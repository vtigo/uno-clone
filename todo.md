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
- [X] Define Card struct and related types in game/card.go:
  - [X] CardColor enum (Red, Blue, Green, Yellow, Wild)
  - [X] CardType enum (Number, Skip, Reverse, DrawTwo, WildCard, WildDrawFour)
  - [X] Card struct with Color, Type, and Value fields
- [X] Implement deck management functions:
  - [X] CreateDeck function that generates all 108 cards
  - [X] Shuffle function to randomize deck
  - [X] DrawCard function to take cards from the deck
- [X] Write unit tests for card and deck functionality

### Player Implementation
- [ ] Create Player struct in game/player.go:
  - [ ] Name field (string)
  - [ ] Hand field (slice of Card pointers)
  - [ ] Any additional player state fields
- [ ] Implement player methods:
  - [ ] NewPlayer function to create player with name
  - [ ] AddCard method to add cards to hand
  - [ ] PlayCard method to remove/play cards from hand
  - [ ] HasWon method to check for win condition (no cards)
  - [ ] ShouldCallUno method to check if player has one card
- [ ] Write unit tests for player functionality

### Game State Management
- [ ] Create GameState struct in game/state.go:
  - [ ] GameScreen enum for different screens
  - [ ] CurrentScreen field to track current screen
  - [ ] Players array for both players
  - [ ] CurrentPlayer index to track active player
  - [ ] Deck field for remaining cards
  - [ ] DiscardPile field for played cards
  - [ ] CurrentCard field for top card
- [ ] Implement game state methods:
  - [ ] NewGameState function to initialize game
  - [ ] InitGame method to set up a new game
  - [ ] DealCards method to distribute initial hands
  - [ ] NextTurn method for turn management
  - [ ] ResetGame method to start over
- [ ] Write unit tests for game state functionality

## Phase 2: Game Logic and Rules

### Card Validation and Rules
- [ ] Create game/rules.go file:
  - [ ] Implement IsValidPlay function to check if a card can be played
  - [ ] Create HandleCardEffect function for special card effects
- [ ] Implement special card effects:
  - [ ] Skip card logic (current player plays again)
  - [ ] Reverse card logic (acts as Skip in 2-player)
  - [ ] Draw Two card logic
  - [ ] Wild card logic with color selection
  - [ ] Wild Draw Four card logic
- [ ] Implement two-player specific rules:
  - [ ] Reverse and Skip allow current player to play again
  - [ ] Draw cards affect turn appropriately
- [ ] Write unit tests for all validation and card effects

### Game Flow Management
- [ ] Implement turn sequence logic:
  - [ ] Start turn functionality
  - [ ] Play card validation and execution
  - [ ] Draw card functionality
  - [ ] End turn transitions
- [ ] Add UNO call handling:
  - [ ] Track when players should call UNO
  - [ ] Implement penalties for failing to call UNO
- [ ] Implement win condition checking:
  - [ ] Detect when a player has no cards left
  - [ ] Handle game end transitions
- [ ] Write integration tests for complete game flows

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
- [ ] Implement rendering for different card types:
  - [ ] Number cards with color and value
  - [ ] Skip cards with symbol
  - [ ] Reverse cards with symbol
  - [ ] Draw Two cards with symbol
  - [ ] Wild cards with symbol
  - [ ] Wild Draw Four cards with symbol
- [ ] Add scaling and positioning options
- [ ] Test card rendering for all card types

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
  - [ ] Add validation against game rules
- [ ] Implement game actions:
  - [ ] Draw card functionality
  - [ ] Play card execution
  - [ ] UNO button functionality
  - [ ] End Turn button actions
- [ ] Add Wild card color selection:
  - [ ] Create color wheel display
  - [ ] Implement color choice handling
- [ ] Add feedback for invalid moves
- [ ] Test all gameplay interactions

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
  - [ ] Define message types for all game actions
  - [ ] Create message structures for data payloads
  - [ ] Implement serialization functions
  - [ ] Add message validation
- [ ] Implement protocol for:
  - [ ] Player joining/leaving
  - [ ] Card playing
  - [ ] Drawing cards
  - [ ] Calling UNO
  - [ ] Ending turns
  - [ ] Selecting colors
  - [ ] Game state synchronization
- [ ] Test protocol serialization and validation

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
