# Snake & Ladder Game

A command-line implementation of the classic Snake and Ladder board game written in **Go**. This is a low-level design (LLD) project demonstrating game logic, player management, and turn-based gameplay.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation & Setup](#installation--setup)
- [Building the Project](#building-the-project)
- [How to Run](#how-to-run)
- [How to Play](#how-to-play)
- [Project Structure](#project-structure)
- [Game Rules](#game-rules)
- [Ladders & Snakes](#ladders--snakes)
- [Contributing Guidelines](#contributing-guidelines)

## Prerequisites

- **Go 1.24.1** or later (see [go.dev](https://go.dev/dl) to download)
- macOS, Linux, or Windows with a terminal/command prompt

### Verify Go Installation

```bash
go version
```

Should output something like: `go version go1.25.6 darwin/arm64`

## Installation & Setup

### Step 1: Clone or Navigate to the Repository

```bash
cd /path/to/snakeladder
```

### Step 2: Download Dependencies

The project uses Go's standard library only, so dependencies are minimal:

```bash
go mod download
```

### Step 3: Verify Project Structure

Ensure you have the following files in the directory:

```
snakeladder/
├── main.go           # Main game loop and logic
├── board.go          # Ladder and snake definitions
├── dice.go           # Dice rolling logic
├── go.mod            # Go module file
├── README.md         # This file
├── game/
│   └── board.go      # Alternative board generation (unused)
└── player/
    └── player.go     # Player data structure
```

## Building the Project

### Option 1: Build to Binary (Recommended)

```bash
go build -o snakeladder
```

This creates an executable file named `snakeladder` that you can run directly.

### Option 2: Just Compile

```bash
go build
```

## How to Run

### Option 1: Direct Run (Fastest)

```bash
go run .
```

This compiles and runs the game in one command.

### Option 2: Run the Compiled Binary

```bash
./snakeladder
```

(After building with `go build`)

### Option 3: Run with Specific Files

```bash
go run main.go board.go dice.go
```

## How to Play

### Starting the Game

1. Enter **Player 1 name**
2. Enter **Player 2 name**
3. Press **Enter** to start

### During Gameplay

- **Roll Dice**: Press **`r`** and hit **Enter** to roll
- **Dice Range**: Rolls a random number between 1-6
- **Move**: Your position advances by the rolled number
- **Win**: First player to reach position **100** wins!

### Game Controls

| Command | Action |
|---------|--------|
| `r` | Roll the dice and move |
| `q` | Quit game (main menu only) |

### Example Game Flow

```
Who is player 1: Sagar
Who is player 2: Dhruv
=================== starting the game ===================
Current player: Sagar is on 0
Play, it is your turn: r
You rolled: 4
Sagar is now on position 4

Current player: Dhruv is on 0
Play, it is your turn: r
You rolled: 5
Dhruv is now on position 5
(Game continues until someone reaches 100)
```

## Project Structure

### Core Files

- **[main.go](main.go)** - Game loop, player turns, and game flow
- **[board.go](board.go)** - Ladder and snake position definitions
- **[dice.go](dice.go)** - Random dice rolling (1-6)
- **[go.mod](go.mod)** - Go module configuration

### Game Packages

- **[player/player.go](player/player.go)** - Player struct and initialization
- **[game/board.go](game/board.go)** - Alternative board generation logic (currently unused)

## Game Rules

### Objective

Be the first player to reach position **100** on the board.

### Turn System

- Players take turns in order (Player 1, Player 2, etc.)
- Each turn: roll dice → move → check for snakes/ladders

### Movement

1. Roll the dice (generates 1-6)
2. Move forward by that number
3. If you land on a **ladder start**, climb up
4. If you land on a **snake head**, slide down
5. Otherwise, end your turn

### Winning

- First player to reach exactly position **100** wins
- Game announces the winner and asks if players want to play again

## Ladders & Snakes

### Ladders (Climb Up)

| Start → End |
|-------------|
| 1 → 38 |
| 4 → 14 |
| 9 → 31 |
| 21 → 42 |
| 28 → 84 |
| 51 → 67 |
| 72 → 91 |
| 80 → 99 |

### Snakes (Slide Down)

| Start → End |
|-------------|
| 17 → 7 |
| 62 → 19 |
| 64 → 60 |
| 54 → 34 |
| 87 → 36 |
| 93 → 73 |
| 95 → 75 |
| 98 → 79 |

## Contributing Guidelines

### Code Style

- Follow Go's standard conventions (use `gofmt`)
- Write clear, descriptive function and variable names
- Add comments for non-obvious logic

### Format Code

```bash
go fmt ./...
```

### Lint Code

```bash
go vet ./...
```

### Best Practices

1. Keep functions small and focused
2. Use meaningful variable names
3. Add error handling where needed
4. Document public functions and packages
5. Test changes before committing

### Pull Request Process

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/your-feature`)
3. Make your changes
4. Run `go fmt` and `go vet`
5. Commit with clear messages
6. Push and create a pull request

## Troubleshooting

### "go: command not found"

Install Go from [go.dev](https://go.dev/dl)

### "module not found"

Run `go mod download` in the project directory

### Game won't compile

Ensure all files are in the correct directory and run `go build`

## License

This is a learning project. Feel free to use and modify as needed.

---

**Happy Playing!** 🎮
