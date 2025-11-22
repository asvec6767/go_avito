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

func (r *userRepository) GetById(id string) (*domain.User, error) {
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

// func (r *userRepository) GetByName(name string) (*domain.User, error) {
// 	var user domain.User

// 	err := r.db.Where("name = ?", name).First(&user).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, domain.ErrUserNotFound
// 		}
// 		return nil, err
// 	}

// 	return &user, nil
// }

func (r *userRepository) Create(user *domain.User) error {
	var existingUser domain.User
	err := r.db.Where("username = ?", user.Username).First(&existingUser).Error
	if err == nil {
		return domain.ErrUserAlreadyExists
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}

	return r.db.Create(user).Error
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

func (r *userRepository) Delete(id string) error {
	result := r.db.Delete(&domain.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}
