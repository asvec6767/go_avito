package domain

import _ "gorm.io/gorm"

type User struct {
	ID       string `json:"id" gorm:"unique;not null;primarykey"`
	Name     string `json:"name" gorm:"not null"`
	IsActive bool   `json:"is_active" gorm:"default:false"`
	TeamID   string `json:"team_id" gorm:"not null;index"`
	Team     Team   `json:"team,omitempty" gorm:"foreignKey:TeamID"`
}

type UserRepository interface {
	GetById(id string) (*User, error)
	// GetByName(name string) (*User, error)
	// GetListByTeamId(id int) ([]*User, error)
	// GetList(ids []int) ([]*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
}
