package domain

import "context"

type Team struct {
	ID   string
	Name string
}

type TeamRepository interface {
	GetById(ctx context.Context, id string) (*Team, error)
	GetByName(ctx context.Context, name string) (*Team, error)
	Create(ctx context.Context, team *Team) error
	Update(ctx context.Context, team *Team) error
	Delete(ctx context.Context, id string) error
}
