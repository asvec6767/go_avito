package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique;not null"`
	IsActive bool   `json:"is_active" gorm:"default:false"`
	TeamID   int    `json:"team_id" gorm:"not null;index"`
	Team     Team   `json:"team,omitempty" gorm:"foreignKey:TeamID"`
}

type UserRepository interface {
	GetById(id int) (*User, error)
	GetByName(name string) (*User, error)
	// GetListByTeamId(id int) ([]*User, error)
	// GetList(ids []int) ([]*User, error)
	Create(user *User) (int, error)
	Update(user *User) error
	Delete(id int) error
}
