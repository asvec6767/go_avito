package user

import "main/internal/domain"

type userUseCase struct {
	user domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) *userUseCase {
	return &userUseCase{
		user: userRepo,
	}
}

func (uc *userUseCase) Create(name string) (*domain.User, error) {
	user := &domain.User{
		Name:     name,
		IsActive: false,
	}

	id, err := uc.user.Create(user)
	if err != nil {
		return nil, err
	}

	user.ID = id

	return user, nil
}

func (uc *userUseCase) GetById(id int) (*domain.User, error) {
	return uc.user.GetById(id)
}

func (uc *userUseCase) GetByName(name string) (*domain.User, error) {
	return uc.user.GetByName(name)
}

func (uc *userUseCase) GetList(ids []int) ([]*domain.User, error) {
	return uc.user.GetList(ids)
}

func (uc *userUseCase) Activate(id int) (*domain.User, error) {
	user, err := uc.user.GetById(id)
	if err != nil {
		return nil, err
	}

	user.IsActive = true

	err = uc.user.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) Deactivate(id int) (*domain.User, error) {
	user, err := uc.user.GetById(id)
	if err != nil {
		return nil, err
	}

	user.IsActive = false

	err = uc.user.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) Delete(id int) error {
	return uc.user.Delete(id)
}
