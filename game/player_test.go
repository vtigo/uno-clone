package game

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	name := "TestPlayer"
	player := NewPlayer(name)
	
	if player.Name != name {
		t.Errorf("Expected player name to be %s, got %s", name, player.Name)
	}

	if len(player.Hand) != 0 {
		t.Errorf("Expected new player to have empty hand, got %d cards", len(player.Hand))
	}

	if player.HasCalledUno {
		t.Errorf("Expected new player to have HasCalledUno = false")
	}

	if player.IsMyTurn {
		t.Errorf("Expected new player to have IsMyTurn = false")
	}

	if player.HasWon() {
		t.Errorf("Expected new player to have HasWon = false")
	}
}

func TestAddCard(t *testing.T) {
	player := NewPlayer("TestPlayer")
	card := &Card{Color: Red, Type: Number, Value: 5}

	player.AddCard(card)

	if len(player.Hand) != 1 {
		t.Errorf("Expected hand size to be 1 after adding a card, got %d", len(player.Hand))
	}

	if player.Hand[0] != card {
		t.Errorf("Expected the added card to be in player's hand")
	}

	handSizeBefore := len(player.Hand)
	player.AddCard(nil)

	if len(player.Hand) != handSizeBefore {
		t.Errorf("Expected player hand size to remain %d after adding nil card, got %d", handSizeBefore, len(player.Hand))
	}

	player.CallUno()

	if !player.HasCalledUno {
		t.Errorf("Expected player to have HasCalledUno = true after calling uno")
	}

	player.AddCard(&Card{Color: Yellow, Type: Number, Value: 6})

	if len(player.Hand) != 2 {
		t.Errorf("Expected hand size to be 2 after adding a second card, got %d", len(player.Hand))
	}

	if player.HasCalledUno {
		t.Errorf("Expected player to have HasCalledUno = false after adding a second card")
	}
}

func TestAddCardsToHand(t *testing.T) {
	player := NewPlayer("TestPlayer")
	cards := []*Card{
		{Color: Red, Type: Number, Value: 5},
		{Color: Blue, Type: Number, Value: 7},
		{Color: Green, Type: Skip},
	}

	player.AddCardsToHand(cards)

	if len(player.Hand) != len(cards) {
		t.Errorf("Expected hand size to be %d after adding multiple cards, got %d", len(cards), len(player.Hand))
	}

	for i, card := range cards {
		if player.Hand[i] != card {
			t.Errorf("Expected card at index %d to be %v, got %v", i, card, player.Hand[i])
		}
	}
}

func TestPlayCard(t *testing.T) {
	player := NewPlayer("TestPlayer")
	card1 := &Card{Color: Red, Type: Number, Value: 5}
	card2 := &Card{Color: Yellow, Type: Number, Value: 6}

	player.AddCard(card1)
	player.AddCard(card2)

	playerCard, err := player.PlayCard(0)

	if err != nil {
		t.Errorf("Expected no errors when playing valid card, got %v", err)
	}

	if playerCard != card1 {
		t.Errorf("Expected to play card1, but got a different card")
	}

	if len(player.Hand) != 1 {
		t.Errorf("Expected hand size to be 1 after playing a card, got %d", len(player.Hand))
	}
	
	_, err = player.PlayCard(5)
	if err == nil {
		t.Errorf("Expected error when playing a card with invalid index")
	}

	_, err = player.PlayCard(-1)
	if err == nil {
		t.Errorf("Expected error when playing a card with negative index")
	}
}

func TestGetValidPlays(t *testing.T) {
	player := NewPlayer("TestPlayer")
	topCard := &Card{Color: Red, Type: Number, Value: 5}
	currentColor := Red

	// No cards in hand should return empty array
	validPlays := player.GetValidPlays(topCard, currentColor)
	if len(validPlays) != 0 {
		t.Errorf("Expected 0 valid plays with empty hand, got %d", len(validPlays))
	}

	// Add cards to hand
	redSeven := &Card{Color: Red, Type: Number, Value: 7}
	blueFive := &Card{Color: Blue, Type: Number, Value: 5}
	greenSkip := &Card{Color: Green, Type: Skip}
	wildCard := &Card{Color: Wild, Type: WildCard}

	player.AddCardsToHand([]*Card{redSeven, blueFive, greenSkip, wildCard})

	validPlays = player.GetValidPlays(topCard, currentColor)
	if len(validPlays) != 3 {
		t.Errorf("Expected 3 valid plays, got %d", len(validPlays))
	}

	// Change current color to green
	currentColor = Green
	validPlays = player.GetValidPlays(topCard, currentColor)
	if len(validPlays) != 2 {
		t.Errorf("Expected 2 valid plays with green as current color, got %d", len(validPlays))
	}
}

func TestHasValidPlay(t *testing.T) {
	player := NewPlayer("TestPlayer")
	topCard := &Card{Color: Red, Type: Number, Value: 5}
	currentColor := Red

	// No cards in hand should return false
	if player.HasValidPlay(topCard, currentColor) {
		t.Errorf("Expected HasValidPlay to be false with empty hand")
	}

	// Add cards with no valid plays
	blueSkip := &Card{Color: Blue, Type: Skip}
	greenSkip := &Card{Color: Green, Type: Skip}
	player.AddCardsToHand([]*Card{blueSkip, greenSkip})

	if player.HasValidPlay(topCard, currentColor) {
		t.Errorf("Expected HasValidPlay to be false with no matching cards")
	}

	// Add a wild card which is always valid
	wildCard := &Card{Color: Wild, Type: WildCard}
	player.AddCard(wildCard)

	if !player.HasValidPlay(topCard, currentColor) {
		t.Errorf("Expected HasValidPlay to be true after adding wild card")
	}
}

func TestHasWon(t *testing.T) {
	player := NewPlayer("TestPlayer")

	if player.HasWon() {
		t.Errorf("Expected new player to have HasWon = false")
	}

	player.AddCard(&Card{Color: Red, Type: Number, Value: 5})

	if player.HasWon() {
		t.Errorf("Expected player with cards to have HasWon = false")
	}

	player.Hand = []*Card{}

	if player.HasWon() {
		t.Errorf("Expected player with empty hand but who never played any cards to have HasWon = false")
	}

	player.AddCard(&Card{Color: Yellow, Type: Number, Value: 6})
	player.PlayCard(0)

	if !player.HasWon() {
		t.Errorf("Expected player who played all their cards to have HasWon = true")
	}
}

func TestShouldCallUno(t *testing.T) {
	player := NewPlayer("TestPlayer")

	if player.ShouldCallUno() {
		t.Errorf("Expected new player with empty hands to have ShouldCallUno = false")
	}

	player.AddCard(&Card{Color: Red, Type: Number, Value: 5})
	player.AddCard(&Card{Color: Yellow, Type: Number, Value: 6})

	if player.ShouldCallUno() {
		t.Errorf("Expected a player who never played any cards to have ShouldCallUno = false")
	}

	player.PlayCard(0)

	if !player.ShouldCallUno() {
		t.Errorf("Expected player who have played cards and has only 1 card in hand to have ShouldCallUno = true")
	}

	player.AddCard(&Card{Color: Green, Type: Number, Value: 7})

	if player.ShouldCallUno() {
		t.Errorf("Expected player with two cards in hand to have ShouldCallUno = false")
	}
}

func TestCallUno(t *testing.T) {
	player := NewPlayer("Test Player")

	if player.HasCalledUno {
		t.Errorf("Expected new player to have HasCalledUno = false")
	}

	player.CallUno()

	if !player.HasCalledUno {
		t.Errorf("Expected player to have HasCalledUno = true after calling uno")
	}

	player.ResetUnoCall()

	if player.HasCalledUno {
		t.Errorf("Expected player to have HasCalledUno = false after reset")
	}
}

func TestString(t *testing.T) {
	player := NewPlayer("TestPlayer")
	expected := "Player TestPlayer (0 cards) | Uno(false)"
	
	if player.String() != expected {
		t.Errorf("Expected string representation to be '%s', got '%s'", expected, player.String())
	}

	player.AddCard(&Card{Color: Red, Type: Number, Value: 5})
	player.CallUno()
	
	expected = "Player TestPlayer (1 cards) | Uno(true)"
	if player.String() != expected {
		t.Errorf("Expected string representation to be '%s', got '%s'", expected, player.String())
	}
}
