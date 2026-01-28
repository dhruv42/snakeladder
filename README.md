<h2 style="border-bottom:none;">Snake &amp; Ladder Game</h2>
<p><em>(Developed by human (not AI) with love 💙)</em></p>

---

A command-line implementation of the classic Snake and Ladder board game written in **Go**. This is a low-level design (LLD) project demonstrating game logic, player management, and turn-based gameplay.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation & Setup](#installation--setup)
- [How to Run](#how-to-run)
- [How to Play](#how-to-play)
- [Game Rules](#game-rules)
- [Ladders & Snakes](#ladders--snakes)

## Prerequisites

- **Go 1.24.1**

### Verify Go Installation

```bash
go version
```

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

## How to Run

### Direct Run

```bash
go run .
```

This compiles and runs the game in one command.

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
Who is player 1: Player A
Who is player 2: Player B
=================== starting the game ===================
Current player: Player A is on 0
Play, it is your turn: r
You rolled: 4
Player A is now on position 4

Current player: Player B is on 0
Play, it is your turn: r
You rolled: 5
Player B is now on position 5
(Game continues until someone reaches 100)
```

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

---

**Happy Playing!** 🎮
