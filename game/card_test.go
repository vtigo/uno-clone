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

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	
	if deck.Size() != 108 {
		t.Errorf("Expected deck to have 108 cards, got %d", deck.Size())
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

func TestDraw(t *testing.T) {
	deck := NewDeck()
	initialSize := deck.Size()
	
	_, err := deck.Draw()
	if err != nil {
		t.Errorf("Expected no error when drawing from non-empty deck, got %v", err)
	}
	
	if deck.Size() != initialSize-1 {
		t.Errorf("Expected deck size to be %d after drawing, got %d", initialSize-1, deck.Size())
	}
	
	// Draw all remaining cards
	for i := range initialSize-1 {
		_, err := deck.Draw()
		if err != nil && i < initialSize-2 {
			t.Errorf("Expected no error when drawing card %d, got %v", i, err)
		}
	}
	
	// Try to draw from empty deck
	_, err = deck.Draw()
	if err == nil {
		t.Error("Expected error when drawing from empty deck, got nil")
	}
}

func TestDrawN(t *testing.T) {
	deck := NewDeck()
	initialSize := deck.Size()
	
	// Test drawing 5 cards
	cards, err := deck.DrawN(5)
	if err != nil {
		t.Errorf("Expected no error when drawing 5 cards, got %v", err)
	}
	
	if len(cards) != 5 {
		t.Errorf("Expected to draw 5 cards, got %d", len(cards))
	}
	
	if deck.Size() != initialSize-5 {
		t.Errorf("Expected deck size to be %d after drawing 5 cards, got %d", initialSize-5, deck.Size())
	}
	
	// Test drawing 0 cards (should error)
	_, err = deck.DrawN(0)
	if err == nil {
		t.Error("Expected error when drawing 0 cards, got nil")
	}
	
	// Test drawing more cards than are in the deck
	_, err = deck.DrawN(initialSize)
	if err == nil {
		t.Error("Expected error when drawing more cards than in deck, got nil")
	}
}

func TestAddToBottom(t *testing.T) {
	deck := NewDeck()
	initialSize := deck.Size()
	
	card := Card{Color: Red, Type: Number, Value: 5}
	deck.AddToBottom(card)
	
	if deck.Size() != initialSize+1 {
		t.Errorf("Expected deck size to be %d after adding a card to the bottom, got %d", initialSize+1, deck.Size())
	}
	
	// The card added should be at the bottom (index 0)
	bottomCard := deck.Cards[0]
	if bottomCard.Color != card.Color || bottomCard.Type != card.Type || bottomCard.Value != card.Value {
		t.Errorf("Expected bottom card to be %+v, got %+v", card, bottomCard)
	}
}

func TestIsEmpty(t *testing.T) {
	deck := NewDeck()
	
	if deck.IsEmpty() {
		t.Error("Expected new deck not to be empty")
	}
	
	// Draw all cards
	for i := range 108 {
		_, err := deck.Draw()
		if err != nil && i < 107 {
			t.Errorf("Unexpected error drawing card %d: %v", i, err)
		}
	}
	
	if !deck.IsEmpty() {
		t.Error("Expected deck to be empty after drawing all cards")
	}
}

func TestCreateDiscardPile(t *testing.T) {
	initialCard := Card{Color: Blue, Type: Number, Value: 3}
	discardPile := CreateDiscardPile(initialCard)
	
	if discardPile.Size() != 1 {
		t.Errorf("Expected discard pile to have 1 card, got %d", discardPile.Size())
	}
	
	topCard := discardPile.Cards[0]
	if topCard.Color != initialCard.Color || topCard.Type != initialCard.Type || topCard.Value != initialCard.Value {
		t.Errorf("Expected top card of discard pile to be %v, got %v", initialCard, topCard)
	}
}

func TestShuffleRandomness(t *testing.T) {
	deck := NewDeck()
	
	// Store the initial order
	initialOrder := make([]Card, len(deck.Cards))
	copy(initialOrder, deck.Cards)
	
	// Shuffle the deck
	deck.Shuffle()
	
	// Check that the order has changed (this test could occasionally fail by chance,
	// but the probability is extremely low with 108 cards)
	samePosition := 0
	for i := range deck.Cards {
		if i < len(initialOrder) && deck.Cards[i] == initialOrder[i] {
			samePosition++
		}
	}
	
	// Allow a small number of cards to remain in the same position by chance
	if samePosition > 20 {
		t.Errorf("Shuffle appears not to be sufficiently random. %d cards remained in the same position", samePosition)
	}
}
