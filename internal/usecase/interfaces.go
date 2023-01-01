package usecase

import (
	"context"

	"github.com/ferralucho/go-task-management/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Card -.
	Card interface {
		CreateTask(context.Context, entity.Task) (entity.Card, error)
		CreateBug(context.Context, entity.Bug) (entity.Card, error)
		CreateIssue(context.Context, entity.Issue) (entity.Card, error)
	}

	// CardRepo -.
	CardRepo interface {
		CreateCard(context.Context, entity.InternalCard) (entity.Card, error)
	}
)
