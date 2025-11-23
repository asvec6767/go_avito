package pullrequest

import (
	"context"
	"main/internal/domain"
	"math/rand/v2"
	"slices"
	"time"
)

const reviewers_count_default = 2 //кол-во ревьюверов на пулл реквесте

type pullRequestUseCase struct {
	pr   domain.PRRepository
	user domain.UserRepository
	team domain.TeamRepository
}

type CreatePRRequest struct {
	ID       string
	Name     string
	AuthorID string
}

func NewPullRequestUseCase(prRepo domain.PRRepository, userRepo domain.UserRepository, teamRepo domain.TeamRepository) *pullRequestUseCase {
	return &pullRequestUseCase{
		pr:   prRepo,
		user: userRepo,
		team: teamRepo,
	}
}

func (uc *pullRequestUseCase) Create(ctx context.Context, req *CreatePRRequest) (*domain.PR, error) {
	user, err := uc.user.GetById(ctx, req.AuthorID)
	if err != nil {
		return nil, err
	}

	if !user.IsActive {
		return nil, domain.ErrUserNotActive
	}

	team, err := uc.team.GetById(ctx, user.TeamID)
	if err != nil {
		return nil, err
	}

	pr := &domain.PR{
		ID:       req.ID,
		Name:     req.Name,
		Status:   domain.PullRequestStatusOpen,
		AuthorID: req.AuthorID,
		TeamID:   team.ID,
	}

	if err := uc.pr.Create(ctx, pr); err != nil {
		return nil, err
	}

	pr, err = uc.ChangeAllReviewers(ctx, pr.ID)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

func (uc *pullRequestUseCase) GetById(ctx context.Context, id string) (*domain.PR, error) {
	return uc.pr.GetById(ctx, id)
}

func (uc *pullRequestUseCase) GetPRWithReviewers(ctx context.Context, id string) (*domain.PR, error) {
	pr, err := uc.pr.GetWithReviewers(ctx, id)
	if err != nil {
		return nil, err
	}
	return pr, nil
}

func (uc *pullRequestUseCase) ChangeReviewer(ctx context.Context, pr_id, old_reviewer_id string) (*domain.PR, *domain.User, error) {
	pr, err := uc.pr.GetById(ctx, pr_id)
	if err != nil {
		return nil, nil, err
	}

	if pr.MergedAt != nil || pr.Status != domain.PullRequestStatusOpen {
		return nil, nil, domain.ErrPRAlreadyMerged
	}

	old_reviewers_ids := pr.ReviewerIDs

	err = uc.assignReviewers(ctx, pr_id, append(old_reviewers_ids, pr.AuthorID), 1)
	if err != nil {
		return nil, nil, err
	}

	pr, err = uc.pr.GetById(ctx, pr_id)
	if err != nil {
		return nil, nil, err
	}

	id_reviewer_replaced_by := ""
	if len(pr.ReviewerIDs) >= 1 {
		id_reviewer_replaced_by = pr.ReviewerIDs[0]
	}

	for _, id := range old_reviewers_ids {
		if id != old_reviewer_id {
			pr.ReviewerIDs = append(pr.ReviewerIDs, id)
		}
	}

	reviewer_replaced_by, err := uc.user.GetById(ctx, id_reviewer_replaced_by)
	if err != nil {
		return nil, nil, err
	}

	return pr, reviewer_replaced_by, nil
}

func (uc *pullRequestUseCase) ChangeAllReviewers(ctx context.Context, pr_id string) (*domain.PR, error) {
	pr, err := uc.pr.GetById(ctx, pr_id)
	if err != nil {
		return nil, err
	}

	if pr.MergedAt != nil || pr.Status != domain.PullRequestStatusOpen {
		return nil, domain.ErrPRAlreadyMerged
	}

	err = uc.assignReviewers(ctx, pr_id, []string{pr.AuthorID}, reviewers_count_default)
	if err != nil {
		return nil, err
	}

	pr, err = uc.pr.GetById(ctx, pr_id)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

func (uc *pullRequestUseCase) assignReviewers(ctx context.Context, pr_id string, notAvailableReviewerIDs []string, reviewers_count_default int) error {
	pr, err := uc.pr.GetById(ctx, pr_id)
	if err != nil {
		return err
	}

	teamUsers, err := uc.user.GetByActiveAndTeam(ctx, pr.TeamID)
	if err != nil {
		return err
	}

	var availableReviewers []domain.User
	for _, user := range teamUsers {
		if idx := slices.IndexFunc(notAvailableReviewerIDs, func(c string) bool { return c == user.ID }); idx < 0 {
			availableReviewers = append(availableReviewers, user)
		}
	}

	reviewerIDs := uc.selectReviewers(availableReviewers, min(len(availableReviewers), reviewers_count_default))

	pr.ReviewerIDs = reviewerIDs
	if err := uc.pr.Update(ctx, pr); err != nil {
		return err
	}

	return nil
}

func (uc *pullRequestUseCase) selectReviewers(users []domain.User, count int) []string {
	if len(users) <= count {
		ids := make([]string, len(users))
		for i, user := range users {
			ids[i] = user.ID
		}
		return ids
	}

	rand.Shuffle(len(users), func(i, j int) {
		users[i], users[j] = users[j], users[i]
	})

	reviewer_ids := make([]string, count)
	for i := 0; i < count; i++ {
		reviewer_ids[i] = users[i].ID
	}

	return reviewer_ids
}

func (uc *pullRequestUseCase) SetMergedStatus(ctx context.Context, pr_id string) (*domain.PR, error) {
	pullrequest, err := uc.pr.GetById(ctx, pr_id)
	if err != nil {
		return nil, err
	}

	if pullrequest.Status != domain.PullRequestStatusOpen {
		return nil, domain.ErrPRAlreadyMerged
	}

	if pullrequest.MergedAt != nil {
		now := time.Now()
		pullrequest.Status = domain.PullRequestStatusMerged
		pullrequest.MergedAt = &now
	}

	if err := uc.pr.Update(ctx, pullrequest); err != nil {
		return nil, err
	}

	return pullrequest, nil
}

func (uc *pullRequestUseCase) SetOpenStatus(ctx context.Context, pr_id string) (*domain.PR, error) {
	pullrequest, err := uc.pr.GetById(ctx, pr_id)
	if err != nil {
		return nil, err
	}

	pullrequest.Status = domain.PullRequestStatusOpen

	err = uc.pr.Update(ctx, pullrequest)
	if err != nil {
		return nil, err
	}

	return pullrequest, nil
}

func (uc *pullRequestUseCase) Delete(ctx context.Context, id string) error {
	return uc.pr.Delete(ctx, id)
}

func (uc *pullRequestUseCase) GetListByUserId(ctx context.Context, user_id string) ([]domain.PR, error) {
	return uc.pr.GetByReviewerAndStatus(ctx, user_id, domain.PullRequestStatusOpen)
}
