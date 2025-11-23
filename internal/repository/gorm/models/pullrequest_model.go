package models

import (
	"main/internal/domain"
	"time"
)

type PRModel struct {
	ID       string `gorm:"primaryKey;type:varchar(36)"`
	Name     string `gorm:"not null"`
	AuthorID string `gorm:"not null;index"`
	Status   string `gorm:"not null;default:'OPEN'"`
	TeamID   string `gorm:"not null;index"`
	MergedAt *time.Time

	Author    UserModel   `gorm:"foreignKey:AuthorID"`
	Reviewers []UserModel `gorm:"many2many:pr_reviewers"`
}

func (m PRModel) ToDomain() domain.PR {
	var reviewerIDs []string
	for _, reviewer := range m.Reviewers {
		reviewerIDs = append(reviewerIDs, reviewer.UserID)
	}

	// TODO: сделать время мерджа только при первом мердже
	var mergedAt *time.Time
	if m.MergedAt != nil {
		mergedAt = m.MergedAt
	}

	return domain.PR{
		ID:          m.ID,
		Name:        m.Name,
		AuthorID:    m.AuthorID,
		Status:      domain.PRStatus(m.Status),
		TeamID:      m.TeamID,
		ReviewerIDs: reviewerIDs,
		MergedAt:    mergedAt,
	}
}

func PRToModel(pr domain.PR) PRModel {
	return PRModel{
		ID:       pr.ID,
		Name:     pr.Name,
		AuthorID: pr.AuthorID,
		Status:   string(pr.Status),
		TeamID:   pr.TeamID,
		MergedAt: pr.MergedAt,
	}
}
