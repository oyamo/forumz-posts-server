package main

import (
	"go.uber.org/zap"
	"log"
	"posts-server/internal/domain/connections"
	"posts-server/internal/domain/persons"
	"posts-server/internal/domain/posts"
	"posts-server/internal/infrastructure/postgres"
	"posts-server/internal/interfaces/kafka"
	"posts-server/internal/interfaces/web"
	"posts-server/internal/pkg"
)

func main() {
	logConf, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	logger := logConf.Sugar()

	conf, err := pkg.NewConfig()
	if err != nil {
		logger.Fatal(err)
	}

	publicKey, err := getPublicKeyFromFile(conf.PublicKey)
	if err != nil {
		logger.Fatal(err)
	}

	conn, err := pkg.NewPostgresClient(conf.DatabaseDSN)
	defer conn.Close()
	if err != nil {
		logger.Fatal(err)
	}

	jsonSender := pkg.NewJSONSender(conf)

	_ = jsonSender

	personRepo := postgres.NewPersonRepository(conn)
	connectionRepo := postgres.NewConnectionRepository(conn)
	postsRepo := postgres.NewPostsRepository(conn)
	connectionsUC := connections.NewUseCase(connectionRepo)
	personsUC := persons.NewUseCase(personRepo, logger)
	postsUC := posts.NewUseCase(postsRepo, logger)

	consumer := pkg.NewConsumer("posts-server", conf.KafkaConsumerServer)
	kafkaCRouter := kafka.NewConsumerRouter(consumer, logger, connectionsUC, personsUC)
	kafkaCRouter.Consume()

	router := web.NewRouter(logger, publicKey, postsUC)
	engine := router.InitializeRoutes()
	engine.Run(":3010")
}
