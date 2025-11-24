package handlers

import (
	"main/internal/usecase"
	"main/internal/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

type CreateUserRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	Username string `json:"username" binding:"required"`
	TeamID   string `json:"team_id,omitempty"`
	IsActive bool   `json:"is_active" binding:"required"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
	TeamID   string `json:"team_id,omitempty"`
}

type UpdateActiveRequest struct {
	IsActive bool `json:"is_active" binding:"required"`
}

type SetStatusUserRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	IsActive bool   `json:"is_active" binding:"required"`
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) GetUserGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "id пользователя не найден"})
		return
	}

	user, err := h.userUseCase.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:       user.ID,
		Username: user.Username,
		IsActive: user.IsActive,
		TeamID:   user.TeamID,
	})
}

func (h *UserHandler) PostUserCreate(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "не валидное тело запроса " + err.Error(),
		})
		return
	}

	createReq := &user.CreateUserRequest{
		ID:       req.UserID,
		Username: req.Username,
		IsActive: req.IsActive,
		TeamID:   req.TeamID,
	}

	user, err := h.userUseCase.Create(c.Request.Context(), createReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, UserResponse{
		ID:       user.ID,
		Username: user.Username,
		IsActive: user.IsActive,
		TeamID:   user.TeamID,
	})
}

func (h *UserHandler) PostUsersSetIsActive(c *gin.Context) {
	var req SetStatusUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userUseCase.SetIsActive(c.Request.Context(), req.UserID, req.IsActive)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:       user.ID,
		Username: user.Username,
		TeamID:   user.TeamID,
		IsActive: user.IsActive,
	})
}
