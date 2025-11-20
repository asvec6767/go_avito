package usecase

import "main/internal/domain"

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
	Create()
}
