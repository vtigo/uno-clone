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
		LastPlayedBy:  -1, // Nobody played yet
	}
	
	// Create and shuffle deck
	deck := NewDeck()
	deck.Shuffle()
	
	// Draw initial hands
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
	
	// Draw initial card
	initialCard, err := deck.Draw()
	if err != nil {
		return nil, fmt.Errorf("failed to draw initial card: %v", err)
	}
	
	// Create the discard pile with the initial card as the first card and "put the deck on the draw pile"
	state.DiscardPile = CreateDiscardPile(initialCard)
	state.DrawPile = deck
	
	// Set the current color based on the initial card
	if initialCard.Color == Wild {
		state.ActiveColor = Red
	} else {
		state.ActiveColor = initialCard.Color
	}
	
	// Handle initial card effects
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
	
	// Start play phase
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
		return gr.handleWildCard(state, *chosenColor)
	case WildDrawFour:
		if chosenColor == nil {
			state.Phase = PhaseColorSelection
			return nil
		}
		return gr.handleWildDrawFourCard(state, *chosenColor)
	default:
		return fmt.Errorf("unknown card type: %v", card.Type)
	}
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

func (gr * GameRules) handleWildCard(state *GameState, chosenColor CardColor) error {
	if chosenColor < Red || chosenColor > Yellow {
		return errors.New("invalid color choice for Wild Card")
	}

	state.ActiveColor = chosenColor

	gr.NextTurn(state)
	return nil
}

func (gr *GameRules) handleWildDrawFourCard(state *GameState, chosenColor CardColor) error {
	if chosenColor < Red || chosenColor > Yellow {
		return errors.New("invalid color choice for Wild Draw Four Card")
	}

	state.ActiveColor = chosenColor

	playerToDraw := (state.CurrentPlayer + 1) % len(state.Players)

	cardsDrawn, err := state.DrawPile.DrawN(4)
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
	for i:= range cardsDrawn {
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

func (gr *GameRules) HandleUnoCall(playerIndex int, state *GameState) (bool, string) {
	if playerIndex < 0 || playerIndex >= len(state.Players) {
		return false, "Invalid player index"
	}

	player := state.Players[playerIndex]

	if !player.ShouldCallUno() {
		return false, "Player does not have exactly one card left"
	}

	player.CallUno()
	return true, "UNO called successfully"
}

func (gr *GameRules) HandleUnoChallenge(targetIndex int, state *GameState) (bool, string) {
	if targetIndex < 0 || targetIndex > len(state.Players) {
		return false, "Invalid target index"
	}

	target := state.Players[targetIndex]

	if !target.ShouldCallUno() {
		return false, "Target does not need to call uno"
	}

	if target.HasCalledUno {
		return false, "Target allready called uno"
	}

	cardsDrawn, err := state.DrawPile.DrawN(2)
	if err != nil {
		if err.Error() == "not enough cards in deck" {
			// TODO: If there are cards in the discard pile, shuffle them into the draw pile
			// Keep the top card in the discard pile
			// Add the cards from the discard pile to the draw pile
			// Shuffle the draw pile
			// Try to draw again
		} else {
			return false, fmt.Sprintf("failed to draw cards: %v", err) 
		}
	}

	cardPtrs := make([]*Card, len(cardsDrawn))
	for i := range cardsDrawn {
		cardPtrs[i] = &cardsDrawn[i]
	}

	target.AddCardsToHand(cardPtrs)

	return true, "Challenge successfull! Target has drawn 2 cards"
}

func (gr *GameRules) HandlePlayCard(player *Player, cardIndex int, state *GameState, chosenColor *CardColor) error {
	valid, message := gr.ValidateMove(player, cardIndex, state)
	if(!valid) {
		return errors.New(message)
	}

	card, err := player.PlayCard(cardIndex)
	if err != nil {
		return fmt.Errorf("failed to play card: %v", err)
	}

	state.DiscardPile.AddToBottom(*card)

	state.LastPlayedBy = state.CurrentPlayer

	if chosenColor != nil || card.Color != Wild {
		err = gr.HandleCardEffect(card, state, chosenColor)
		if err != nil {
			return fmt.Errorf("failed to handle card effect: %v", err)
		}
	} else {
		state.Phase = PhaseColorSelection
	}

	if player.HasWon() {
		state.Phase = PhaseGameOver
	}

	return nil
}

func (gr *GameRules) HandleDrawCard(player *Player, state *GameState) error {
	if !player.IsMyTurn {
		return errors.New("it is not your turn")
	}

	if state.Phase != PhasePlay {
		return errors.New("game is not in the play phase")
	}

	card, err := state.DrawPile.Draw()
	if err != nil {
		// Handle the case where there aren't enough cards in the draw pile
		if err.Error() == "cannot draw from an empty deck" {
			// If there are cards in the discard pile, shuffle them into the draw pile
			if len(state.DiscardPile.Cards) > 1 { // Keep the top card in the discard pile
				topCard := state.DiscardPile.Cards[len(state.DiscardPile.Cards)-1]
				state.DiscardPile.Cards = state.DiscardPile.Cards[:len(state.DiscardPile.Cards)-1]

				// Add the cards from the discard pile to the draw pile
				state.DrawPile.Cards = append(state.DrawPile.Cards, state.DiscardPile.Cards...)
				state.DiscardPile.Cards = []Card{topCard}

				// Shuffle the draw pile
				state.DrawPile.Shuffle()

				// Try to draw again
				card, err = state.DrawPile.Draw()
				if err != nil {
					return fmt.Errorf("failed to draw card after reshuffling: %v", err)
				}
			} else {
				// Not enough cards even after reshuffling
				return errors.New("no more cards to draw")
			}
		} else {
			return fmt.Errorf("failed to draw card: %v", err)
		}
	}

	player.AddCard(&card)
	return nil
}

func (gr *GameRules) EndTurn(state *GameState) error {
	if state.Phase != PhasePlay {
		return errors.New("game phase is not play phase")
	}

	gr.NextTurn(state)
	return nil
}
