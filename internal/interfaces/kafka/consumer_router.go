package kafka

import (
	"go.uber.org/zap"
	"posts-server/internal/domain/connections"
	"posts-server/internal/domain/persons"
	"posts-server/internal/interfaces/kafka/handlers"
	"posts-server/internal/pkg"
)

type ConsumerRouter struct {
	consumer      *pkg.KakfaConsumer
	logger        *zap.SugaredLogger
	connectionsUC *connections.UseCase
	personsUC     *persons.UseCase
}

func (r *ConsumerRouter) Consume() {
	personsHandler := handlers.NewPersonHandler(r.logger, r.personsUC)
	connectionsHandler := handlers.NewConnectionHandler(r.connectionsUC, r.logger)

	r.consumer.ConsumeAndHandle("Put-Person-v1", personsHandler.PutPerson)
	r.consumer.ConsumeAndHandle("Put-Connection-v1", connectionsHandler.Connect)
	r.consumer.ConsumeAndHandle("Delete-Connection-v1", connectionsHandler.Disconnect)
}

func NewConsumerRouter(consumer *pkg.KakfaConsumer, logger *zap.SugaredLogger, connectionsUC *connections.UseCase, personsUC *persons.UseCase) *ConsumerRouter {
	return &ConsumerRouter{
		consumer:      consumer,
		logger:        logger,
		connectionsUC: connectionsUC,
		personsUC:     personsUC,
	}
}
