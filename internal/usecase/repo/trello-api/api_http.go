package trello_api

import (
	"context"
	"fmt"
	"github.com/adlio/trello"
	"github.com/ferralucho/go-task-management/internal/entity"
	"os"
)

type TrelloApi struct {
	Client *trello.Client
}

// New -.
func New() *TrelloApi {
	appKey := os.Getenv("TRELLO_DEVELOPER_PUBLIC_KEY")
	token := os.Getenv("TRELLO_MEMBER_TOKEN")
	c := trello.NewClient(appKey, token)

	return &TrelloApi{
		Client: c,
	}
}

// CreateCard -.
func (r *TrelloApi) CreateCard(ctx context.Context, card entity.InternalCard) (entity.Card, error) {
	member, err := r.Client.GetMember(os.Getenv("TRELLO_USERNAME"), trello.Defaults())
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
