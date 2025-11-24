package domain

import (
	"context"
)

type User struct {
	ID       string
	Username string
	Email    string
	IsActive bool
	TeamID   string
}

type UserRepository interface {
	GetById(ctx context.Context, id string) (*User, error)
	GetByActiveAndTeam(ctx context.Context, team_id string) ([]User, error)
	GetByTeam(ctx context.Context, team_id string) ([]User, error)
	// GetByName(name string) (*User, error)
	// GetListByTeamId(id int) ([]*User, error)
	// GetList(ids []int) ([]*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
}
