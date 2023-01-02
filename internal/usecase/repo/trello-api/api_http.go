package trello_api

import (
	"context"
	"fmt"
	"github.com/adlio/trello"
	"github.com/ferralucho/go-task-management/config"
	"github.com/ferralucho/go-task-management/internal/entity"
	"strings"
)

type TrelloApi struct {
	Client   *trello.Client
	Username string
	Board    string
}

// New -.
func New(cfg *config.Config) *TrelloApi {
	c := trello.NewClient(cfg.Trello.PublicKey, cfg.Trello.MemberToken)

	return &TrelloApi{
		Client:   c,
		Username: cfg.Trello.Username,
		Board:    cfg.Trello.MemberBoard,
	}
}

// CreateCard -.
func (r *TrelloApi) CreateCard(ctx context.Context, card entity.InternalCard) (entity.Card, error) {
	member, err := r.Client.GetMember(r.Username, trello.Defaults())
	if err != nil {
		return entity.Card{}, fmt.Errorf("CardUseCase - CreateCard - Board GetMember: %w", err)
	}

	boards, err := member.GetBoards(trello.Defaults())
	if err != nil {
		return entity.Card{}, fmt.Errorf("CardUseCase - CreateCard - Board GetBoards: %w", err)
	}
	var board *trello.Board
	for _, v := range boards {
		if v.Name == r.Board {
			board = v
		}
	}

	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		return entity.Card{}, fmt.Errorf("CardUseCase - CreateCard - Board GetLists: %w", err)
	}

	var list *trello.List
	for _, v := range lists {
		if strings.ToLower(v.Name) == "to do" {
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
		if e := list.AddCard(ca, trello.Defaults()); e != nil {
			return entity.Card{}, fmt.Errorf("CardUseCase - CreateCard - Board GetLists: %w", e)
		}
	}

	externalCard := entity.Card{
		Name: card.Name, Desc: card.Desc, IDLabels: card.IdLabels, IDMembers: []string{memberToAssign}, IDList: list.ID,
	}

	return externalCard, nil
}
