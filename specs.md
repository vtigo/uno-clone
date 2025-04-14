# Two-Player UNO Clone Specification

## Project Overview
This specification outlines the development of a two-player UNO card game clone implemented in Go using the Ebiten game engine. The game will feature a graphical user interface with pixel art style, networked multiplayer over LAN, and will follow the special rules for two-player UNO.

## Game Rules
### Two-Player UNO Rules
1. Playing a Reverse card acts like a Skip. The player who plays the Reverse may immediately play another card.
2. The person playing a Skip card may immediately play another card.
3. When one person plays a Draw Two card and the other player has drawn the 2 cards, the play is back to the first person. The same principle applies to the Wild Draw Four card.
4. The usual UNO card game rules apply in all other instances.

## Architecture
### Technology Stack
- **Language**: Go
- **Game Engine**: Ebiten (Ebitengine)
- **Networking**: WebSockets (using Gorilla WebSocket library)
- **Serialization**: JSON
- **GUI Framework**: Ebiten's built-in rendering capabilities

### Project Structure
Simple package-based structure:
```
uno/
├── assets/
│   ├── images/
│   │   └── cards/ 
│   ├── audio/
│   └── fonts/
├── game/
│   ├── card.go      # Card definitions and deck management
│   ├── player.go    # Player information and hand management
│   ├── rules.go     # Game rules and logic
│   └── state.go     # Game state management
├── ui/
│   ├── screens.go   # Different game screens (title, gameplay, results)
│   ├── render.go    # Rendering logic
│   └── input.go     # Input handling
├── net/
│   ├── server.go    # WebSocket server implementation
│   ├── client.go    # WebSocket client implementation
│   ├── discovery.go # LAN game discovery
│   └── protocol.go  # Network protocol definitions
├── main.go          # Entry point
├── go.mod
└── go.sum
```

### Dependencies
- github.com/hajimehoshi/ebiten/v2
- github.com/hajimehoshi/ebiten/v2/audio
- github.com/hajimehoshi/ebiten/v2/ebitenutil
- github.com/hajimehoshi/ebiten/v2/inpututil
- github.com/hajimehoshi/ebiten/v2/text
- golang.org/x/image/font
- github.com/golang/freetype/truetype
- github.com/gorilla/websocket

## Game Components

### Card Representation
```go
type CardColor int
type CardType int

const (
    Red CardColor = iota
    Blue
    Green
    Yellow
    Wild // For wild cards
)

const (
    Number CardType = iota
    Skip
    Reverse
    DrawTwo
    WildCard
    WildDrawFour
)

// Card struct to represent each card
type Card struct {
    Color CardColor
    Type  CardType
    Value int // Only used for number cards (0-9)
}
```

### Deck Management
- Standard UNO deck with original card distribution
- 108 cards in total:
  - 19 red cards (0-9, with duplicates of 1-9)
  - 19 blue cards (0-9, with duplicates of 1-9)
  - 19 green cards (0-9, with duplicates of 1-9)
  - 19 yellow cards (0-9, with duplicates of 1-9)
  - 8 Skip cards (2 of each color)
  - 8 Reverse cards (2 of each color)
  - 8 Draw Two cards (2 of each color)
  - 4 Wild cards
  - 4 Wild Draw Four cards

### Player Management
- Two players only (human vs human over network)
- Players identified by customizable names
- Each player starts with 7 cards

### Game State Management
- Simple state machine for different game screens:
  - TitleScreen
  - RulesScreen
  - SettingsScreen
  - GameplayScreen
  - ResultsScreen
- Direct function calls for handling card effects:
```go
func handleCardPlay(card Card) {
    // Basic card effect handling
    switch card.Type {
    case Skip:
        skipNextPlayer()
        allowCurrentPlayerPlayAgain()
    case Reverse:
        // In two-player game, acts like Skip
        skipNextPlayer()
        allowCurrentPlayerPlayAgain()
    case DrawTwo:
        currentPlayer.DrawCards(2)
        skipNextPlayer()
    // other cases...
    }
}
```

## User Interface

### Screens
1. **Title Screen**
   - "Play Game" button
   - "Rules" button
   - "Settings" button
   - "Exit" button

2. **Rules Screen**
   - Display two-player UNO rules
   - "Back" button

3. **Settings Screen**
   - Sound effects toggle
   - Full screen toggle
   - Player names customization
   - "Back" button

4. **Gameplay Screen**
   - Player hands
   - Draw pile
   - Discard pile
   - Current player indicator
   - "UNO" button (for when a player has one card left)
   - "End Turn" button
   - Color selection wheel for Wild cards

5. **Results Screen**
   - Winner name
   - Cards left in opponent's hand
   - Total time played
   - "Play Again" button
   - "Back to Title" button

### Visual Elements
- Custom pixel art style for all cards and UI elements
- Pre-loaded card images for better performance
- Visual indicators for valid moves (automatic validation)
- Combination of visual indicators and sound effects for notifications

## Networking

### Architecture
- Client-server model where one player acts as host
- WebSockets for network communication
- Server as source of truth for game state
- JSON for data serialization

### Game Discovery
- LAN discovery system to find games on the local network
- Players can join games without knowing IP addresses

### State Synchronization
- Server validates all game actions
- After validation, server broadcasts updated game state to all clients
- Clients render based on the authoritative state from server

### Disconnection Handling
- If a player disconnects, the match is terminated
- Simple popup notification for the remaining player

## Game Flow

### Game Initialization
1. Player starts the game and chooses to host or join
2. If hosting, server starts and waits for connections
3. If joining, client searches for games on LAN
4. Players connect and enter customizable names
5. Host starts the game
6. Cards are dealt (7 to each player)
7. Initial card is placed on discard pile
8. First player is determined

### Turn Sequence
1. Current player selects a card to play (if valid)
2. Card effect is applied
3. Player clicks "End Turn" button
4. If player has one card remaining, they must click "UNO" button (honor system)
5. Turn passes to next player
6. Process repeats until a player plays their last card

### Special Card Handling
- Skip: Current player plays again
- Reverse: In two-player game, acts like Skip, current player plays again
- Draw Two: Other player draws 2 cards, loses turn
- Wild: Player selects new color via color wheel
- Wild Draw Four: Other player draws 4 cards, loses turn, current player selects new color

### Game End
1. Player plays their last card
2. Score summary screen displays:
   - Winner name
   - Cards left in opponent's hand
   - Total time played
3. Players can choose to play again or return to title screen

## Implementation Details

### Error Handling
- Simple but present error handling approach
- Graceful handling of network errors
- Validation of all user inputs
- Clear user feedback for invalid actions

### Audio
- Sound effects for various game events (to be implemented as continuous improvement):
  - Playing a card
  - Drawing a card
  - Saying "UNO"
  - Winning the game
  - Using special cards
  - Background music
  - Invalid move attempt

### Testing Strategy
- **Unit Tests**:
  - Card validation rules
  - Game state transitions
  - Special card effects
  - Winning conditions

- **Integration Tests**:
  - Game state and UI interaction
  - Network communication
  - Complete game flows

- **Manual Testing**:
  - Visual elements and UI
  - Network performance
  - Overall gameplay experience
