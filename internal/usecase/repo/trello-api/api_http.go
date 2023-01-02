package trello_api

import (
	"context"
	"fmt"
	"github.com/adlio/trello"
	"github.com/ferralucho/go-task-management/config"
	"github.com/ferralucho/go-task-management/internal/entity"
)

type TrelloApi struct {
	Client   *trello.Client
	Username string
}

// New -.
func New(cfg *config.Config) *TrelloApi {
	c := trello.NewClient(cfg.Trello.PublicKey, cfg.Trello.MemberToken)

	return &TrelloApi{
		Client:   c,
		Username: cfg.Trello.Username,
	}
}

// CreateCard -.
func (r *TrelloApi) CreateCard(ctx context.Context, card entity.InternalCard) (entity.Card, error) {
	member, err := r.Client.GetMember(r.Username, trello.Defaults())
	if err != nil {
		return entity.Card{}, fmt.Errorf("CardUseCase - CreateCard - Board GetLists: %w", err)
	}

	board, err := r.Client.GetBoard("bOaRdID", trello.Defaults())
	if err != nil {
		return entity.Card{}, fmt.Errorf("CardUseCase - CreateCard - Board GetBoard: %w", err)
	}

	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		return entity.Card{}, fmt.Errorf("CardUseCase - CreateCard - Board GetLists: %w", err)
	}

	var list *trello.List
	for _, v := range lists {
		if v.Name == "To Do" {
			list = v
		}
	}

	var memberToAssign string
	if card.Assign {
		memberToAssign = member.ID
	} else {
		memberToAssign = ""
	}

	ca := &trello.Card{Name: card.Name, Desc: card.Desc, IDLabels: card.IdLabels, IDMembers: []string{memberToAssign}}
	if list != nil {
		list.AddCard(ca, trello.Defaults())
	}

	externalCard := entity.Card{
		Name: card.Name, Desc: card.Desc, IDLabels: card.IdLabels, IDMembers: []string{memberToAssign},
	}

	return externalCard, nil
}
