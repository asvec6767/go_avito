package postgres

import (
	"database/sql"
	"main/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetById(id int) (*domain.User, error) {
	return nil, nil
}
