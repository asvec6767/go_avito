package gorm

import (
	"main/internal/domain"

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

func (r *userRepository) GetById(id int) (*domain.User, error) {
	var user domain.User

	err := r.db.First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByName(name string) (*domain.User, error) {
	var user domain.User

	err := r.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(user *domain.User) (int, error) {
	var existingUser domain.User
	err := r.db.Where("name = ?", user.Name).First(&existingUser).Error
	if err == nil {
		return 0, domain.ErrUserAlreadyExists
	}
	if err != gorm.ErrRecordNotFound {
		return 0, err
	}

	err = r.db.Create(user).Error
	return int(user.ID), err
}

func (r *userRepository) Update(user *domain.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

func (r *userRepository) Delete(id int) error {
	result := r.db.Delete(&domain.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}
