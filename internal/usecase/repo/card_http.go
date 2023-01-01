package repo

import (
	"context"
	"github.com/ferralucho/go-task-management/internal/entity"
)

type CardRepoHttp struct{}

// New -.
func New() *CardRepoHttp {
	return &CardRepoHttp{}
}

// CreateCard -.
func (r *CardRepoHttp) CreateCard(ctx context.Context, card entity.InternalCard) (entity.Card, error) {
	return entity.Card{}, nil
}
