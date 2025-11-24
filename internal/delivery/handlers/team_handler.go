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

type CreateTeamRequest struct {
	Name string `json:"name" binding:"required"`
}

type TeamResponse struct {
	TeamName string       `json:"team_name"`
	Members  UserResponse `json:"members"`
}

type TeamCreateRequest struct {
	TeamName string              `json:"team_name" binding:"required"`
	Members  []CreateUserRequest `json:"members"`
}

func NewTeamHandler(teamUseCase usecase.TeamUseCase) *TeamHandler {
	return &TeamHandler{teamUseCase: teamUseCase}
}

func (h *TeamHandler) PostTeamAdd(c *gin.Context) {
	var req TeamCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := h.teamUseCase.Create(c.Request.Context(), req.TeamName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users := make([]domain.User, len(req.Members))
	for i, user := range req.Members {
		users[i] = domain.User{
			ID:       user.UserID,
			Username: user.Username,
			IsActive: user.IsActive,
			TeamID:   team.ID,
		}
	}

	if err = h.teamUseCase.SetUsers(c.Request.Context(), team.ID, users); err != nil {
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

	team, users, err := h.teamUseCase.GetByNameWithUsers(c.Request.Context(), teamName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	resp := TeamResponse{
		TeamName: team.Name,
	}
	for _, user := range *users {
		resp.Members.ID = user.ID
		resp.Members.Username = user.Username
		resp.Members.IsActive = user.IsActive
	}

	c.JSON(http.StatusOK, resp)
}
