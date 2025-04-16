package game

import (
	"errors"
	"fmt"
)

type GamePhase int

const (
	PhaseSetup GamePhase = iota
	PhasePlay
	PhaseColorSelection
	PhaseGameOver
)

type GameState struct {
	Players       []*Player
	CurrentPlayer int
	DrawPile      *Deck
	DiscardPile   *Deck
	ActiveColor   CardColor
	Phase         GamePhase
	LastPlayedBy  int
}

type GameRules struct {}

func NewGameRules() *GameRules {
	return &GameRules{}
}

func NewGameState(players []*Player) (*GameState, error) {
	if len(players) != 2 {
		return nil, errors.New("two players are required")
	}

	state := &GameState{
		Players:       players,
		CurrentPlayer: 0,
		Phase:         PhaseSetup,
		LastPlayedBy:  -1,
	}

	deck := NewDeck()
	deck.Shuffle()

	for _, player := range players {
		cards, err := deck.DrawN(7)
		if err != nil {
			return nil, fmt.Errorf("Failed to deal initial cards %v", err)
		}

		cardPtrs := make([]*Card, len(cards))
		for i := range cards {
			cardPtrs[i] = &cards[i]
		}
		player.AddCardsToHand(cardPtrs)
	}

	initialCard, err := deck.Draw()
	if err != nil {
		return nil, fmt.Errorf("failed to draw initial card: %v", err)
	}

	state.DiscardPile = CreateDiscardPile(initialCard)
	state.DrawPile = deck

	if initialCard.Color == Wild {
		state.ActiveColor = Red
	} else {
		state.ActiveColor = initialCard.Color
	}

	if initialCard.Type != Number {
		switch initialCard.Type {
		case Skip, Reverse:
			// Skip or Reverse as initial card: First player gets another turn
			// Nothing to do here as current player is already 0
		case DrawTwo:
			// Draw Two as initial card: Second player draws 2 cards and first player takes a turn
			secondPlayer := (state.CurrentPlayer + 1) % len(state.Players)
			cardsDrawn, err := state.DrawPile.DrawN(2)
			if err != nil {
				return nil, fmt.Errorf("failed to draw cards for initial Draw Two: %v", err)
			}
			cardPtrs := make([]*Card, len(cardsDrawn))
			for i := range cardsDrawn {
				cardPtrs[i] = &cardsDrawn[i]
			}
			state.Players[secondPlayer].AddCardsToHand(cardPtrs)
		case WildCard:
			// Wild card as initial card: First player chooses the color
			// Default to Red until player chooses
			state.ActiveColor = Red
			state.Phase = PhaseColorSelection
		case WildDrawFour:
			// Wild Draw Four as initial card: Second player draws 4 cards, first player chooses color
			secondPlayer := (state.CurrentPlayer + 1) % len(state.Players)
			cardsDrawn, err := state.DrawPile.DrawN(4)
			if err != nil {
				return nil, fmt.Errorf("failed to draw cards for initial Wild Draw Four: %v", err)
			}
			cardPtrs := make([]*Card, len(cardsDrawn))
			for i := range cardsDrawn {
				cardPtrs[i] = &cardsDrawn[i]
			}
			state.Players[secondPlayer].AddCardsToHand(cardPtrs)
			state.ActiveColor = Red
			state.Phase = PhaseColorSelection
		}
	}

	if state.Phase != PhaseColorSelection {
		state.Phase = PhasePlay
	}

	state.Players[state.CurrentPlayer].IsMyTurn = true

	return state, nil
}
