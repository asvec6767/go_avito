package handlers

import (
	"main/internal/domain"
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
	// type TeamMember struct {
	// 	IsActive bool   `json:"is_active"`
	// 	UserId   string `json:"user_id"`
	// 	Username string `json:"username"`
	// }

	// Team defines model for Team.
	// type Team struct {
	// 	Members  []TeamMember `json:"members"`
	// 	TeamName string       `json:"team_name"`
	// }

	// var request Team

	var request struct {
		Members  []domain.User `json:"members"`
		TeamName string        `json:"team_name"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := h.teamUseCase.Create(request.TeamName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = h.teamUseCase.SetUsers(team.ID, request.Members); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, team)
}

func (h *TeamHandler) GetTeamGet(c *gin.Context) {
	teamName := c.Param("team_name")
	if teamName == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Query argument is required, but not found"})
		return
	}

	team, err := h.teamUseCase.GetByName(teamName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}
