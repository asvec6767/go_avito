package gorm

import (
	"context"
	"main/internal/domain"
	"main/internal/repository/gorm/models"

	"gorm.io/gorm"
)

type prRepository struct {
	db *gorm.DB
}

func NewPRRepository(db *gorm.DB) domain.PRRepository {
	return &prRepository{
		db: db,
	}
}

func (r *prRepository) GetById(ctx context.Context, id string) (*domain.PR, error) {
	var model models.PRModel

	err := r.db.WithContext(ctx).First(&model, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrPRNotFound
		}
		return nil, err
	}

	pr := model.ToDomain()

	return &pr, nil
}

func (r *prRepository) GetWithReviewers(ctx context.Context, id string) (*domain.PR, error) {
	var model models.PRModel

	err := r.db.WithContext(ctx).
		Preload("Reviewers", "is_active = ?", true).
		First(&model, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrPRNotFound
		}
		return nil, err
	}

	pr := model.ToDomain()

	return &pr, nil
}

func (r *prRepository) Create(ctx context.Context, pr *domain.PR) error {
	var existingPRModel models.PRModel

	err := r.db.WithContext(ctx).First(&existingPRModel, "name = ?", pr.Name).Error
	if err == nil {
		return domain.ErrPRAlreadyExists
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}

	model := models.PRToModel(*pr)

	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *prRepository) Update(ctx context.Context, pr *domain.PR) error {
	model := models.PRToModel(*pr)

	result := r.db.WithContext(ctx).Save(&model)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrPRNotFound
	}
	return nil
}

func (r *prRepository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&models.PRModel{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrPRNotFound
	}
	return nil
}

func (r *prRepository) GetByReviewerAndStatus(ctx context.Context, reviewerID string, status domain.PRStatus) ([]domain.PR, error) {
	var models []models.PRModel

	// TODO: перепроверить правильность JOIN
	err := r.db.WithContext(ctx).
		Joins("JOIN pr_reviewers ON pr_reviewers.pr_model_id = pr_models.id").
		Where("pr_reviewers.user_model_user_id = ? AND status = ?", reviewerID, status).
		Find(&models).Error
	if err != nil {
		return nil, err
	}

	prs := make([]domain.PR, len(models))
	for i, model := range models {
		prs[i] = model.ToDomain()
	}

	return prs, nil
}

func (r *prRepository) GetByTeam(ctx context.Context, teamID string) ([]domain.PR, error) {
	var models []models.PRModel

	err := r.db.WithContext(ctx).
		Where("team_id = ?", teamID).
		Find(&models).Error
	if err != nil {
		return nil, err
	}

	prs := make([]domain.PR, len(models))
	for i, model := range models {
		prs[i] = model.ToDomain()
	}

	return prs, nil
}
