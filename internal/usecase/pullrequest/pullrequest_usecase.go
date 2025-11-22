package pullrequest

import (
	"main/internal/domain"
	"math/rand/v2"
	"slices"
)

type pullRequestUseCase struct {
	pr   domain.PRRepository
	user domain.UserRepository
	team domain.TeamRepository
}

func NewPullRequestUseCase(prRepo domain.PRRepository, userRepo domain.UserRepository, teamRepo domain.TeamRepository) *pullRequestUseCase {
	return &pullRequestUseCase{
		pr:   prRepo,
		user: userRepo,
		team: teamRepo,
	}
}

func (uc *pullRequestUseCase) Create(name string, author_id string) (*domain.PR, error) {
	user, err := uc.user.GetById(author_id)
	if err != nil {
		return nil, err
	}

	team, err := uc.team.GetById(user.TeamID)
	if err != nil {
		return nil, err
	}

	pr := &domain.PR{
		Name:     name,
		Status:   domain.PullRequestStatusOpen,
		AuthorID: author_id,
		TeamID:   team.ID,
	}

	if err := uc.pr.Create(pr); err != nil {
		return nil, err
	}

	return pr, nil
}

func (uc *pullRequestUseCase) GetById(id string) (*domain.PR, error) {
	return uc.pr.GetById(id)
}

// func (uc *pullRequestUseCase) GetByName(name string) (*domain.PR, error) {
// 	return uc.pr.GetByName(name)
// }

// TODO: сделать норм ошибку
func (uc *pullRequestUseCase) ChangeAllReviewers(id string) (*domain.PR, error) {
	pullrequest, err := uc.pr.GetById(id)
	if err != nil {
		return nil, err
	}

	if pullrequest.Status == domain.PullRequestStatusMerged {
		return nil, domain.ErrAccessDenied
		// return nil, fmt.Errorf("невозможно изменить pullrequest в статусе merged")
	}

	author, err := uc.user.GetById(pullrequest.AuthorID)
	if err != nil {
		return nil, err
	}

	team, err := uc.team.GetById(author.TeamID)
	if err != nil {
		return nil, err
	}

	users := team.Users

	users = removeUserFromUsersList(users, author)
	pullrequest.Reviewers = setNRandomReviewers(users, 2)

	return pullrequest, nil
}

func removeUserFromUsersList(users []*domain.User, user *domain.User) []*domain.User {
	idx := slices.IndexFunc(users, func(u *domain.User) bool { return u.ID == user.ID })

	new_slice := make([]*domain.User, 0)
	new_slice = append(new_slice, users[:idx]...)
	return append(new_slice, users[idx+1:]...)
}

func addRandomReviewer(reviewers []*domain.User, users []*domain.User) []*domain.User {
	rand_user := users[rand.IntN(len(users))]

	return append(reviewers, rand_user)
}

func setNRandomReviewers(users []*domain.User, count int) []*domain.User {
	var reviewers []*domain.User
	for range count {
		reviewers = addRandomReviewer(reviewers, users)
		users = removeUserFromUsersList(users, reviewers[len(reviewers)-1])
	}
	return reviewers
}

// TODO: сделать норм ошибку
// TODO: сделать чтобы заменялся один reviewer а не сразу все
func (uc *pullRequestUseCase) ChangeReviewer(pr_id, old_reviewer_id string) (*domain.PR, error) {
	pullrequest, err := uc.pr.GetById(pr_id)
	if err != nil {
		return nil, err
	}

	if pullrequest.Status == domain.PullRequestStatusMerged {
		return nil, domain.ErrAccessDenied
		// return nil, fmt.Errorf("невозможно изменить pullrequest в статусе merged")
	}

	old_reviewer, err := uc.user.GetById(old_reviewer_id)
	if err != nil {
		return nil, err
	}

	old_true_reviewer := removeUserFromUsersList(pullrequest.Reviewers, old_reviewer)[0]

	team, err := uc.team.GetById(old_reviewer.TeamID)
	if err != nil {
		return nil, err
	}

	users := team.Users

	author := pullrequest.Author

	users = removeUserFromUsersList(users, &author)
	users = removeUserFromUsersList(users, old_reviewer)
	pullrequest.Reviewers = setNRandomReviewers(users, 1)

	pullrequest.Reviewers = append(pullrequest.Reviewers, old_true_reviewer)

	return pullrequest, nil
}

func (uc *pullRequestUseCase) SetMergedStatus(pr_id string) (*domain.PR, error) {
	pullrequest, err := uc.pr.GetById(pr_id)
	if err != nil {
		return nil, err
	}

	pullrequest.Status = domain.PullRequestStatusMerged

	err = uc.pr.Update(pullrequest)
	if err != nil {
		return nil, err
	}

	return pullrequest, nil
}

func (uc *pullRequestUseCase) SetOpenStatus(pr_id string) (*domain.PR, error) {
	pullrequest, err := uc.pr.GetById(pr_id)
	if err != nil {
		return nil, err
	}

	pullrequest.Status = domain.PullRequestStatusOpen

	err = uc.pr.Update(pullrequest)
	if err != nil {
		return nil, err
	}

	return pullrequest, nil
}

func (uc *pullRequestUseCase) Delete(id string) error {
	return uc.pr.Delete(id)
}
