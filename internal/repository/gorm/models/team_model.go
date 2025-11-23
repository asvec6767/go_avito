package models

import (
	"main/internal/domain"
)

type TeamModel struct {
	ID   string `gorm:"primaryKey;type:varchar(31)"`
	Name string `gorm:"not null;unique"`
}

func (m TeamModel) ToDomain() domain.Team {
	return domain.Team{
		ID:   m.ID,
		Name: m.Name,
	}
}

func TeamToModel(t domain.Team) TeamModel {
	return TeamModel{
		ID:   t.ID,
		Name: t.Name,
	}
}
