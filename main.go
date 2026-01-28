package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dhruv42/snakeladder/player"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println("text ---", text)

	var input string
	for input != "q" {
		fmt.Print("Who is player 1: ")
		name1 := takeinput()
		fmt.Println()

		fmt.Print("Who is player 2: ")
		name2 := takeinput()

		allPlayers := player.New(name1, name2)
		// board := game.New(10, 8, 8)
		// temp, _ := json.Marshal(*board)
		// fmt.Println("board ------", string(temp))
		allLadders, allSnakes, dice := InitializeBoard()

		play(allPlayers, allLadders, allSnakes, dice)
		// break
	}

}

func takeinput() string {
	scanner.Scan()
	return scanner.Text()
}

func play(ap *player.AllPlayers, l *AllLadders, s *AllSnakes, d *Dice) {
	currPos := 0
	fmt.Println("=================== starting the game ===================")
	for currPos != 100 {
		for _, cp := range ap.Players {
			// cp := cp
			// fmt.Println("--------------------")
			displayPlayerPosition(cp)

			fmt.Print("Play, it is your turn: ")
			if takeinput() != "r" {
				continue
			}

			dicenumber := d.Roll()
			if cp.CurrentPosition + dicenumber > 100 {
				continue
			}
			cp.CurrentPosition += dicenumber
			displayFinalPosition(cp)

			if cp.CurrentPosition == 100 {
				fmt.Printf("The winner is: %s\n", cp.Name)
				return
			}

			snakebit, spos := checksnake(s, cp)
			if snakebit {
				cp.CurrentPosition = spos
				displayFinalPosition(cp)
				if cp.CurrentPosition == 100 {
        			fmt.Printf("The winner is: %s\n", cp.Name)
        			return
    			}
				currPos = cp.CurrentPosition
				continue
			}

			gotladder, lpos := checkladder(l, cp)
			if gotladder {
				cp.CurrentPosition = lpos
				displayFinalPosition(cp)
				if cp.CurrentPosition == 100 {
        			fmt.Printf("The winner is: %s\n", cp.Name)
       				return
    			}
				currPos = cp.CurrentPosition
				continue
			}
			fmt.Println("--------------------")
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
	fmt.Printf("Current player: %s is on %d\n", cp.Name, cp.CurrentPosition)
}

func displayFinalPosition(cp *player.Player) {
	fmt.Printf("%s reaches on %d\n", cp.Name, cp.CurrentPosition)
}
