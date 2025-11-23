package gorm

import (
	"context"
	"main/internal/domain"
	"main/internal/repository/gorm/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetById(ctx context.Context, id string) (*domain.User, error) {
	var model models.UserModel

	err := r.db.WithContext(ctx).Where("user_id = ?", id).First(&model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	user := model.ToDomain()

	return &user, nil
}

func (r *userRepository) GetByActiveAndTeam(ctx context.Context, teamID string) ([]domain.User, error) {
	var models []models.UserModel

	err := r.db.WithContext(ctx).
		Where("team_id = ? AND is_active = ?", teamID, true).
		Find(&models).Error
	if err != nil {
		return nil, err
	}

	users := make([]domain.User, len(models))
	for i, model := range models {
		users[i] = model.ToDomain()
	}

	return users, nil
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	model := models.UserToModel(*user)

	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	model := models.UserToModel(*user)

	result := r.db.WithContext(ctx).Save(&model)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&models.UserModel{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}
