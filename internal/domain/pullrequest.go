package domain

import (
	"context"
	"time"
)

type PRStatus string

const (
	PullRequestStatusOpen   PRStatus = "OPEN"
	PullRequestStatusMerged PRStatus = "MERGED"
)

type PR struct {
	ID          string
	Name        string
	AuthorID    string
	Status      PRStatus
	TeamID      string
	ReviewerIDs []string
	MergedAt    *time.Time
}

type PRRepository interface {
	GetById(ctx context.Context, id string) (*PR, error)
	GetWithReviewers(ctx context.Context, id string) (*PR, error)
	GetByReviewerAndStatus(ctx context.Context, reviewerID string, status PRStatus) ([]PR, error)
	GetByTeam(ctx context.Context, teamID string) ([]PR, error)
	// GetByName(name string) (*PR, error)
	// GetByTeamId(id int) (*PR, error)
	Create(ctx context.Context, team *PR) error
	Update(ctx context.Context, team *PR) error
	Delete(ctx context.Context, id string) error
}
