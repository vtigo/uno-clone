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
		t.Errorf("Expected new playuer to have HasCalledUno = false")
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

	if(len(player.Hand) != handSizeBefore) {
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
		t.Errorf("Expected to play card1(c: %s, t: %s, v: %d), but got a different card(c: %s, t: %s, v: %d)",
		CardColor(card1.Color).String(), CardType(card1.Type).String(), card1.Value,
		CardColor(playerCard.Color).String(), CardType(playerCard.Type).String(), playerCard.Value)
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
		t.Errorf("Expcted player with empty hand but who never played ant cards to have HasWon = false")
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

func TestCallUno(t * testing.T) {
	player := NewPlayer("Test Player")

	if player.HasCalledUno {
		t.Errorf("Expected new player to have HasCalledUno = false")
	}

	player.CallUno()

	if !player.HasCalledUno {
		t.Errorf("Expected player to have HasCalledUno = true after calling uno")
	}
}
