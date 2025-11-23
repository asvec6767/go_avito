package handlers

import (
	"main/internal/usecase"
	"net/http"

	prusecase "main/internal/usecase/pullrequest"

	"github.com/gin-gonic/gin"
)

type PRHandler struct {
	prUseCase usecase.PullRequestUseCase
}

type PRResponse struct {
	PullRequestId     string   `json:"pull_request_id"`
	PullRequestName   string   `json:"pull_request_name"`
	AuthorId          string   `json:"author_id"`
	Status            string   `json:"status"`
	AssignedReviewers []string `json:"assigned_reviewers,omitempty"`
	MergedAt          string   `json:"mergedAt,omitempty"`
}

func NewPRHandler(prUseCase usecase.PullRequestUseCase) *PRHandler {
	return &PRHandler{prUseCase: prUseCase}
}

func (h *PRHandler) PostPullRequestCreate(c *gin.Context) {
	var request struct {
		AuthorId        string `json:"author_id"`
		PullRequestId   string `json:"pull_request_id"`
		PullRequestName string `json:"pull_request_name"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	if _, err := h.prUseCase.GetById(c.Request.Context(), request.PullRequestId); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "pull request уже существует"})
	}

	pr, err := h.prUseCase.Create(c.Request.Context(), &prusecase.CreatePRRequest{
		ID:       request.PullRequestId,
		Name:     request.PullRequestName,
		AuthorID: request.AuthorId,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	pr, err = h.prUseCase.ChangeAllReviewers(c.Request.Context(), pr.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, PRResponse{
		PullRequestId:     pr.ID,
		PullRequestName:   pr.Name,
		AuthorId:          pr.AuthorID,
		Status:            string(pr.Status),
		AssignedReviewers: pr.ReviewerIDs,
	})
}

func (h *PRHandler) PostPullRequestMerge(c *gin.Context) {
	var request struct {
		PullRequestId string `json:"pull_request_id"`
	}

	type response struct {
		Pr PRResponse `json:"pr"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	pr, err := h.prUseCase.SetMergedStatus(c.Request.Context(), request.PullRequestId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, response{
		Pr: PRResponse{
			PullRequestId:     pr.ID,
			PullRequestName:   pr.Name,
			AuthorId:          pr.AuthorID,
			Status:            string(pr.Status),
			AssignedReviewers: pr.ReviewerIDs,
			MergedAt:          pr.MergedAt.Format("2025-10-24T12:34:56Z"),
		},
	})
}

func (h *PRHandler) PostPullRequestReassign(c *gin.Context) {
	var request struct {
		PullRequestId string `json:"pull_request_id"`
		OldUserId     string `json:"old_user_id"`
	}

	type response struct {
		Pr         PRResponse `json:"pr"`
		ReplacedBy string     `json:"replaced_by"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	// TODO: доделать все варианты ошибок
	pr, replaced_by, err := h.prUseCase.ChangeReviewer(c.Request.Context(), request.PullRequestId, request.OldUserId)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, response{
		Pr: PRResponse{
			PullRequestId:     pr.ID,
			PullRequestName:   pr.Name,
			AuthorId:          pr.AuthorID,
			Status:            string(pr.Status),
			AssignedReviewers: pr.ReviewerIDs,
		},
		ReplacedBy: replaced_by.ID,
	})
}

func (h *PRHandler) GetUsersGetReview(c *gin.Context) {
	type response struct {
		UserId       string       `json:"user_id"`
		PullRequests []PRResponse `json:"pull_requests"`
	}

	userId := c.Param("user_id")
	if userId == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Query argument is required, but not found"})
		return
	}

	prs, err := h.prUseCase.GetListByUserId(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var short_prs []PRResponse
	for _, pr := range prs {
		short_prs = append(short_prs, PRResponse{
			PullRequestId:   pr.ID,
			PullRequestName: pr.Name,
			AuthorId:        pr.AuthorID,
			Status:          string(pr.Status),
		})
	}

	c.JSON(http.StatusOK, response{
		UserId:       userId,
		PullRequests: short_prs,
	})
}
