package user

import (
	"context"
	"main/internal/domain"
)

type userUseCase struct {
	user domain.UserRepository
	team domain.TeamRepository
}

type CreateUserRequest struct {
	ID       string
	Username string
	TeamID   string
	IsActive bool
}

func NewUserUseCase(userRepo domain.UserRepository, teamRepo domain.TeamRepository) *userUseCase {
	return &userUseCase{
		user: userRepo,
		team: teamRepo,
	}
}

func (uc *userUseCase) Create(ctx context.Context, req *CreateUserRequest) (*domain.User, error) {
	_, err := uc.team.GetById(ctx, req.TeamID)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:       req.ID,
		Username: req.Username,
		IsActive: req.IsActive, // default inactive
		TeamID:   req.TeamID,
	}

	if err := uc.user.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) GetById(ctx context.Context, id string) (*domain.User, error) {
	return uc.user.GetById(ctx, id)
}

func (uc *userUseCase) Activate(ctx context.Context, id string) (*domain.User, error) {
	user, err := uc.user.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	user.IsActive = true

	err = uc.user.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) Deactivate(ctx context.Context, id string) (*domain.User, error) {
	user, err := uc.user.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	user.IsActive = false

	err = uc.user.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) SetIsActive(ctx context.Context, id string, status bool) (*domain.User, error) {
	user, err := uc.user.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	user.IsActive = status

	err = uc.user.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) Delete(ctx context.Context, id string) error {
	return uc.user.Delete(ctx, id)
}
