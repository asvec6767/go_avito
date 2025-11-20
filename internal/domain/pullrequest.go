package domain

type PullRequestStatus string

const (
	PullRequestStatusOpen   PullRequestStatus = "OPEN"
	PullRequestStatusMerged PullRequestStatus = "MERGED"
)

type PR struct {
	ID        int
	Name      string
	Author    *User
	Status    PullRequestStatus
	Reviewers []*User
}

type PRRepository interface {
	GetById(id int) (*PR, error)
	GetByName(name string) (*PR, error)
	Create(team *PR) (int, error)
	Update(team *PR) error
	Delete(id int) error
}
