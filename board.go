package main

import "github.com/dhruv42/snakeladder/game"

// type Board struct {
// 	Size    int
// 	Ladders []Ladder
// 	Snakes  []Snake
// }

type AllLadders struct {
	Ladders []game.Ladder
}

type AllSnakes struct {
	Snakes []game.Snake
}

// type Ladder struct {
// 	Start, End int
// }

// type Snake struct {
// 	Start, End int
// }

// func InitializeBoard() (*AllLadders, *AllSnakes, *Dice) {
// 	ladders := &AllLadders{
// 		Ladders: []Ladder{
// 			{Start: 1, End: 38},
// 			{Start: 4, End: 14},
// 			{Start: 9, End: 31},
// 			{Start: 21, End: 42},
// 			{Start: 28, End: 84},
// 			{Start: 51, End: 67},
// 			{Start: 72, End: 91},
// 			{Start: 80, End: 99},
// 		},
// 	}

// 	snakes := &AllSnakes{
// 		Snakes: []Snake{
// 			{Start: 17, End: 7},
// 			{Start: 62, End: 19},
// 			{Start: 64, End: 60},
// 			{Start: 54, End: 34},
// 			{Start: 87, End: 36},
// 			{Start: 93, End: 73},
// 			{Start: 95, End: 75},
// 			{Start: 98, End: 79},
// 		},
// 	}

// 	return ladders, snakes, NewDice()
// }
