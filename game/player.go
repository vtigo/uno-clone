package game

import (
	"errors"
	"fmt"
)

// Player represents a player in the UNO game
type Player struct {
	Name         string   // The player's display name
	Hand         []*Card  // The player's current card hand
	HasCalledUno bool     // Whether the player has called UNO
	IsMyTurn     bool     // Whether it's currently this player's turn
	hasPlayedCard bool    // Internal tracking for if the player has played at least one card
}

// String returns a string representation of the player
func (p *Player) String() string {
	return fmt.Sprintf("Player %s (%d cards) | Uno(%t)", p.Name, len(p.Hand), p.HasCalledUno)
}

// NewPlayer creates a new player with the given name
func NewPlayer(name string) *Player {
	player := &Player{
		Name: name,
		Hand: make([]*Card, 0),
		HasCalledUno: false,
		IsMyTurn: false,
		hasPlayedCard: false,
	}
	return player
}

// AddCard adds a single card to the player's hand
// the UNO call status is reset to false
func (p *Player) AddCard(card *Card) {
	if card != nil {
		p.Hand = append(p.Hand, card)
		p.ResetUnoCall()
	}
}

// AddCardsToHand adds multiple cards to the player's hand
func (p *Player) AddCardsToHand(cards []*Card) {
	for _, card := range cards {
		p.AddCard(card)
	}
}

// PlayCard removes and returns a card at the specified index
// Returns an error if the index is out of bounds
func (p *Player) PlayCard(index int) (*Card, error) {
	if index < 0 || index >= len(p.Hand) {
		return nil, errors.New("Invalid card index")
	}

	card := p.Hand[index]

	// Remove the card by replacing it with the last card and truncating
	p.Hand[index] = p.Hand[len(p.Hand)-1]
	p.Hand = p.Hand[:len(p.Hand)-1]
	
	if !p.hasPlayedCard {
		p.hasPlayedCard = true
	}

	return card, nil
}

// HasValidPlay checks if the player has any valid moves
// against the top card and current color
func (p *Player) HasValidPlay(topCard *Card, currentColor CardColor) bool {
	return len(p.GetValidPlays(topCard, currentColor)) > 0
}

// GetValidPlays returns indices of valid cards to play
// based on the top card and current color
func (p *Player) GetValidPlays(topCard *Card, currentColor CardColor) []int {
	validPlays := make([]int, 0)
	
	for i, card := range p.Hand {
		// Regular Wild cards can always be played
		if card.Color == Wild && card.Type == WildCard {
			validPlays = append(validPlays, i)
			continue
		}
		
		// Wild Draw Four cards can only be played if the player has no cards
		// matching the current color
		if card.Color == Wild && card.Type == WildDrawFour {
			// Use the existing function to check if Wild Draw Four is valid
			if IsWildDrawFourValid(p.Hand, currentColor) {
				validPlays = append(validPlays, i)
			}
			continue
		}
		
		// If the top card color and current color match, use standard rules
		if topCard.Color == currentColor {
			if card.CanPlayOn(*topCard, currentColor) {
				validPlays = append(validPlays, i)
			}
			continue
		}
		
		// If the top card color and current color differ (after a Wild card),
		// only cards matching the current color are valid
		if card.Color == currentColor {
			validPlays = append(validPlays, i)
		}
	}
	
	return validPlays
}

// HandSize returns the number of cards in the player's hand
func (p *Player) HandSize() int {
	return len(p.Hand)
}

// HasWon checks if the player has won (no cards in hand and has played at least one card)
func (p *Player) HasWon() bool {
	return p.hasPlayedCard && len(p.Hand) == 0
}

// CallUno sets HasCalledUno to true
func (p *Player) CallUno() {
	p.HasCalledUno = true
}

// ResetUnoCall sets HasCalledUno to false
func (p *Player) ResetUnoCall() {
	p.HasCalledUno = false
}

// ShouldCallUno returns true if the player has only one card left after playing
func (p *Player) ShouldCallUno() bool {
	return p.hasPlayedCard && len(p.Hand) == 1
}
