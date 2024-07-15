package web

import (
	"crypto/rsa"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"posts-server/internal/domain/posts"
	"posts-server/internal/interfaces/web/handlers"
)

type Router struct {
	logger  *zap.SugaredLogger
	postsUC *posts.UseCase
	pub     *rsa.PublicKey
}

func (router *Router) InitializeRoutes() *gin.Engine {
	postsHandler := handlers.NewPostsHandler(router.postsUC, router.logger)
	middlewareHandler := handlers.NewMiddlewareHandler(router.pub, router.logger)

	engine := gin.Default()
	v1 := engine.Group("/api/v1")
	v1.Use(middlewareHandler.AddRequestID)

	postsSubRouter := v1.Group("/posts")
	postsSubRouter.POST("/", middlewareHandler.AuthenticateRequest, postsHandler.CreatePost)
	postsSubRouter.GET("/:postId", middlewareHandler.AuthenticateRequest, postsHandler.ViewPost)
	postsSubRouter.DELETE("/:postId", middlewareHandler.AuthenticateRequest, postsHandler.DeletePost)
	return engine

}

func NewRouter(logger *zap.SugaredLogger, pub *rsa.PublicKey, postsUC *posts.UseCase) *Router {
	return &Router{
		logger:  logger,
		pub:     pub,
		postsUC: postsUC,
	}
}
