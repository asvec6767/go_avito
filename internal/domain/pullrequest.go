package domain

import "gorm.io/gorm"

type PRStatus string

const (
	PullRequestStatusOpen   PRStatus = "OPEN"
	PullRequestStatusMerged PRStatus = "MERGED"
)

type PR struct {
	gorm.Model
	Name      string   `json:"name" gorm:"not null;unique"`
	AuthorID  int      `json:"author_id" gorm:"not null;index"`
	Author    User     `json:"author,omitempty" gorm:"foreignKey:AuthorID"`
	Status    PRStatus `json:"status" gorm:"default:'OPEN'"`
	TeamID    int      `json:"team_id" gorm:"not null;index"`
	Reviewers []*User  `json:"reviewers" gorm:"many2many:pr_users"`
}

type PRRepository interface {
	GetById(id int) (*PR, error)
	GetByName(name string) (*PR, error)
	// GetByTeamId(id int) (*PR, error)
	Create(team *PR) (int, error)
	Update(team *PR) error
	Delete(id int) error
}
