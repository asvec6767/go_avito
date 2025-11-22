package domain

import _ "gorm.io/gorm"

type Team struct {
	ID    string  `json:"id" gorm:"unique;not null;primarykey"`
	Name  string  `json:"name" gorm:"not null;unique"`
	Users []*User `json:"users,omitempty" gorm:"foreignKey:TeamID"`
	PRs   []*PR   `json:"prs,omitempty" gorm:"foreignKey:TeamID"`
}

type TeamRepository interface {
	GetById(id string) (*Team, error)
	// GetByName(name string) (*Team, error)
	Create(team *Team) error
	Update(team *Team) error
	Delete(id string) error
}
