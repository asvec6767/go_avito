package team

import "main/internal/domain"

type teamUseCase struct {
	Team domain.TeamRepository
	User domain.UserRepository
}

func NewTeamUseCase(team domain.TeamRepository, user domain.UserRepository) *teamUseCase {
	return &teamUseCase{
		Team: team,
		User: user,
	}
}

func (uc *teamUseCase) Create(name string) (*domain.Team, error) {
	team := &domain.Team{
		Name: name,
	}

	id, err := uc.Team.Create(team)
	if err != nil {
		return nil, err
	}

	team.ID = id

	return team, nil
}

func (uc *teamUseCase) GetById(id int) (*domain.Team, error) {
	return uc.Team.GetById(id)
}

func (uc *teamUseCase) GetByName(name string) (*domain.Team, error) {
	return uc.Team.GetByName(name)
}

func (uc *teamUseCase) AddUser(team_id, user_id int) error {
	if _, err := uc.Team.GetById(team_id); err != nil {
		return err
	}

	user, err := uc.User.GetById(user_id)
	if err != nil {
		return err
	}

	user.TeamID = team_id

	err = uc.User.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *teamUseCase) RemoveUser(user_id int) error {
	user, err := uc.User.GetById(user_id)
	if err != nil {
		return err
	}

	user.TeamID = 0

	err = uc.User.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *teamUseCase) SetUsers(team_id int, user_ids []int) error {
	if _, err := uc.Team.GetById(team_id); err != nil {
		return err
	}

	for _, user_id := range user_ids {
		user, err := uc.User.GetById(user_id)
		if err != nil {
			return err
		}

		user.TeamID = team_id

		err = uc.User.Update(user)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: сделать норм обработку ошибки в этой функции
// func (uc *teamUseCase) RemoveAllUsers(team_id int) error {
// 	for user, err := uc.User.GetByTeamId(team_id); user!=nil && err==nil; {
// 		user.TeamID=0;

// 		err = uc.User.Update(user)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

func (uc *teamUseCase) Delete(id int) error {
	return uc.Team.Delete(id)
}
