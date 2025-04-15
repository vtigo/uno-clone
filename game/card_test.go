package game

import (
	"testing"
)

func TestCanPlayOn(t *testing.T) {
	// Test cases for playing on a number card
	redFive := Card{Color: Red, Type: Number, Value: 5}
	blueFive := Card{Color: Blue, Type: Number, Value: 5}
	redSeven := Card{Color: Red, Type: Number, Value: 7}
	redSkip := Card{Color: Red, Type: Skip}
	blueSkip := Card{Color: Blue, Type: Skip}
	yellowReverse := Card{Color: Yellow, Type: Reverse}
	wildCard := Card{Color: Wild, Type: WildCard}
	wildDrawFour := Card{Color: Wild, Type: WildDrawFour}
	
	// Case 1: Same color, different number
	if !redSeven.CanPlayOn(redFive, Red) {
		t.Error("Expected to be able to play red 7 on red 5")
	}
	
	// Case 2: Different color, same number
	if !blueFive.CanPlayOn(redFive, Red) {
		t.Error("Expected to be able to play blue 5 on red 5")
	}
	
	// Case 3: Different color, different number
	if redSeven.CanPlayOn(blueFive, Blue) {
		t.Error("Expected not to be able to play red 7 on blue 5")
	}
	
	// Case 4: Same color, action card on number
	if !redSkip.CanPlayOn(redFive, Red) {
		t.Error("Expected to be able to play red Skip on red 5")
	}
	
	// Case 5: Different color, action card on number
	if blueSkip.CanPlayOn(redFive, Red) {
		t.Error("Expected not to be able to play blue Skip on red 5")
	}
	
	// Case 6: Same action type, different color
	if !blueSkip.CanPlayOn(redSkip, Red) {
		t.Error("Expected to be able to play blue Skip on red Skip (same type)")
	}
	
	// Case 7: Wild card can be played on any card
	if !wildCard.CanPlayOn(redFive, Red) {
		t.Error("Expected to be able to play Wild on any card")
	}
	
	if !wildDrawFour.CanPlayOn(yellowReverse, Yellow) {
		t.Error("Expected to be able to play Wild Draw Four on any card")
	}
	
	// Case 8: Playing on a Wild card with active color
	if !redFive.CanPlayOn(wildCard, Red) {
		t.Error("Expected to be able to play red 5 on Wild when active color is Red")
	}
	
	if redFive.CanPlayOn(wildCard, Blue) {
		t.Error("Expected not to be able to play red 5 on Wild when active color is Blue")
	}
}

func TestIsWildDrawFourValid(t *testing.T) {
	// Create a hand with various cards
	redFive := &Card{Color: Red, Type: Number, Value: 5}
	blueSeven := &Card{Color: Blue, Type: Number, Value: 7}
	greenSkip := &Card{Color: Green, Type: Skip}
	wildCard := &Card{Color: Wild, Type: WildCard}
	
	// Test case 1: Hand with no cards of active color
	hand1 := []*Card{blueSeven, greenSkip, wildCard}
	if !IsWildDrawFourValid(hand1, Red) {
		t.Error("Expected Wild Draw Four to be valid when hand has no cards of active color")
	}
	
	// Test case 2: Hand with cards of active color
	hand2 := []*Card{redFive, blueSeven, greenSkip, wildCard}
	if IsWildDrawFourValid(hand2, Red) {
		t.Error("Expected Wild Draw Four to be invalid when hand has cards of active color")
	}
	
	// Test case 3: Empty hand
	var hand3 []*Card
	if !IsWildDrawFourValid(hand3, Red) {
		t.Error("Expected Wild Draw Four to be valid with empty hand")
	}
}

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
