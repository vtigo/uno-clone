package main

// Game window configuration
const (
	ScreenWidth  = 1280
	ScreenHeight = 720
	WindowTitle  = "UNO Clone"
	TPS          = 60 // Ticks per second
)

// Game state constants
const (
	StateTitle int = iota
	StateRules
	StateSettings
	StateNetworkSetup
	StateGameplay
	StateResults
)

// Card constants
const (
	CardWidth  = 64
	CardHeight = 96
)

// Player constants
const (
	InitialHandSize = 7
	MaxPlayers      = 2
)

// Networking constants
const (
	DefaultPort         = "8080"
	DiscoveryPort       = "8081"
	BroadcastInterval   = 1000 // milliseconds
	DiscoveryTimeout    = 5000 // milliseconds
	ConnectionTimeout   = 30   // seconds
	DisconnectThreshold = 5    // missed heartbeats
)

// Asset paths
const (
	FontPath    = "assets/fonts/main.ttf"
	CardImgPath = "assets/images/cards/"
	AudioPath   = "assets/audio/"
)
