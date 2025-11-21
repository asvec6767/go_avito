package delivery

import (
	"main/internal/delivery/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	userHandler *handlers.UserHandler
	// teamHandler *handlers.TeamHandler
	// prHandler *handlers.PRHandler
}

func NewRouter(userHandler *handlers.UserHandler) *Router {
	return &Router{
		userHandler: userHandler,
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	userGroup := router.Group("/users")
	{
		userGroup.POST("/setIsActive", r.userHandler.SetUserIsActive)
	}

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return router
}
