package team

import (
	"context"
	"main/internal/domain"

	"gorm.io/gorm"
)

type teamUseCase struct {
	team domain.TeamRepository
	user domain.UserRepository
}

type CreateTeamRequest struct {
	Name string
}

func NewTeamUseCase(teamRepo domain.TeamRepository, userRepo domain.UserRepository) *teamUseCase {
	return &teamUseCase{
		team: teamRepo,
		user: userRepo,
	}
}

func (uc *teamUseCase) Create(ctx context.Context, name string) (*domain.Team, error) {
	team := &domain.Team{
		Name: name,
	}

	if err := uc.team.Create(ctx, team); err != nil {
		return nil, err
	}

	return team, nil
}

func (uc *teamUseCase) GetById(ctx context.Context, id string) (*domain.Team, error) {
	return uc.team.GetById(ctx, id)
}

func (uc *teamUseCase) GetByName(ctx context.Context, name string) (*domain.Team, error) {
	return uc.team.GetByName(ctx, name)
}

func (uc *teamUseCase) AddUser(ctx context.Context, team_id, user_id string) error {
	if _, err := uc.team.GetById(ctx, team_id); err != nil {
		return err
	}

	user, err := uc.user.GetById(ctx, user_id)
	if err != nil {
		return err
	}

	user.TeamID = team_id

	err = uc.user.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *teamUseCase) RemoveUser(ctx context.Context, user_id string) error {
	user, err := uc.user.GetById(ctx, user_id)
	if err != nil {
		return err
	}

	user.TeamID = ""

	err = uc.user.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *teamUseCase) SetUsers(ctx context.Context, team_id string, users []domain.User) error {
	if _, err := uc.team.GetById(ctx, team_id); err != nil {
		return err
	}

	for _, user := range users {
		_, err := uc.user.GetById(ctx, user.ID)
		switch err {
		case gorm.ErrRecordNotFound: // Нет такого юзера - создать
			if err := uc.user.Create(ctx, &user); err != nil {
				return err
			}
		case nil: // Есть такой юзер - задать команду
			user.TeamID = team_id
			if err := uc.user.Update(ctx, &user); err != nil {
				return err
			}
		default:
			return err
		}
	}

	return nil
}

func (uc *teamUseCase) RemoveAllUsers(ctx context.Context, team_id string) error {
	_, err := uc.team.GetById(ctx, team_id)
	if err != nil {
		return err
	}

	users, err := uc.user.GetByActiveAndTeam(ctx, team_id)
	if err != nil {
		return err
	}

	for _, user := range users {
		user.TeamID = ""

		err = uc.user.Update(ctx, &user)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc *teamUseCase) Delete(ctx context.Context, id string) error {
	if err := uc.RemoveAllUsers(ctx, id); err != nil {
		return err
	}
	return uc.team.Delete(ctx, id)
}
