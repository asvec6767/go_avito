package gorm

import (
	"main/internal/domain"

	"gorm.io/gorm"
)

type teamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) domain.TeamRepository {
	return &teamRepository{
		db: db,
	}
}

func (r *teamRepository) GetById(id string) (*domain.Team, error) {
	var team domain.Team

	err := r.db.First(&team, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrTeamNotFound
		}
		return nil, err
	}

	return &team, nil
}

func (r *teamRepository) GetByName(name string) (*domain.Team, error) {
	var team domain.Team

	err := r.db.Where("name = ?", name).First(&team).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &team, nil
}

func (r *teamRepository) Create(team *domain.Team) error {
	var existingTeam domain.Team
	err := r.db.Where("name = ?", team.Name).First(&existingTeam).Error
	if err == nil {
		return domain.ErrTeamAlreadyExists
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}

	err = r.db.Create(team).Error
	return err
}

func (r *teamRepository) Update(team *domain.Team) error {
	result := r.db.Save(team)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrTeamNotFound
	}
	return nil
}

func (r *teamRepository) Delete(id string) error {
	result := r.db.Delete(&domain.Team{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrTeamNotFound
	}
	return nil
}
