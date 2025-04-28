package game

import (
	"testing"
)

// Helper function to create a simple game state for testing
func createTestGameState() *GameState {
	player1 := NewPlayer("Player 1")
	player2 := NewPlayer("Player 2")
	
	deck := NewDeck()
	initialCard := Card{Color: Red, Type: Number, Value: 5}
	discardPile := CreateDiscardPile(initialCard)
	
	state := &GameState{
		Players:       []*Player{player1, player2},
		CurrentPlayer: 0,
		DrawPile:      deck,
		DiscardPile:   discardPile,
		ActiveColor:   Red,
		Phase:         PhasePlay,
		LastPlayedBy:  -1,
	}
	
	// Set player 1's turn
	player1.IsMyTurn = true
	
	return state
}

// Test NewGameRules
func TestNewGameRules(t *testing.T) {
	rules := NewGameRules()
	if rules == nil {
		t.Error("Expected NewGameRules to return a non-nil GameRules instance")
	}
}

// Test NewGameState
func TestNewGameState(t *testing.T) {
	player1 := NewPlayer("Player 1")
	player2 := NewPlayer("Player 2")
	
	// Test with correct number of players
	state, err := NewGameState([]*Player{player1, player2})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if state == nil {
		t.Fatal("Expected NewGameState to return a non-nil GameState instance")
	}
	
	// Check initial setup
	if len(state.Players) != 2 {
		t.Errorf("Expected 2 players, got %d", len(state.Players))
	}
	
	if state.CurrentPlayer != 0 {
		t.Errorf("Expected currentPlayer to be 0, got %d", state.CurrentPlayer)
	}
	
	if player1.HandSize() != 7 {
		t.Errorf("Expected player 1 to have 7 cards, got %d", player1.HandSize())
	}
	
	if player2.HandSize() != 7 {
		t.Errorf("Expected player 2 to have 7 cards, got %d", player2.HandSize())
	}
	
	if state.Phase != PhasePlay && state.Phase != PhaseColorSelection {
		t.Errorf("Expected game phase to be Play or ColorSelection, got %d", state.Phase)
	}
	
	// Test with incorrect number of players
	_, err = NewGameState([]*Player{player1})
	if err == nil {
		t.Error("Expected error when initializing game with 1 player")
	}
	
	_, err = NewGameState([]*Player{player1, player2, NewPlayer("Player 3")})
	if err == nil {
		t.Error("Expected error when initializing game with 3 players")
	}
}

// Test ValidateMove
func TestValidateMove(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Add cards to player 1's hand
	redSeven := &Card{Color: Red, Type: Number, Value: 7}
	blueFive := &Card{Color: Blue, Type: Number, Value: 5}
	blueSkip := &Card{Color: Blue, Type: Skip}
	wildCard := &Card{Color: Wild, Type: WildCard}
	wildDrawFour := &Card{Color: Wild, Type: WildDrawFour}
	
	state.Players[0].AddCardsToHand([]*Card{redSeven, blueFive, blueSkip, wildCard, wildDrawFour})
	
	// Test valid moves
	valid, _ := rules.ValidateMove(state.Players[0], 0, state) // Red 7 on Red 5
	if !valid {
		t.Error("Expected playing Red 7 on Red 5 to be valid")
	}
	
	valid, _ = rules.ValidateMove(state.Players[0], 3, state) // Wild card on Red 5
	if !valid {
		t.Error("Expected playing Wild card to be valid")
	}
	
	// Test invalid moves
	valid, _ = rules.ValidateMove(state.Players[0], 2, state) // Blue Skip on Red 5
	if valid {
		t.Error("Expected playing Blue Skip on Red 5 to be invalid")
	}
	
	// Test playing when it's not the player's turn
	state.Players[0].IsMyTurn = false
	state.Players[1].IsMyTurn = true
	
	valid, msg := rules.ValidateMove(state.Players[0], 0, state)
	if valid {
		t.Error("Expected playing when it's not the player's turn to be invalid")
	}
	if msg != "It's not your turn" {
		t.Errorf("Expected error message 'It's not your turn', got '%s'", msg)
	}
	
	// Reset turn
	state.Players[0].IsMyTurn = true
	state.Players[1].IsMyTurn = false
	
	// Test invalid card index
	valid, msg = rules.ValidateMove(state.Players[0], 10, state)
	if valid {
		t.Error("Expected playing with invalid card index to be invalid")
	}
	if msg != "Invalid card index" {
		t.Errorf("Expected error message 'Invalid card index', got '%s'", msg)
	}
	
	// Test playing in wrong phase
	state.Phase = PhaseColorSelection
	
	valid, msg = rules.ValidateMove(state.Players[0], 0, state)
	if valid {
		t.Error("Expected playing in wrong phase to be invalid")
	}
	if msg != "Game is not in the play phase" {
		t.Errorf("Expected error message 'Game is not in the play phase', got '%s'", msg)
	}
	
	// Reset phase
	state.Phase = PhasePlay
	
	// Test invalid Wild Draw Four play (when player has matching color)
	state.ActiveColor = Blue
	valid, msg = rules.ValidateMove(state.Players[0], 4, state) // Wild Draw Four when player has Blue card
	if valid {
		t.Error("Expected playing Wild Draw Four when having matching color cards to be invalid")
	}
	if msg != "Wild Draw Four can only be played if you don't have any cards of the active color" {
		t.Errorf("Expected error message about Wild Draw Four restriction, got '%s'", msg)
	}
}

// Test handling Number card effect
func TestHandleNumberCardEffect(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Record initial state
	initialCurrentPlayer := state.CurrentPlayer
	
	// Handle number card effect
	err := rules.handleNumberCard(state)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	// Verify the current player changed (next turn)
	if state.CurrentPlayer != (initialCurrentPlayer+1)%len(state.Players) {
		t.Error("Expected current player to change after playing number card")
	}
	
	// Verify player turn flags
	if state.Players[initialCurrentPlayer].IsMyTurn {
		t.Error("Expected previous player's turn flag to be false")
	}
	
	if !state.Players[state.CurrentPlayer].IsMyTurn {
		t.Error("Expected new current player's turn flag to be true")
	}
}

// Test handling Skip card effect
func TestHandleSkipCardEffect(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Record initial state
	initialCurrentPlayer := state.CurrentPlayer
	
	// Handle skip card effect
	err := rules.handleSkipCard(state)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	// For 2-player game, Skip should keep the turn with the current player
	if state.CurrentPlayer != initialCurrentPlayer {
		t.Error("Expected current player to remain the same after playing skip card in 2-player game")
	}
	
	// Verify player turn flags
	if !state.Players[initialCurrentPlayer].IsMyTurn {
		t.Error("Expected current player's turn flag to still be true")
	}
	
	if state.Players[(initialCurrentPlayer+1)%len(state.Players)].IsMyTurn {
		t.Error("Expected other player's turn flag to be false")
	}
}

// Test handling Reverse card effect
func TestHandleReverseCardEffect(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Record initial state
	initialCurrentPlayer := state.CurrentPlayer
	
	// Handle reverse card effect
	err := rules.handleReverseCard(state)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	// For 2-player game, Reverse should act like Skip
	if state.CurrentPlayer != initialCurrentPlayer {
		t.Error("Expected current player to remain the same after playing reverse card in 2-player game")
	}
	
	// Verify player turn flags
	if !state.Players[initialCurrentPlayer].IsMyTurn {
		t.Error("Expected current player's turn flag to still be true")
	}
	
	if state.Players[(initialCurrentPlayer+1)%len(state.Players)].IsMyTurn {
		t.Error("Expected other player's turn flag to be false")
	}
}

// Test handling Draw Two card effect
func TestHandleDrawTwoCardEffect(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Record initial state
	initialCurrentPlayer := state.CurrentPlayer
	nextPlayer := (initialCurrentPlayer + 1) % len(state.Players)
	initialHandSize := state.Players[nextPlayer].HandSize()
	
	// Handle draw two card effect
	err := rules.handleDrawTwoCard(state)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	// Verify next player drew 2 cards
	expectedHandSize := initialHandSize + 2
	if state.Players[nextPlayer].HandSize() != expectedHandSize {
		t.Errorf("Expected next player to have %d cards, got %d", expectedHandSize, state.Players[nextPlayer].HandSize())
	}
	
	// Verify turn remains with current player
	if state.CurrentPlayer != initialCurrentPlayer {
		t.Error("Expected current player to remain the same after playing draw two card")
	}
	
	// Verify player turn flags
	if !state.Players[initialCurrentPlayer].IsMyTurn {
		t.Error("Expected current player's turn flag to still be true")
	}
	
	if state.Players[nextPlayer].IsMyTurn {
		t.Error("Expected other player's turn flag to be false")
	}
}

// Test handling Wild card effect
func TestHandleWildCardEffect(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Record initial state
	initialCurrentPlayer := state.CurrentPlayer
	
	// Handle wild card effect with chosen color
	chosenColor := Blue
	err := rules.handleWildCard(state, chosenColor)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	// Verify active color changed
	if state.ActiveColor != chosenColor {
		t.Errorf("Expected active color to be %v, got %v", chosenColor, state.ActiveColor)
	}
	
	// Verify turn moved to next player
	expectedNextPlayer := (initialCurrentPlayer + 1) % len(state.Players)
	if state.CurrentPlayer != expectedNextPlayer {
		t.Errorf("Expected current player to be %d, got %d", expectedNextPlayer, state.CurrentPlayer)
	}
	
	// Verify player turn flags
	if state.Players[initialCurrentPlayer].IsMyTurn {
		t.Error("Expected previous player's turn flag to be false")
	}
	
	if !state.Players[expectedNextPlayer].IsMyTurn {
		t.Error("Expected new current player's turn flag to be true")
	}
	
	// Test invalid color
	state = createTestGameState()
	invalidColor := Wild
	err = rules.handleWildCard(state, invalidColor)
	
	if err == nil {
		t.Error("Expected error when choosing invalid color for Wild card")
	}
}

// Test handling Wild Draw Four card effect
func TestHandleWildDrawFourCardEffect(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Record initial state
	initialCurrentPlayer := state.CurrentPlayer
	nextPlayer := (initialCurrentPlayer + 1) % len(state.Players)
	initialHandSize := state.Players[nextPlayer].HandSize()
	
	// Handle wild draw four card effect with chosen color
	chosenColor := Green
	err := rules.handleWildDrawFourCard(state, chosenColor)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	// Verify active color changed
	if state.ActiveColor != chosenColor {
		t.Errorf("Expected active color to be %v, got %v", chosenColor, state.ActiveColor)
	}
	
	// Verify next player drew 4 cards
	expectedHandSize := initialHandSize + 4
	if state.Players[nextPlayer].HandSize() != expectedHandSize {
		t.Errorf("Expected next player to have %d cards, got %d", expectedHandSize, state.Players[nextPlayer].HandSize())
	}
	
	// Verify turn remains with current player
	if state.CurrentPlayer != initialCurrentPlayer {
		t.Error("Expected current player to remain the same after playing wild draw four card")
	}
	
	// Verify player turn flags
	if !state.Players[initialCurrentPlayer].IsMyTurn {
		t.Error("Expected current player's turn flag to still be true")
	}
	
	if state.Players[nextPlayer].IsMyTurn {
		t.Error("Expected other player's turn flag to be false")
	}
	
	// Test invalid color
	state = createTestGameState()
	invalidColor := Wild
	err = rules.handleWildDrawFourCard(state, invalidColor)
	
	if err == nil {
		t.Error("Expected error when choosing invalid color for Wild Draw Four card")
	}
}

// Test NextTurn
func TestNextTurn(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Record initial state
	initialCurrentPlayer := state.CurrentPlayer
	
	// Move to next turn
	rules.NextTurn(state)
	
	// Verify turn moved to next player
	expectedNextPlayer := (initialCurrentPlayer + 1) % len(state.Players)
	if state.CurrentPlayer != expectedNextPlayer {
		t.Errorf("Expected current player to be %d, got %d", expectedNextPlayer, state.CurrentPlayer)
	}
	
	// Verify player turn flags
	if state.Players[initialCurrentPlayer].IsMyTurn {
		t.Error("Expected previous player's turn flag to be false")
	}
	
	if !state.Players[expectedNextPlayer].IsMyTurn {
		t.Error("Expected new current player's turn flag to be true")
	}
}

// Test SkipTurn, RepeatTurn, and ReverseTurn in 2-player game
func TestSpecialTurnsInTwoPlayerGame(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Record initial state
	initialCurrentPlayer := state.CurrentPlayer
	
	// Test SkipTurn
	rules.SkipTurn(state)
	
	// In 2-player game, SkipTurn should keep the current player
	if state.CurrentPlayer != initialCurrentPlayer {
		t.Error("Expected SkipTurn to keep the same current player in 2-player game")
	}
	
	// Test RepeatTurn
	rules.RepeatTurn(state)
	
	// In 2-player game, RepeatTurn should keep the current player
	if state.CurrentPlayer != initialCurrentPlayer {
		t.Error("Expected RepeatTurn to keep the same current player in 2-player game")
	}
	
	// Test ReverseTurn
	rules.ReverseTurn(state)
	
	// In 2-player game, ReverseTurn should keep the current player
	if state.CurrentPlayer != initialCurrentPlayer {
		t.Error("Expected ReverseTurn to keep the same current player in 2-player game")
	}
}

// Test HandleUnoCall
func TestHandleUnoCall(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Test calling UNO with invalid player index
	success, _ := rules.HandleUnoCall(-1, state)
	if success {
		t.Error("Expected UNO call with invalid player index to fail")
	}
	
	// Add cards to player's hand
	redSeven := &Card{Color: Red, Type: Number, Value: 7}
	state.Players[0].AddCard(redSeven)
	
	// Test calling UNO with more than one card
	success, _ = rules.HandleUnoCall(0, state)
	if success {
		t.Error("Expected UNO call with more than one card to fail")
	}
	
	// Remove cards from hand to have one left and make player play at least one card
	state.Players[0].hasPlayedCard = true
	state.Players[0].Hand = []*Card{redSeven}
	
	// Test valid UNO call
	success, _ = rules.HandleUnoCall(0, state)
	if !success {
		t.Error("Expected valid UNO call to succeed")
	}
	
	// Verify UNO flag
	if !state.Players[0].HasCalledUno {
		t.Error("Expected player's UNO flag to be true after calling UNO")
	}
}

// Test HandleUnoChallenge
func TestHandleUnoChallenge(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Test challenging with invalid target index
	success, _ := rules.HandleUnoChallenge(-1, state)
	if success {
		t.Error("Expected UNO challenge with invalid target index to fail")
	}
	
	// Add cards to player's hand
	redSeven := &Card{Color: Red, Type: Number, Value: 7}
	state.Players[0].AddCard(redSeven)
	
	// Test challenging when player has more than one card
	success, _ = rules.HandleUnoChallenge(0, state)
	if success {
		t.Error("Expected UNO challenge when player has more than one card to fail")
	}
	
	// Remove cards from hand to have one left and make player play at least one card
	state.Players[0].hasPlayedCard = true
	state.Players[0].Hand = []*Card{redSeven}
	
	// Test challenging a player with 1 card but who called UNO
	state.Players[0].CallUno()
	success, _ = rules.HandleUnoChallenge(0, state)
	if success {
		t.Error("Expected UNO challenge when player has called UNO to fail")
	}
	
	// Reset UNO call
	state.Players[0].ResetUnoCall()
	
	// Test valid UNO challenge
	initialHandSize := state.Players[0].HandSize()
	success, _ = rules.HandleUnoChallenge(0, state)
	if !success {
		t.Error("Expected valid UNO challenge to succeed")
	}
	
	// Verify player drew 2 cards as penalty
	expectedHandSize := initialHandSize + 2
	if state.Players[0].HandSize() != expectedHandSize {
		t.Errorf("Expected player to have %d cards after UNO challenge, got %d", expectedHandSize, state.Players[0].HandSize())
	}
}

// Test HandlePlayCard
func TestHandlePlayCard(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Add cards to player 1's hand
	redSeven := &Card{Color: Red, Type: Number, Value: 7}
	redSkip := &Card{Color: Red, Type: Skip}
	wildCard := &Card{Color: Wild, Type: WildCard}
	
	state.Players[0].AddCardsToHand([]*Card{redSeven, redSkip, wildCard})
	
	// Test playing a number card
	initialHandSize := state.Players[0].HandSize()
	initialDiscardPileSize := state.DiscardPile.Size()
	err := rules.HandlePlayCard(state.Players[0], 0, state, nil)
	
	if err != nil {
		t.Errorf("Expected no error when playing valid card, got %v", err)
	}
	
	// Verify card was removed from hand
	if state.Players[0].HandSize() != initialHandSize-1 {
		t.Errorf("Expected hand size to decrease by 1, got %d", state.Players[0].HandSize())
	}
	
	// Verify card was added to discard pile
	if state.DiscardPile.Size() != initialDiscardPileSize+1 {
		t.Errorf("Expected discard pile size to increase by 1, got %d", state.DiscardPile.Size())
	}
	
	// Verify turn moved to next player for number card
	if state.CurrentPlayer != 1 {
		t.Errorf("Expected current player to be 1, got %d", state.CurrentPlayer)
	}
	
	// Test playing a wild card without choosing color
	state.CurrentPlayer = 0 // Reset turn
	state.Players[0].IsMyTurn = true
	state.Players[1].IsMyTurn = false
	
	// Add a wild card to the hand (since we played cards earlier)
	state.Players[0].AddCard(wildCard)
	
	err = rules.HandlePlayCard(state.Players[0], state.Players[0].HandSize()-1, state, nil) // Play the wild card
	
	if err != nil {
		t.Errorf("Expected no error when playing wild card without color, got %v", err)
	}
	
	// Verify phase changed to ColorSelection
	// Note: If this test fails, check if the actual behavior of HandlePlayCard matches expected behavior
	if state.Phase != PhaseColorSelection {
		t.Errorf("Expected phase to change to ColorSelection, got %d", state.Phase)
	}
	
	// Test playing a wild card with chosen color
	state.CurrentPlayer = 0 // Reset turn
	state.Players[0].IsMyTurn = true
	state.Players[1].IsMyTurn = false
	state.Phase = PhasePlay // Reset phase
	
	// Update player's hand with a new wild card for testing
	state.Players[0].AddCard(wildCard)
	
	chosenColor := Blue
	err = rules.HandlePlayCard(state.Players[0], state.Players[0].HandSize()-1, state, &chosenColor)
	
	if err != nil {
		t.Errorf("Expected no error when playing wild card with color, got %v", err)
	}
	
	// Verify active color changed
	if state.ActiveColor != chosenColor {
		t.Errorf("Expected active color to be %v, got %v", chosenColor, state.ActiveColor)
	}
	
	// Verify LastPlayedBy was updated
	if state.LastPlayedBy != 0 {
		t.Errorf("Expected LastPlayedBy to be 0, got %d", state.LastPlayedBy)
	}
}

// Test HandleDrawCard
func TestHandleDrawCard(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Record initial hand size
	initialHandSize := state.Players[0].HandSize()
	
	// Test drawing a card
	err := rules.HandleDrawCard(state.Players[0], state)
	
	if err != nil {
		t.Errorf("Expected no error when drawing card, got %v", err)
	}
	
	// Verify hand size increased
	if state.Players[0].HandSize() != initialHandSize+1 {
		t.Errorf("Expected hand size to increase by 1, got %d", state.Players[0].HandSize())
	}
	
	// Test drawing when it's not player's turn
	state.Players[0].IsMyTurn = false
	
	err = rules.HandleDrawCard(state.Players[0], state)
	if err == nil {
		t.Error("Expected error when drawing card when it's not player's turn")
	}
	
	// Test drawing when game is not in play phase
	state.Players[0].IsMyTurn = true
	state.Phase = PhaseColorSelection
	
	err = rules.HandleDrawCard(state.Players[0], state)
	if err == nil {
		t.Error("Expected error when drawing card when game is not in play phase")
	}
}

// Test EndTurn
func TestEndTurn(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Record initial state
	initialCurrentPlayer := state.CurrentPlayer
	
	// Test ending turn
	err := rules.EndTurn(state)
	
	if err != nil {
		t.Errorf("Expected no error when ending turn, got %v", err)
	}
	
	// Verify turn moved to next player
	expectedNextPlayer := (initialCurrentPlayer + 1) % len(state.Players)
	if state.CurrentPlayer != expectedNextPlayer {
		t.Errorf("Expected current player to be %d, got %d", expectedNextPlayer, state.CurrentPlayer)
	}
	
	// Verify player turn flags
	if state.Players[initialCurrentPlayer].IsMyTurn {
		t.Error("Expected previous player's turn flag to be false")
	}
	
	if !state.Players[expectedNextPlayer].IsMyTurn {
		t.Error("Expected new current player's turn flag to be true")
	}
	
	// Test ending turn when game is not in play phase
	state.Phase = PhaseColorSelection
	
	err = rules.EndTurn(state)
	if err == nil {
		t.Error("Expected error when ending turn when game is not in play phase")
	}
}

// Test winning condition
func TestWinningCondition(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Set up a scenario where player has one card
	redSeven := &Card{Color: Red, Type: Number, Value: 7}
	state.Players[0].Hand = []*Card{redSeven}
	state.Players[0].hasPlayedCard = true // Simulate having played cards before
	
	// Play the final card
	err := rules.HandlePlayCard(state.Players[0], 0, state, nil)
	
	if err != nil {
		t.Errorf("Expected no error when playing final card, got %v", err)
	}
	
	// Verify game is over
	if state.Phase != PhaseGameOver {
		t.Errorf("Expected game phase to be GameOver, got %d", state.Phase)
	}
	
	// Verify player has won
	if !state.Players[0].HasWon() {
		t.Error("Expected player to have won after playing final card")
	}
}

// Test recycling discard pile when draw pile is empty
func TestRecycleDiscardPile(t *testing.T) {
	rules := NewGameRules()
	state := createTestGameState()
	
	// Empty the draw pile
	for !state.DrawPile.IsEmpty() {
		card, _ := state.DrawPile.Draw()
		state.DiscardPile.AddToBottom(card)
	}
	
	// Add several cards to discard pile
	for i := range 5 {
		state.DiscardPile.AddToBottom(Card{Color: CardColor(i % 4), Type: Number, Value: i})
	}
	
	discardPileSize := state.DiscardPile.Size()
	playerHandSize := state.Players[0].HandSize()
	
	// Try to draw a card - should recycle the discard pile
	err := rules.HandleDrawCard(state.Players[0], state)
	
	if err != nil {
		t.Errorf("Expected no error when drawing card after recycling, got %v", err)
	}
	
	// Verify a card was drawn
	if state.Players[0].HandSize() != playerHandSize + 1 {
		t.Errorf("Expected player hand size to increase by 1, got %d instead of %d", 
			state.Players[0].HandSize(), playerHandSize + 1)
	}
	
	// Verify discard pile was recycled
	if state.DrawPile.IsEmpty() {
		t.Error("Expected draw pile to have cards after recycling")
	}
	
	// Verify discard pile has only one card (the top card)
	if state.DiscardPile.Size() != 1 {
		t.Errorf("Expected discard pile to have 1 card, got %d", state.DiscardPile.Size())
	}
	
	// Verify the math works - verify we're not losing or creating cards
	// The player drew 1 card, and 1 card should remain in the discard pile
	// So the draw pile should have (discardPileSize - 2) cards
	expectedDrawPileSize := discardPileSize - 2
	if state.DrawPile.Size() != expectedDrawPileSize {
		t.Errorf("Expected draw pile to have %d cards, got %d", expectedDrawPileSize, state.DrawPile.Size())
	}
}
