package usecase

import (
	"main/internal/domain"
)

type UserUseCase interface {
	Create(name string) (*domain.User, error)
	GetById(id int) (*domain.User, error)
	GetByName(name string) (*domain.User, error)
	GetList(ids []int) ([]*domain.User, error)
	Activate(id int) (*domain.User, error)
	Deactivate(id int) (*domain.User, error)
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
