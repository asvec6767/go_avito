package domain

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name  string  `json:"name" gorm:"not null;unique"`
	Users []*User `json:"users,omitempty" gorm:"foreignKey:TeamID"`
	PRs   []*PR   `json:"prs,omitempty" gorm:"foreignKey:TeamID"`
}

type TeamRepository interface {
	GetById(id int) (*Team, error)
	GetByName(name string) (*Team, error)
	Create(team *Team) (int, error)
	Update(team *Team) error
	Delete(id int) error
}
