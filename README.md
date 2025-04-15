# Uno Clone - Under Construction

A two-player networked implementation of the classic UNO card game built with Go and Ebiten.

## Overview

This project is a faithful recreation of the classic UNO card game with custom pixel art graphics. It was creating for learning purposes. This implementation is specifically designed for two players over a local network, with special rules adapted for two-player gameplay.

## Features

- **Two-Player Gameplay**: Optimized rules for head-to-head UNO matches
- **Custom Pixel Art**: Unique visual style with hand-crafted pixel art cards and UI
- **LAN Multiplayer**: Play with a friend on your local network
- **Simple Discovery**: Easily find and join games without entering IP addresses
- **Full Game Experience**: All the classic UNO cards and special effects

## Getting Started

### Prerequisites

- Go 1.18 or later
- Network connection (for multiplayer)

### Installation

1. Clone the repository:
```
git clone https://github.com/vtigo/uno-vlone.git
cd uno-clone
```

2. Build the game:
```
go build -o uno-clone
```

3. Run the game:
```
./uno-clone
```

On Windows, run `uno-clone.exe`

## How to Play

1. **Start the game** and choose "Play Game"
2. Select whether to **Host** or **Join** a game
   - If hosting, wait for another player to join
   - If joining, the game will automatically discover available hosts on your network
3. Once connected, both players will be dealt 7 cards
4. Take turns playing cards that match the top card on the discard pile by color or number
5. Special cards have unique effects:
   - **Skip**: Skip the opponent's turn (you play again)
   - **Reverse**: Acts like Skip in two-player mode (you play again)
   - **Draw Two**: Opponent draws 2 cards and loses their turn
   - **Wild**: Choose any color
   - **Wild Draw Four**: Opponent draws 4 cards, loses their turn, and you choose the color
6. Don't forget to click the **UNO** button when you have only one card left!
7. First player to play all their cards wins

## Two-Player Special Rules

- Playing a Reverse card acts like a Skip. The player who plays the Reverse may immediately play another card.
- The person playing a Skip card may immediately play another card.
- When one person plays a Draw Two card and the other player has drawn the 2 cards, the play is back to the first person. The same principle applies to the Wild Draw Four card.

## Development

This project uses:
- [Go](https://golang.org/) as the programming language
- [Ebiten](https://ebiten.org/) as the game engine
- [Gorilla WebSocket](https://github.com/gorilla/websocket) for network communication

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the classic UNO card game by Mattel
- Special thanks to the Ebiten project and Go community
