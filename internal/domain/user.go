package domain

import _ "gorm.io/gorm"

type User struct {
	UserId       string `json:"user_id" gorm:"unique;not null;primarykey"`
	Username     string `json:"username" gorm:"not null"`
	IsActive     bool   `json:"is_active" gorm:"default:false"`
	TeamID       string `json:"team_id,omitempty" gorm:"not null;index"`
	Team         Team   `json:"team,omitempty" gorm:"foreignKey:TeamID"`
	PullRequests []*PR  `json:"pullrequests" gorm:"many2many:pr_users"`
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
