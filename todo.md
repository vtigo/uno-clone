# UNO Clone Implementation Todo Checklist

## Phase 1: Project Setup and Core Components

### Project Initialization
- [X] Create directory structure according to specs:
  - [X] Create `assets/` directory (with subdirectories for images, audio, fonts)
  - [X] Create `game/` directory for game logic
  - [X] Create `ui/` directory for UI components
  - [X] Create `net/` directory for network components
  - [X] Create `audio/` directory for sound management
- [X] Initialize Go module (`go mod init`)
- [X] Add required dependencies:
  - [X] github.com/hajimehoshi/ebiten/v2
  - [X] github.com/hajimehoshi/ebiten/v2/audio
  - [X] github.com/hajimehoshi/ebiten/v2/ebitenutil
  - [X] github.com/hajimehoshi/ebiten/v2/inpututil
  - [X] github.com/hajimehoshi/ebiten/v2/text
  - [X] golang.org/x/image/font
  - [X] github.com/golang/freetype/truetype
  - [X] github.com/gorilla/websocket
- [X] Create config.go file with game constants
- [X] Set up basic main.go with Ebiten game loop

### Card and Deck Implementation
- [ ] Define Card struct and related types in game/card.go:
  - [X] CardColor enum (Red, Blue, Green, Yellow, Wild)
  - [X] CardType enum (Number, Skip, Reverse, DrawTwo, WildCard, WildDrawFour)
  - [X] Card struct with Color, Type, and Value fields
  - [ ] Methods for card comparison and matching (according to UNO rules)
  - [X] String/display formatting methods for cards
- [ ] Implement standard UNO deck management:
  - [X] CreateDeck function that generates all 108 cards with proper distribution:
    - [X] 19 red cards (0-9, with duplicates of 1-9)
    - [X] 19 blue cards (0-9, with duplicates of 1-9)
    - [X] 19 green cards (0-9, with duplicates of 1-9)
    - [X] 19 yellow cards (0-9, with duplicates of 1-9)
    - [X] 8 Skip cards (2 of each color)
    - [X] 8 Reverse cards (2 of each color)
    - [X] 8 Draw Two cards (2 of each color)
    - [X] 4 Wild cards
    - [X] 4 Wild Draw Four cards
  - [X] Shuffle function using Fisher-Yates algorithm
  - [X] DrawCard function to take cards from the top of the deck
  - [ ] Reshuffle function to reuse discarded cards when deck is empty
- [ ] Write comprehensive unit tests:
  - [X] Test card creation and properties
  - [X] Test deck generation (correct number and distribution of cards)
  - [X] Test shuffle randomness
  - [X] Test draw functionality
  - [ ] Test card matching rules

### Player Implementation
- [ ] Create Player struct in game/player.go:
  - [X] Name field (string)
  - [X] Hand field (slice of Card pointers)
  - [X] HasCalledUno flag to track UNO calls
  - [ ] IsConnected flag for network play status
- [ ] Implement player methods:
  - [X] NewPlayer function to create player with name
  - [X] AddCard method to add cards to hand
  - [X] PlayCard method to remove/play cards from hand
  - [X] HasWon method to check for win condition (no cards)
  - [X] ShouldCallUno method to check if player has one card
  - [X] CallUno method to mark that player has called UNO
  - [ ] GetPlayableCards method to find all valid plays
  - [ ] CanPlayWildDrawFour method to validate Wild Draw Four plays
- [ ] Implement hand management:
  - [ ] Sort hand by color and number for easier play
  - [ ] Track cards drawn this turn (for draw-and-play rule)
  - [ ] Methods to find cards matching a specific color or value
- [ ] Write unit tests for player functionality:
  - [X] Test hand management
  - [ ] Test valid play detection
  - [X] Test UNO call status
  - [X] Test win condition
  - [ ] Test Wild Draw Four validation

### Game State Management
- [ ] Create GameState struct in game/state.go:
  - [ ] GameScreen enum for different screens
  - [ ] CurrentScreen field to track current screen
  - [ ] Players array for both players
  - [ ] CurrentPlayer index to track active player
  - [ ] CurrentColor field to track active color (may differ from top card after Wild)
  - [ ] Deck field for remaining cards
  - [ ] DiscardPile field for played cards
  - [ ] CurrentCard field for top card
  - [ ] GameStatus enum (NotStarted, InProgress, GameOver)
  - [ ] DrawCount to track cards that need to be drawn (Draw Two/Four)
  - [ ] WaitingForColorSelection flag
  - [ ] CanPlayAfterDraw flag
  - [ ] CanPlayAgain flag (for Skip/Reverse effects)
  - [ ] GameStartTime and GameEndTime for timing
- [ ] Implement game initialization:
  - [ ] NewGameState function to initialize game
  - [ ] InitGame method to set up a new game
  - [ ] DealCards method to distribute initial 7 cards per player
  - [ ] SetupFirstCard method (ensure first card isn't a special card)
- [ ] Implement turn management:
  - [ ] StartTurn method to begin a player's turn
  - [ ] NextTurn method for turn progression
  - [ ] HandleCardPlay method for playing a card
  - [ ] HandleDrawCard method for drawing a card
  - [ ] HandleUnoCall method
  - [ ] HandleColorSelection for Wild cards
  - [ ] HandlePlayAgain for Skip/Reverse in two-player mode
- [ ] Implement game progression:
  - [ ] CheckGameOver method
  - [ ] CalculateScore method
  - [ ] ResetGame method for starting over
  - [ ] SaveGameState and LoadGameState for persistence
- [ ] Write unit tests for all game state functionality:
  - [ ] Test initialization
  - [ ] Test turn progression
  - [ ] Test card effects specific to two-player rules
  - [ ] Test UNO call handling
  - [ ] Test game over conditions

## Phase 2: Game Logic and Rules

### Two-Player UNO Rules Implementation
- [ ] Implement core card matching rules:
  - [ ] Cards must match by color, number, or action type
  - [ ] Number cards (0-9) playable on matching color or number
  - [ ] Action cards (Skip, Reverse, Draw Two) playable on matching color or action
  - [ ] Wild cards can be played on any card
  - [ ] Wild Draw Four only playable when player has no matching color in hand
- [ ] Implement special two-player mechanics:
  - [ ] Skip card: Current player plays again immediately
  - [ ] Reverse card: Acts like Skip (current player plays again immediately)
  - [ ] Draw Two card: Other player draws 2 cards, loses turn, play returns to first player
  - [ ] Wild card: Player selects the new color, turn passes to opponent
  - [ ] Wild Draw Four: Other player draws 4 cards, loses turn, first player selects color
- [ ] Implement draw mechanics:
  - [ ] If player cannot play, they must draw one card
  - [ ] If drawn card is playable, player may play it immediately
  - [ ] If drawn card cannot be played, turn passes to other player
- [ ] Implement turn progression:
  - [ ] Turns alternate between the two players
  - [ ] Turn ends when a card is played (except for Skip/Reverse)
  - [ ] Turn ends when a drawn card cannot be played
- [ ] Implement UNO call rules:
  - [ ] Player must call "UNO" when down to one card
  - [ ] Penalty (draw two cards) if caught not calling "UNO"
- [ ] Implement win condition:
  - [ ] Game ends when a player has no cards left
  - [ ] First player to discard all cards wins

### Game Flow Management
- [ ] Implement turn sequence logic:
  - [ ] Start turn functionality
  - [ ] Play card validation and execution
  - [ ] Draw card functionality
  - [ ] Handle post-draw playability check
  - [ ] Handle special two-player turn sequences (Skip/Reverse)
  - [ ] End turn transitions
- [ ] Add UNO call handling:
  - [ ] Track when players should call UNO
  - [ ] Add button/mechanism for players to call UNO
  - [ ] Implement detection for missed UNO calls
  - [ ] Implement penalties for failing to call UNO (draw 2 cards)
- [ ] Implement special game states:
  - [ ] Color selection after Wild cards
  - [ ] Multiple plays after Skip/Reverse
  - [ ] Waiting for opponent to draw cards
- [ ] Implement win condition checking:
  - [ ] Detect when a player has no cards left
  - [ ] Handle game end transitions
  - [ ] Calculate final score based on opponent's remaining cards
- [ ] Write integration tests for complete game flows
  - [ ] Test standard card play sequences
  - [ ] Test all special card effects in two-player context
  - [ ] Test proper UNO call handling
  - [ ] Test win conditions

## Phase 3: User Interface Foundation

### Basic UI Framework
- [ ] Create ui/screens.go file:
  - [ ] Define Screen interface with Update, Draw, HandleInput methods
  - [ ] Create base screen implementations
- [ ] Create ui/input.go file:
  - [ ] Implement mouse click detection
  - [ ] Add button interaction logic
  - [ ] Create card selection handling
- [ ] Create ui/render.go file:
  - [ ] Add basic rendering utilities
  - [ ] Implement text rendering functions
  - [ ] Create simple shape rendering (buttons, cards)
- [ ] Integrate UI framework with game state
- [ ] Write tests for UI component interactions

### Title and Rules Screen
- [ ] Implement TitleScreen struct:
  - [ ] Create layout with proper button positioning
  - [ ] Add game title and visual elements
  - [ ] Implement navigation button actions
- [ ] Implement RulesScreen struct:
  - [ ] Display formatted text for two-player UNO rules
  - [ ] Add back button functionality
- [ ] Test screen navigation and display

### Settings and Setup Screens
- [ ] Implement SettingsScreen struct:
  - [ ] Add toggle components for sound effects and fullscreen
  - [ ] Create player name input fields
  - [ ] Implement settings storage and retrieval
- [ ] Create ConnectionScreen struct:
  - [ ] Add host game option
  - [ ] Add join game option
  - [ ] Implement player name input
  - [ ] Add back button functionality
- [ ] Test settings persistence and UI interactions

## Phase 4: Gameplay UI Implementation

### Gameplay Screen Framework
- [ ] Implement GameplayScreen struct:
  - [ ] Define regions for all UI elements
  - [ ] Create placeholders for player hands
  - [ ] Add draw and discard pile areas
  - [ ] Position action buttons (UNO, End Turn)
  - [ ] Add turn indicator placement
- [ ] Handle basic gameplay screen transitions
- [ ] Create layout tests for different screen sizes

### Card Rendering
- [ ] Create CardRenderer struct in ui package:
  - [ ] Implement RenderCard method for front-facing cards
  - [ ] Create RenderCardBack method for face-down cards
  - [ ] Add methods for card animations (play, draw, shuffle)
- [ ] Implement UNO card design standards:
  - [ ] Number cards with proper color and centered number
  - [ ] Skip cards with international "no" symbol
  - [ ] Reverse cards with direction change arrows
  - [ ] Draw Two cards with +2 symbol
  - [ ] Wild cards with four-color quadrant design
  - [ ] Wild Draw Four cards with four-color design and +4 symbol
  - [ ] All cards should have the UNO logo/text on back design
- [ ] Add card state visualization:
  - [ ] Highlight effect for playable cards
  - [ ] Selection indicator for chosen card
  - [ ] Dimming effect for unplayable cards
  - [ ] Special effect for "UNO" status (last card)
- [ ] Implement layout options:
  - [ ] Scaling for different screen sizes
  - [ ] Positioning for hand, draw pile, and discard pile
  - [ ] Card fanning for player's hand
  - [ ] Stacking for draw pile
- [ ] Test card rendering:
  - [ ] Test all card types and colors
  - [ ] Test different states (playable, selected, etc.)
  - [ ] Test animations
  - [ ] Test layout with different numbers of cards

### Hand and Game Elements
- [ ] Implement player hand rendering:
  - [ ] Create fanned card layout for player's own hand
  - [ ] Add card back rendering for opponent's hand
  - [ ] Implement card highlighting for selection
- [ ] Add game element rendering:
  - [ ] Draw pile visualization
  - [ ] Discard pile showing top card
  - [ ] Current player indicator
  - [ ] Turn action buttons
- [ ] Implement color wheel for Wild cards
- [ ] Test rendering of all game elements

### Gameplay Interaction
- [ ] Implement card selection:
  - [ ] Add click detection for cards in hand
  - [ ] Create visual feedback for selected cards
  - [ ] Add validation against UNO rules
  - [ ] Highlight playable cards in hand
  - [ ] Show tooltip for why cards aren't playable
- [ ] Implement standard game actions:
  - [ ] Draw card functionality
  - [ ] Play card execution
  - [ ] UNO button functionality (must be clicked when down to one card)
  - [ ] End Turn button actions
  - [ ] Automatic turn end after playing last card
- [ ] Implement special card interactions:
  - [ ] Create color wheel display for Wild cards
  - [ ] Implement color choice handling
  - [ ] Handle consecutive plays after Skip/Reverse
  - [ ] Manage Draw Two and Wild Draw Four effects
  - [ ] Validate Wild Draw Four plays (no matching color rule)
- [ ] Add player feedback:
  - [ ] Visual/audio cues for invalid moves
  - [ ] Highlight current player's turn
  - [ ] Countdown timer for UNO call
  - [ ] Animation for card draws and plays
  - [ ] Notification for opponent actions
- [ ] Test all gameplay interactions:
  - [ ] Test standard card plays
  - [ ] Test all special card effects
  - [ ] Test UNO call mechanic
  - [ ] Test appropriate feedback for all actions

### Results Screen
- [ ] Implement ResultsScreen struct:
  - [ ] Create winner display
  - [ ] Add remaining cards counter
  - [ ] Implement game time display
  - [ ] Add Play Again and Back to Title buttons
- [ ] Track and format play time
- [ ] Add transitions to and from results screen
- [ ] Test results screen functionality

## Phase 5: Networking Implementation

### WebSocket Server
- [ ] Create net/server.go file:
  - [ ] Implement GameServer struct
  - [ ] Add connection listening and handling
  - [ ] Create client tracking
  - [ ] Add game session management
- [ ] Implement server methods:
  - [ ] Start and stop server functionality
  - [ ] Handle incoming connections
  - [ ] Broadcast game state updates
- [ ] Add error handling and logging
- [ ] Test server functionality

### WebSocket Client
- [ ] Create net/client.go file:
  - [ ] Implement GameClient struct
  - [ ] Add server connection management
  - [ ] Create message sending and receiving
  - [ ] Implement reconnection logic
- [ ] Add client methods:
  - [ ] Connect and disconnect functionality
  - [ ] Send game actions to server
  - [ ] Receive and process game state updates
- [ ] Add connection status tracking
- [ ] Test client connectivity

### Network Protocol
- [ ] Create net/protocol.go file:
  - [ ] Define message types for all UNO game actions
  - [ ] Create message structures for data payloads
  - [ ] Implement serialization functions
  - [ ] Add message validation
- [ ] Implement protocol for standard UNO actions:
  - [ ] Player joining/leaving
  - [ ] Game start/initialization
  - [ ] Card playing
  - [ ] Drawing cards
  - [ ] Calling UNO
  - [ ] Ending turns
  - [ ] Selecting colors for Wild cards
- [ ] Implement protocol for special cases:
  - [ ] Wild Draw Four challenge (optional rule)
  - [ ] Penalty for missed UNO call
  - [ ] Consecutive plays after Skip/Reverse
  - [ ] Game state synchronization
  - [ ] Player timeout handling
- [ ] Implement security features:
  - [ ] Action validation (ensure only valid moves are allowed)
  - [ ] Turn validation (ensure players act only on their turn)
  - [ ] State validation (prevent inconsistencies)
  - [ ] Anti-cheat measures (server authority for rules)
- [ ] Test protocol:
  - [ ] Test serialization/deserialization
  - [ ] Test message validation
  - [ ] Test handling of all game actions
  - [ ] Test error cases and edge conditions

### LAN Discovery
- [ ] Create net/discovery.go file:
  - [ ] Implement DiscoveryServer struct for broadcasting
  - [ ] Create DiscoveryClient struct for finding games
  - [ ] Add GameBroadcast structure
- [ ] Implement discovery methods:
  - [ ] Start and stop broadcasting
  - [ ] Listen for available games
  - [ ] Track and update game listings
- [ ] Add UDP broadcasting and listening
- [ ] Test game discovery on LAN

### Network Integration
- [ ] Update GameState for networked play:
  - [ ] Add GameMode tracking (local vs network)
  - [ ] Implement network status reporting
- [ ] Modify UI for network status:
  - [ ] Show connection information
  - [ ] Display opponent details
  - [ ] Add waiting indicators
- [ ] Handle networked gameplay:
  - [ ] Process network events
  - [ ] Validate actions against game state
  - [ ] Synchronize state after actions
- [ ] Implement disconnection handling
- [ ] Test complete networked gameplay

## Phase 6: Polish and Finalization

### Audio System
- [ ] Create audio/audio.go file:
  - [ ] Implement AudioManager struct
  - [ ] Define sound effect types
  - [ ] Add audio loading and playback
- [ ] Implement sound triggers for:
  - [ ] Playing cards
  - [ ] Drawing cards
  - [ ] Calling UNO
  - [ ] Winning games
  - [ ] Special card effects
  - [ ] Invalid moves
- [ ] Add background music playback
- [ ] Implement volume control from settings
- [ ] Test audio system functionality

### UI Polish
- [ ] Implement screen transitions:
  - [ ] Add fade in/out effects
  - [ ] Create slide transitions
- [ ] Add card animations:
  - [ ] Card dealing animation
  - [ ] Card play movement
  - [ ] Card draw animation
- [ ] Implement UI feedback:
  - [ ] Button hover effects
  - [ ] Click feedback
  - [ ] Game event highlights
- [ ] Refine visual styling:
  - [ ] Consistent color scheme
  - [ ] Improved layouts
  - [ ] Visual hierarchy
- [ ] Test animations and transitions

### Performance Optimization
- [ ] Optimize rendering:
  - [ ] Implement sprite batching
  - [ ] Add asset caching
  - [ ] Reduce unnecessary redraws
- [ ] Improve network performance:
  - [ ] Optimize message sizes
  - [ ] Add compression if needed
  - [ ] Reduce update frequency
- [ ] Profile and optimize:
  - [ ] Memory usage
  - [ ] CPU utilization
  - [ ] Network traffic
- [ ] Test performance under various conditions

### Final Testing
- [ ] Implement comprehensive unit tests:
  - [ ] Card validation rules
  - [ ] Game state transitions
  - [ ] Special card effects
  - [ ] Winning conditions
- [ ] Add integration tests:
  - [ ] Game state and UI
  - [ ] Network communication
  - [ ] Complete game flows
- [ ] Perform manual testing:
  - [ ] UI responsiveness
  - [ ] Visual elements
  - [ ] Network performance
  - [ ] Gameplay experience
- [ ] Fix identified bugs
- [ ] Test on multiple platforms

### Documentation and Deployment
- [ ] Complete code documentation:
  - [ ] Add package documentation
  - [ ] Document functions and methods
  - [ ] Include examples where helpful
- [ ] Update README:
  - [ ] Add installation instructions
  - [ ] Include game rules
  - [ ] Explain controls and UI
  - [ ] Create network setup guide
- [ ] Create build scripts:
  - [ ] Windows executable
  - [ ] macOS application
  - [ ] Linux package
- [ ] Package assets with executables
- [ ] Test cross-platform functionality

## Final Checklist

### Verify Core Functionality
- [ ] Cards render correctly
- [ ] Game rules work properly
- [ ] Special card effects function as expected
- [ ] Win conditions trigger appropriately
- [ ] All screens navigate correctly

### Verify Network Play
- [ ] Games can be hosted and joined
- [ ] LAN discovery finds available games
- [ ] Game state synchronizes correctly
- [ ] Disconnection handling works properly
- [ ] Actions validate on server side

### Verify User Experience
- [ ] UI is intuitive and responsive
- [ ] Sound effects play appropriately
- [ ] Animations are smooth
- [ ] Error messages are clear
- [ ] Game flow is enjoyable

### Verify Cross-Platform
- [ ] Game runs on Windows
- [ ] Game runs on macOS
- [ ] Game runs on Linux
- [ ] Assets load correctly on all platforms
- [ ] Performance is acceptable everywhere

### Project Documentation
- [ ] All code is well-documented
- [ ] README contains all necessary information
- [ ] Build instructions are clear
- [ ] Game rules are well-explained
- [ ] Network setup is documented
