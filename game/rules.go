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

func (gr *GameRules) ValidateMove(player *Player, cardIndex int, state *GameState) (bool, string) {
	if !player.IsMyTurn {
		return false, "It's not your turn"
	}

	if cardIndex < 0 || cardIndex >= len(player.Hand) {
		return false, "Invalid card index"
	} 

	if state.Phase != PhasePlay {
		return false, "Game is not in the play phase"
	}

	card := player.Hand[cardIndex]

	if len(state.DiscardPile.Cards) == 0 {
		return false, "Discard pile is empty"
	}

	topCard := state.DiscardPile.Cards[len(state.DiscardPile.Cards)-1]

	if !card.CanPlayOn(topCard, state.ActiveColor) {
		return false, "Card cannot be played on top of the current discard pile"
	}

	if card.Color == Wild && card.Type == WildDrawFour {
		if !IsWildDrawFourValid(player.Hand, state.ActiveColor) {
			return false, "Wild Draw Four can only be played if you don't have any cards of the active color"
		}
	}

	return true, ""
}

func (gr *GameRules) HandleCardEffect(card *Card, state *GameState, chosenColor *CardColor) error {
	if state.Phase != PhasePlay && state.Phase != PhaseColorSelection {
		return errors.New("game is not in the play or color selection phase")
	}

	switch card.Type {
	case Number:
		return gr.handleNumberCard(state)
	case Skip:
		return gr.handleSkipCard(state)
	case Reverse:
		return gr.handleReverseCard(state)
	case DrawTwo:
		return gr.handleDrawTwoCard(state)
	case WildCard:
		if chosenColor == nil {
			state.Phase = PhaseColorSelection
			return nil
		}
		// return gr.handleWildCard(card, state, *chosenColor)
	case WildDrawFour:
		if chosenColor == nil {
			state.Phase = PhaseColorSelection
			return nil
		}
		// return gr.handleWildDrawFourCard(card, state, *chosenColor)
	default:
		return fmt.Errorf("unknown card type: %v", card.Type)
	}
	return nil
}

func (gr *GameRules) handleNumberCard(state *GameState) error {
	// Number cards have no special effects
	gr.NextTurn(state)
	return nil
}

func (gr *GameRules) handleSkipCard(state *GameState) error {
	gr.SkipTurn(state)
	return nil
}

func (gr *GameRules) handleReverseCard(state *GameState) error {
	gr.ReverseTurn(state)
	return nil
}

func (gr *GameRules) handleDrawTwoCard(state *GameState) error {
	playerToDraw := (state.CurrentPlayer + 1) % len(state.Players)

	cardsDrawn, err := state.DrawPile.DrawN(2)
	if err != nil {
		if err.Error() == "not enough cards in deck" {
			// TODO: If there are cards in the discard pile, shuffle them into the draw pile
			// Keep the top card in the discard pile
			// Add the cards from the discard pile to the draw pile
			// Shuffle the draw pile
			// Try to draw again
		} else {
			return fmt.Errorf("failed to draw cards: %v", err)
		}
	}

	cardPtrs := make([]*Card, len(cardsDrawn))
	for i:= range cardPtrs {
		cardPtrs[i] = &cardsDrawn[i]
	}
	state.Players[playerToDraw].AddCardsToHand(cardPtrs)
	
	gr.SkipTurn(state)
	return nil
}

func (gr *GameRules) NextTurn(state *GameState) {
	state.Players[state.CurrentPlayer].IsMyTurn = false
	state.CurrentPlayer = (state.CurrentPlayer + 1) % len(state.Players)
	state.Players[state.CurrentPlayer].IsMyTurn = true
}

func (gr *GameRules) SkipTurn(state *GameState) {
	// In a two player game, skipping means staying with the same current player
	// Can be implemented in the future to handle more than two players
}

func (gr *GameRules) RepeatTurn(state *GameState) {
	// In a two player game, repeating means staying with the same current player
	// Can be implemented in the future to handle more than two players
}

func (gr *GameRules) ReverseTurn(state *GameState) {
	// In a two player game, reversing means staying with the same current player
	// Can be implemented in the future to handle more than two players
}
