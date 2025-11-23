package delivery

import (
	"main/internal/delivery/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	userHandler *handlers.UserHandler
	teamHandler *handlers.TeamHandler
	prHandler   *handlers.PRHandler
}

func NewRouter(userHandler *handlers.UserHandler, teamHandler *handlers.TeamHandler, prHandler *handlers.PRHandler) *Router {
	return &Router{
		userHandler: userHandler,
		teamHandler: teamHandler,
		prHandler:   prHandler,
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	userGroup := router.Group("/users")
	{
		userGroup.POST("/setIsActive", r.userHandler.PostUsersSetIsActive)
		userGroup.GET("/getReview", r.prHandler.GetUsersGetReview)
	}
	teamGroup := router.Group("/team")
	{
		teamGroup.POST("/add", r.teamHandler.PostTeamAdd)
		teamGroup.GET("/get", r.teamHandler.GetTeamGet)
	}
	prGroup := router.Group("/pullRequest")
	{
		prGroup.POST("/create", r.prHandler.PostPullRequestCreate)
		prGroup.POST("/merge", r.prHandler.PostPullRequestMerge)
		prGroup.POST("/reassign", r.prHandler.PostPullRequestReassign)
	}

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return router
}
