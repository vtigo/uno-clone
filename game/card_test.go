package game

import (
	"testing"
)

func TestCreateDeck(t *testing.T) {
	deck := CreateDeck()
	
	if len(deck.Cards) != 108 {
		t.Errorf("Expected deck to have 108 cards, got %d", len(deck.Cards))
	}

	typeCounts := make(map[CardType]int)
	colorCounts := make(map[CardColor]int)

	for _, card := range deck.Cards {
		typeCounts[card.Type] ++
		colorCounts[card.Color] ++
	}

	if typeCounts[Number] != 76 {
		t.Errorf("Expected 76 Number cards, got %d", typeCounts[Number])
	}

	if typeCounts[Skip] != 8 {
		t.Errorf("Expected 8 Skip cards, got %d", typeCounts[Skip])
	}

	if typeCounts[Reverse] != 8 {
		t.Errorf("Expected 8 Reverse cards, got %d", typeCounts[Reverse])
	}

	if typeCounts[DrawTwo] != 8 {
		t.Errorf("Expected 8 Draw Two cards, got %d", typeCounts[DrawTwo])
	}

	if typeCounts[WildCard] != 4 {
		t.Errorf("Expected 4 Wild cards, got %d", typeCounts[WildCard])
	}

	if typeCounts[WildDrawFour] != 4 {
		t.Errorf("Expected 4 Wild Draw Four cards, got %d", typeCounts[WildDrawFour])
	}

	for color:= Red; color <= Yellow; color++ {
		if colorCounts[color] != 25 {
			t.Errorf("Expected 25 %s cards, got %d", CardColor(color).String(), colorCounts[color])
		}
	}

	if colorCounts[Wild] != 8 {
		t.Errorf("Expected 8 Wild colored cards, got %d", colorCounts[Wild])
	}
}

func TestDrawCard(t *testing.T) {
	deck := CreateDeck()
	initialCount := deck.Count()

	card := deck.DrawCard()

	if card == nil {
		t.Errorf("Expected to draw a card, got nil")
	}

	if deck.Count() != initialCount-1 {
		t.Errorf("Expected deck to have %d cards after drawing, got %d", initialCount-1, deck.Count())
	}

	for range initialCount-1 {
		deck.DrawCard()
	}

	emptyCard := deck.DrawCard()
	if emptyCard != nil {
		t.Error("Expected nil when drawing from empty deck, got a card")
	}
}
