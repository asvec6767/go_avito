package team

import "main/internal/domain"

type teamUseCase struct {
	team domain.TeamRepository
	user domain.UserRepository
}

func NewTeamUseCase(teamRepo domain.TeamRepository, userRepo domain.UserRepository) *teamUseCase {
	return &teamUseCase{
		team: teamRepo,
		user: userRepo,
	}
}

func (uc *teamUseCase) Create(name string) (*domain.Team, error) {
	team := &domain.Team{
		Name: name,
	}

	if err := uc.team.Create(team); err != nil {
		return nil, err
	}

	return team, nil
}

func (uc *teamUseCase) GetById(id string) (*domain.Team, error) {
	return uc.team.GetById(id)
}

// func (uc *teamUseCase) GetByName(name string) (*domain.Team, error) {
// 	return uc.team.GetByName(name)
// }

func (uc *teamUseCase) AddUser(team_id, user_id string) error {
	if _, err := uc.team.GetById(team_id); err != nil {
		return err
	}

	user, err := uc.user.GetById(user_id)
	if err != nil {
		return err
	}

	user.TeamID = team_id

	err = uc.user.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *teamUseCase) RemoveUser(user_id string) error {
	user, err := uc.user.GetById(user_id)
	if err != nil {
		return err
	}

	user.TeamID = ""

	err = uc.user.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *teamUseCase) SetUsers(team_id string, user_ids []string) error {
	if _, err := uc.team.GetById(team_id); err != nil {
		return err
	}

	for _, user_id := range user_ids {
		user, err := uc.user.GetById(user_id)
		if err != nil {
			return err
		}

		user.TeamID = team_id

		err = uc.user.Update(user)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc *teamUseCase) RemoveAllUsers(team_id string) error {
	team, err := uc.team.GetById(team_id)
	if err != nil {
		return err
	}

	// users, err := uc.user.GetListByTeamId(team_id)
	// if err != nil {
	// 	return err
	// }

	for _, user := range team.Users {
		user.TeamID = ""

		err = uc.user.Update(user)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc *teamUseCase) Delete(id string) error {
	if err := uc.RemoveAllUsers(id); err != nil {
		return err
	}
	return uc.team.Delete(id)
}
