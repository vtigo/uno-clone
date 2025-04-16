package game

import (
	"crypto/rand"
	"errors"
	"math/big"
)

type CardColor int
const (
	Red CardColor = iota
	Blue
	Green
	Yellow
	Wild
)

type CardType int
const (
	Number CardType = iota
	Skip
	Reverse
	DrawTwo
	WildCard
	WildDrawFour
)

// Card represents a UNO card with a color, type, and value
type Card struct {
	Color	CardColor
	Type	CardType
	Value 	int // Only used for number cards (0-9)
}

// Deck represents a collection of UNO cards with thread-safe operations
type Deck struct {
	Cards []Card
}

func (c CardColor) String() string {
	switch(c) {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	case Wild:
		return "Wild"
	default:
		return "Unknown"
	}
}

func (t CardType) String() string {
	switch t {
	case Number:
		return "Number"
	case Skip:
		return "Skip"
	case Reverse:
		return "Reverse"
	case DrawTwo:
		return "Draw Two"
	case WildCard:
		return "Wild"
	case WildDrawFour:
		return "Wild Draw Four"
	default:
		return "Unknown"
	}
}

func (c Card) String() string {
	if c.Type == Number {
		return c.Color.String() + " " + c.Type.String() + " " + string(rune('0'+c.Value))
	}
	return c.Color.String() + " " + c.Type.String()
}

func (c Card) CanPlayOn(topCard Card, activeColor CardColor) bool {
	// Wild cards and Wild Draw Four cards can always be played
	if c.Color == Wild {
		return true
	}
	
	// If the top card is a Wild card, we need to check against the active color
	if topCard.Color == Wild {
		return c.Color == activeColor
	}
	
	// Check if the colors match
	if c.Color == topCard.Color {
		return true
	}
	
	// Check if the types match for action cards (Skip, Reverse, Draw Two)
	// but only for the SAME action card type
	if c.Type == topCard.Type && c.Type != Number {
		return true
	}
	
	// Check if the values match for number cards
	if c.Type == Number && topCard.Type == Number && c.Value == topCard.Value {
		return true
	}
	
	// No match found
	return false
}

func IsWildDrawFourValid(hand []*Card, activeColor CardColor) bool {
	for _, card := range hand {
		if card.Color == activeColor {
			return false
		}
	}

	return true
}

// NewDeck creates a new standard 108-card UNO deck
func NewDeck() *Deck {
	deck := &Deck{Cards: make([]Card, 0, 108)}
	
	// Add number cards (0-9) for each color
	for color := Red; color <= Yellow; color++ {
		// Add one 0 card for each color
		deck.Cards = append(deck.Cards, Card{Color: color, Type: Number, Value: 0})

		// Add two of each number 1-9 for each color
		for i := 1; i <= 9; i++ {
			deck.Cards = append(deck.Cards, Card{Color: color, Type: Number, Value: i})
			deck.Cards = append(deck.Cards, Card{Color: color, Type: Number, Value: i})
		}

		// Add two of each action card (Skip, Reverse, Draw Two) for each color
		for range 2 {
			deck.Cards = append(deck.Cards, Card{Color: color, Type: Skip})
			deck.Cards = append(deck.Cards, Card{Color: color, Type: Reverse})
			deck.Cards = append(deck.Cards, Card{Color: color, Type: DrawTwo})
		}
	}

	// Add Wild cards and Wild Draw Four cards
	for range 4 {
		deck.Cards = append(deck.Cards, Card{Color: Wild, Type: WildCard})
		deck.Cards = append(deck.Cards, Card{Color: Wild, Type: WildDrawFour})
	}
	
	return deck
}

// Shuffle randomizes the order of cards in the deck using a cryptographically secure random source
func (d *Deck) Shuffle() {
	// Fisher-Yates shuffle algorithm with crypto/rand
	for i := len(d.Cards) - 1; i > 0; i-- {
		// Generate a secure random number
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			// If crypto/rand fails, skip this iteration
			continue
		}
		
		j := nBig.Int64()
		d.Cards[i], d.Cards[int(j)] = d.Cards[int(j)], d.Cards[i]
	}
}

// Draw removes and returns the top card from the deck
func (d *Deck) Draw() (Card, error) {
	if len(d.Cards) == 0 {
		return Card{}, errors.New("cannot draw from an empty deck")
	}
	
	card := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	
	return card, nil
}

// DrawN draws n cards from the deck
func (d *Deck) DrawN(n int) ([]Card, error) {
	if n <= 0 {
		return nil, errors.New("cannot draw a non-positive number of cards")
	}
	
	if len(d.Cards) < n {
		return nil, errors.New("not enough cards in deck")
	}
	
	cards := make([]Card, n)
	for i := range n {
		cards[i] = d.Cards[len(d.Cards)-1-i]
	}
	
	// Remove the drawn cards from the deck
	d.Cards = d.Cards[:len(d.Cards)-n]
	
	return cards, nil
}

// AddToBottom adds a card to the bottom of the deck
func (d *Deck) AddToBottom(card Card) {
	// Prepend the card to the slice
	d.Cards = append([]Card{card}, d.Cards...)
}

// IsEmpty checks if the deck is empty
func (d *Deck) IsEmpty() bool {
	return len(d.Cards) == 0
}

// Size returns the number of cards in the deck
func (d *Deck) Size() int {
	return len(d.Cards)
}

// CreateDiscardPile creates a new discard pile with a single card
func CreateDiscardPile(initialCard Card) *Deck {
	discardPile := &Deck{Cards: []Card{initialCard}}
	return discardPile
}
