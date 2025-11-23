package models

import "main/internal/domain"

type UserModel struct {
	UserID   string    `gorm:"primaryKey;type:varchar(31)"`
	Username string    `gorm:"not null;unique"`
	IsActive bool      `gorm:"default:false"`
	TeamID   string    `gorm:"not null;index"`
	Team     TeamModel `gorm:"foreignKey:TeamID"`
}

func (m UserModel) ToDomain() domain.User {
	return domain.User{
		ID:       m.UserID,
		Username: m.Username,
		IsActive: m.IsActive,
		TeamID:   m.TeamID,
	}
}

func UserToModel(u domain.User) UserModel {
	return UserModel{
		UserID:   u.ID,
		Username: u.Username,
		IsActive: u.IsActive,
		TeamID:   u.TeamID,
	}
}
