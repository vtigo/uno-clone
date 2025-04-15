package game

import (
	"math/rand"
	"time"
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

type Card struct{
	Color	CardColor
	Type	CardType
	Value 	int // Only used for number cards (0-9)
}

type Deck struct{
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

func CreateDeck() *Deck {
	deck := &Deck{Cards: make([]Card, 0, 108)}

	for color := Red; color <= Yellow; color++ {
		deck.Cards = append(deck.Cards, Card{Color: color, Type: Number, Value: 0})

		for i := 1; i <= 9; i++ {
			deck.Cards = append(deck.Cards, Card{Color: color, Type: Number, Value: i})
			deck.Cards = append(deck.Cards, Card{Color: color, Type: Number, Value: i})
		}

		for range 2 {
			deck.Cards = append(deck.Cards, Card{Color: color, Type: Skip})
			deck.Cards = append(deck.Cards, Card{Color: color, Type: Reverse})
			deck.Cards = append(deck.Cards, Card{Color: color, Type: DrawTwo})
		}
	}

	for range 4 {
		deck.Cards = append(deck.Cards, Card{Color: Wild, Type: WildCard})
		deck.Cards = append(deck.Cards, Card{Color: Wild, Type: WildDrawFour})
	}
	
	return deck
}

func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	// Fisher-Yates shuffle algorithm
	for i := len(d.Cards) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *Deck) DrawCard() *Card {
	if len(d.Cards) == 0 {
		return nil
	}

	card := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]

	return &card
}

func (d *Deck) Count() int {
	return len(d.Cards)
}
