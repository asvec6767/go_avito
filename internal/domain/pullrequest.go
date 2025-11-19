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
