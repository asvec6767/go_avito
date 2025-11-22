package handlers

import (
	"fmt"
	"main/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	teamUseCase usecase.TeamUseCase
}

func NewTeamHandler(teamUseCase usecase.TeamUseCase) *TeamHandler {
	return &TeamHandler{teamUseCase: teamUseCase}
}

// PostTeamAdd operation middleware
func (h *TeamHandler) PostTeamAdd(c *gin.Context) {
	// TeamMember defines model for TeamMember.
	type TeamMember struct {
		IsActive bool   `json:"is_active"`
		UserId   string `json:"user_id"`
		Username string `json:"username"`
	}

	// Team defines model for Team.
	type Team struct {
		Members  []TeamMember `json:"members"`
		TeamName string       `json:"team_name"`
	}

	var request Team

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		c.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTeamAdd(ctx, request.(PostTeamAddRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTeamAdd")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostTeamAddResponseObject); ok {
		if err := validResponse.VisitPostTeamAddResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
