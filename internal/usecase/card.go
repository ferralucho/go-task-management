package usecase

import (
	"context"
	"fmt"
	"github.com/ferralucho/go-task-management/internal/entity"
	"math/rand"
	"strconv"
	"strings"
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
	c, err := uc.convertToCard(task.Title, task.Category, []string{task.Category}, true)
	card, err := uc.repo.CreateCard(ctx, c)
	if err != nil {
		return entity.Card{}, fmt.Errorf("CardUseCase - CreateTask - s.repo.CreateCard: %w", err)
	}

	return card, nil
}

func (uc *CardUseCase) CreateIssue(ctx context.Context, issue entity.Issue) (entity.Card, error) {
	c, err := uc.convertToCard(issue.Title, issue.Description, []string{}, false)
	card, err := uc.repo.CreateCard(ctx, c)
	if err != nil {
		return entity.Card{}, fmt.Errorf("CardUseCase - CreateIssue - s.repo.CreateCard: %w", err)
	}

	return card, nil
}

func (uc *CardUseCase) CreateBug(ctx context.Context, bug entity.Bug) (entity.Card, error) {
	descriptionTitle := "bug-"
	if bug.Description == "" {
		descriptionTitle += "space-"
	} else {
		i := strings.Index(bug.Description, " ")
		if i != -1 {
			descriptionTitle += bug.Description[:i] + "-"
		}
	}

	descriptionTitle += strconv.Itoa(rangeIn(0, 1000))

	c, err := uc.convertToCard(descriptionTitle, bug.Description, []string{"Bug"}, true)
	card, err := uc.repo.CreateCard(ctx, c)
	if err != nil {
		return entity.Card{}, fmt.Errorf("CardUseCase - CreateBug - s.repo.CreateCard: %w", err)
	}

	return card, nil
}

func (uc *CardUseCase) convertToCard(name string, desc string, labels []string, assign bool) (entity.InternalCard, error) {
	return entity.InternalCard{
		Name:     name,
		Desc:     desc,
		IdLabels: labels,
		Assign:   assign,
	}, nil
}

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}
