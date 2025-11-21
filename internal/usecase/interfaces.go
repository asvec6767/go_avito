package usecase

import (
	"main/internal/domain"
)

type UserUseCase interface {
	Create(name string) (*domain.User, error)
	GetById(id int) (*domain.User, error)
	GetByName(name string) (*domain.User, error)
	GetList(ids []int) ([]*domain.User, error)
	GetListByTeamId(id int) ([]*domain.User, error)
	Activate(id int) (*domain.User, error)
	Deactivate(id int) (*domain.User, error)
	SetIsActive(id int, status bool) (*domain.User, error)
	Delete(id int) error
}

type TeamUseCase interface {
	Create(name string) (*domain.Team, error)
	GetById(id int) (*domain.Team, error)
	GetByName(name string) (*domain.Team, error)
	SetUsers(team_id int, user_ids []int) error
	AddUser(team_id, user_id int) error
	RemoveUser(user_id int) error
	// RemoveAllUsers(user_id []int) error
	Delete(id int) error
}

type PullRequestUseCase interface {
	Create(name string, author_id int) (*domain.PR, error)
	GetById(id int) (*domain.PR, error)
	GetByName(name string) (*domain.PR, error)
	ChangeAllReviewers(id int) (*domain.PR, error)
	ChangeReviewer(pr_id, old_reviewer_id int) (*domain.PR, error)
	SetMergedStatus(pr_id int) (*domain.PR, error)
	SetOpenStatus(pr_id int) (*domain.PR, error)
	Delete(id int) error
}
