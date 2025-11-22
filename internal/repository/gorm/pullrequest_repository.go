package gorm

import (
	"main/internal/domain"

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

func (r *prRepository) GetById(id string) (*domain.PR, error) {
	var pr domain.PR

	err := r.db.First(&pr, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrPRNotFound
		}
		return nil, err
	}

	return &pr, nil
}

// func (r *prRepository) GetByName(name string) (*domain.PR, error) {
// 	var pr domain.PR

// 	err := r.db.Where("name = ?", name).First(&pr).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, domain.ErrPRNotFound
// 		}
// 		return nil, err
// 	}

// 	return &pr, nil
// }

func (r *prRepository) Create(pr *domain.PR) error {
	var existingPR domain.PR
	err := r.db.Where("name = ?", pr.Name).First(&existingPR).Error
	if err == nil {
		return domain.ErrPRAlreadyExists
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}

	err = r.db.Create(pr).Error
	return err
}

func (r *prRepository) Update(pr *domain.PR) error {
	result := r.db.Save(pr)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrPRNotFound
	}
	return nil
}

func (r *prRepository) Delete(id string) error {
	result := r.db.Delete(&domain.PR{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrPRNotFound
	}
	return nil
}
