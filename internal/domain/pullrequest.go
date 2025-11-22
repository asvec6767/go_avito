package domain

import _ "gorm.io/gorm"

type PRStatus string

const (
	PullRequestStatusOpen   PRStatus = "OPEN"
	PullRequestStatusMerged PRStatus = "MERGED"
)

type PR struct {
	ID        string   `json:"id" gorm:"unique;not null;primarykey"`
	Name      string   `json:"name" gorm:"not null;unique"`
	AuthorID  string   `json:"author_id" gorm:"not null;index"`
	Author    User     `json:"author,omitempty" gorm:"foreignKey:AuthorID"`
	Status    PRStatus `json:"status" gorm:"default:'OPEN'"`
	TeamID    string   `json:"team_id" gorm:"not null;index"`
	Reviewers []*User  `json:"reviewers" gorm:"many2many:pr_users"`
}

type PRRepository interface {
	GetById(id string) (*PR, error)
	// GetByName(name string) (*PR, error)
	// GetByTeamId(id int) (*PR, error)
	Create(team *PR) error
	Update(team *PR) error
	Delete(id string) error
}
