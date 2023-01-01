package usecase

import (
	"context"
	"fmt"
	"github.com/ferralucho/go-task-management/internal/entity"
)

// CardUseCase -.
type CardUseCase struct {
	repo CardRepo
}

// New -.
func New(r CardRepo) *CardUseCase {
	return &CardUseCase{
		repo: r,
	}
}

func (uc *CardUseCase) CreateTask(ctx context.Context, task entity.Task) (entity.Card, error) {
	c, err := uc.convertToCard(task.Title, task.Category)
	card, err := uc.repo.CreateCard(ctx, c)
	if err != nil {
		return entity.Card{}, fmt.Errorf("CardUseCase - CreateTask - s.repo.CreateCard: %w", err)
	}

	return card, nil
}

func (uc *CardUseCase) CreateIssue(ctx context.Context, issue entity.Issue) (entity.Card, error) {
	c, err := uc.convertToCard(issue.Title, issue.Description)
	card, err := uc.repo.CreateCard(ctx, c)
	if err != nil {
		return entity.Card{}, fmt.Errorf("CardUseCase - CreateIssue - s.repo.CreateCard: %w", err)
	}

	return card, nil
}

func (uc *CardUseCase) CreateBug(ctx context.Context, bug entity.Bug) (entity.Card, error) {
	c, err := uc.convertToCard(bug.Title, bug.Description)
	card, err := uc.repo.CreateCard(ctx, c)
	if err != nil {
		return entity.Card{}, fmt.Errorf("CardUseCase - CreateBug - s.repo.CreateCard: %w", err)
	}

	return card, nil
}

func (uc *CardUseCase) convertToCard(name string, desc string) (entity.InternalCard, error) {
	return entity.InternalCard{
		Name: name,
		Desc: desc,
	}, nil
}
