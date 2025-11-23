package usecase

import (
	"context"
	"main/internal/domain"
	userusecase "main/internal/usecase/user"

	prusecase "main/internal/usecase/pullrequest"
)

type UserUseCase interface {
	Create(ctx context.Context, req *userusecase.CreateUserRequest) (*domain.User, error)
	GetById(ctx context.Context, id string) (*domain.User, error)
	Activate(ctx context.Context, id string) (*domain.User, error)
	Deactivate(ctx context.Context, id string) (*domain.User, error)
	SetIsActive(ctx context.Context, id string, status bool) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}

type TeamUseCase interface {
	Create(ctx context.Context, name string) (*domain.Team, error)
	GetById(ctx context.Context, id string) (*domain.Team, error)
	GetByName(ctx context.Context, name string) (*domain.Team, error)
	SetUsers(ctx context.Context, team_id string, users []domain.User) error
	AddUser(ctx context.Context, team_id, user_id string) error
	RemoveUser(ctx context.Context, user_id string) error
	// RemoveAllUsers(user_id []int) error
	Delete(ctx context.Context, id string) error
}

type PullRequestUseCase interface {
	Create(ctx context.Context, req *prusecase.CreatePRRequest) (*domain.PR, error)
	GetById(ctx context.Context, id string) (*domain.PR, error)
	// GetByName(name string) (*domain.PR, error)
	ChangeAllReviewers(ctx context.Context, id string) (*domain.PR, error)
	ChangeReviewer(ctx context.Context, pr_id, old_reviewer_id string) (*domain.PR, *domain.User, error)
	SetMergedStatus(ctx context.Context, pr_id string) (*domain.PR, error)
	SetOpenStatus(ctx context.Context, pr_id string) (*domain.PR, error)
	GetListByUserId(ctx context.Context, user_id string) ([]domain.PR, error)
	Delete(ctx context.Context, id string) error
}
