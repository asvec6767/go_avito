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
		// TeamID: 0,
	}

	if err := uc.user.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) GetById(id string) (*domain.User, error) {
	return uc.user.GetById(id)
}

// func (uc *userUseCase) GetByName(name string) (*domain.User, error) {
// 	return uc.user.GetByName(name)
// }

// func (uc *userUseCase) GetList(ids []int) ([]*domain.User, error) {
// 	return uc.user.GetList(ids)
// }

// func (uc *userUseCase) GetListByTeamId(id int) ([]*domain.User, error) {
// 	return uc.user.GetListByTeamId(id)
// }

func (uc *userUseCase) Activate(id string) (*domain.User, error) {
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

func (uc *userUseCase) Deactivate(id string) (*domain.User, error) {
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

func (uc *userUseCase) SetIsActive(id string, status bool) (*domain.User, error) {
	user, err := uc.user.GetById(id)
	if err != nil {
		return nil, err
	}

	user.IsActive = status

	err = uc.user.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) Delete(id string) error {
	return uc.user.Delete(id)
}
