package usecase

import (
	"main/internal/domain"
)

type UserUseCase interface {
	Create(name string) (*domain.User, error)
	GetById(id string) (*domain.User, error)
	// GetByName(name string) (*domain.User, error)
	GetList(ids []string) ([]*domain.User, error)
	GetListByTeamId(id string) ([]*domain.User, error)
	Activate(id string) (*domain.User, error)
	Deactivate(id string) (*domain.User, error)
	SetIsActive(id string, status bool) (*domain.User, error)
	Delete(id string) error
}

type TeamUseCase interface {
	Create(name string) (*domain.Team, error)
	GetById(id string) (*domain.Team, error)
	GetByName(name string) (*domain.Team, error)
	SetUsers(team_id string, user_ids []string) error
	AddUser(team_id, user_id string) error
	RemoveUser(user_id string) error
	// RemoveAllUsers(user_id []int) error
	Delete(id string) error
}

type PullRequestUseCase interface {
	Create(name string, author_id string) (*domain.PR, error)
	GetById(id string) (*domain.PR, error)
	GetByName(name string) (*domain.PR, error)
	ChangeAllReviewers(id string) (*domain.PR, error)
	ChangeReviewer(pr_id, old_reviewer_id string) (*domain.PR, error)
	SetMergedStatus(pr_id string) (*domain.PR, error)
	SetOpenStatus(pr_id string) (*domain.PR, error)
	Delete(id string) error
}
