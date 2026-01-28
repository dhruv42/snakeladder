package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dhruv42/snakeladder/cmd"
	"github.com/dhruv42/snakeladder/game"
	"github.com/dhruv42/snakeladder/player"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	cmd.Banner()
	input := ""
	for input != "q" {
		fmt.Print("Who is player 1: ")
		name1 := cmd.Takeinput()
		fmt.Println()

		fmt.Print("Who is player 2: ")
		name2 := cmd.Takeinput()
		fmt.Println()

		allPlayers := player.New(name1, name2)
		board := game.New(10, 8, 8)

		play(
			allPlayers,
			&AllLadders{Ladders: board.Ladders},
			&AllSnakes{Snakes: board.Snakes},
			NewDice())
	}
}

func play(ap *player.AllPlayers, l *AllLadders, s *AllSnakes, d *Dice) {
	currPos := 0
	fmt.Print("=================== starting the game ===================\n\n")

	for currPos != 100 {
		for _, cp := range ap.Players {
			displayPlayerPosition(cp)
			if cmd.Takeinput() != "r" {
				continue
			}

			dicenumber := d.Roll()
			if cp.CurrentPosition+dicenumber > 100 {
				continue
			}
			cp.CurrentPosition += dicenumber
			haswon := displayFinalPosition(cp)
			if haswon {
				fmt.Printf("The winner is: %s\n", cp.Name)
				return
			}

			snakebit, spos := checksnake(s, cp)
			if snakebit {
				cp.CurrentPosition = spos
				haswon = fallingDown(cp)
				if haswon {
					fmt.Printf("The winner is: %s\n", cp.Name)
					return
				}
				currPos = cp.CurrentPosition
				fmt.Println("######################################################\n")
				continue
			}

			gotladder, lpos := checkladder(l, cp)
			if gotladder {
				cp.CurrentPosition = lpos
				haswon = climbingTheLadder(cp)
				if haswon {
					fmt.Printf("The winner is: %s\n", cp.Name)
					return
				}
				currPos = cp.CurrentPosition
				fmt.Println("######################################################\n")
				continue
			}
			fmt.Println("######################################################\n")
		}
	}
}

func checksnake(snakes *AllSnakes, p *player.Player) (bool, int) {
	for _, s := range snakes.Snakes {
		if s.Start == p.CurrentPosition {
			return true, s.End // if snake bites then return snake end postition
		}
	}
	return false, 0
}

func checkladder(ladders *AllLadders, p *player.Player) (bool, int) {
	for _, l := range ladders.Ladders {
		if l.Start == p.CurrentPosition {
			return true, l.End // if gets ladder then return ladder end postition
		}
	}
	return false, 0
}

func displayPlayerPosition(cp *player.Player) {
	fmt.Printf("%s, you are on (%d), it's your turn: ", cp.Name, cp.CurrentPosition)
}

func displayFinalPosition(cp *player.Player) bool {
	fmt.Printf("%s, you landed on (%d)\n\n", cp.Name, cp.CurrentPosition)
	return cp.CurrentPosition == 100
}

func climbingTheLadder(cp *player.Player) bool {
	fmt.Printf("%s, guess who got lucky 🥳 to climb 🪜, reaching at .... (%d)\n\n", cp.Name, cp.CurrentPosition)
	return cp.CurrentPosition == 100
}

func fallingDown(cp *player.Player) bool {
	fmt.Printf("%s, oppsiee whopssieee 😿, you got a 🐍 bite, coming back stronger from position ..... (%d)\n\n", cp.Name, cp.CurrentPosition)
	return cp.CurrentPosition == 100
}
