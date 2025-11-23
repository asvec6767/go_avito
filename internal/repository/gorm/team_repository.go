package gorm

import (
	"context"
	"main/internal/domain"
	"main/internal/repository/gorm/models"

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

func (r *teamRepository) GetById(ctx context.Context, id string) (*domain.Team, error) {
	var model models.TeamModel

	err := r.db.WithContext(ctx).First(&model, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrTeamNotFound
		}
		return nil, err
	}

	team := model.ToDomain()

	return &team, nil
}

func (r *teamRepository) GetByName(ctx context.Context, name string) (*domain.Team, error) {
	var model models.TeamModel

	err := r.db.WithContext(ctx).First(&model, "name = ?", name).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	team := model.ToDomain()

	return &team, nil
}

func (r *teamRepository) Create(ctx context.Context, team *domain.Team) error {
	var existingTeamModel models.TeamModel

	err := r.db.WithContext(ctx).First(&existingTeamModel, "name = ?", team.Name).Error
	if err == nil {
		return domain.ErrTeamAlreadyExists
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}

	model := models.TeamToModel(*team)

	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *teamRepository) Update(ctx context.Context, team *domain.Team) error {
	model := models.TeamToModel(*team)

	result := r.db.WithContext(ctx).Save(&model)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrTeamNotFound
	}
	return nil
}

func (r *teamRepository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&models.TeamModel{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrTeamNotFound
	}
	return nil
}
