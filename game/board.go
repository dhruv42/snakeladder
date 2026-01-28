package game

import (
	"math/rand/v2"
)

type Board struct {
	Size         int
	TotalSquares int
	Ladders      []Ladder
	Snakes       []Snake
}

type Ladder struct {
	Start, End int
}

type Snake struct {
	Start, End int
}

var excludeNumbersMap = make(map[int]bool)

func New(size, totalSnakes, totalLadders int) *Board {
	// Clear the exclude map for each new board generation
	excludeNumbersMap = make(map[int]bool)
	totalSquares := size * size
	numbers := []int{}
	for i := 1; i <= 100; i++ {
		numbers = append(numbers, i)
	}
	ladders, n := _GenLadders(size, totalLadders, numbers)
	snakes, n := _GenSnakes(size, totalSnakes, n)

	return &Board{
		Size:         size,
		TotalSquares: totalSquares,
		Ladders:      ladders,
		Snakes:       snakes,
	}
}

func _GenLadders(boardsize, total int, n []int) ([]Ladder, []int) {
	l := []Ladder{}

	// start -> choose any start index, let's keep max start as start - 10, means start can not be more than 90
	// end -> (end > start always)

	for i := 0; i < total; i++ {
		min := 0
		max := len(n) - 1 - boardsize
		startIdx := rand.IntN(max-min) + min
		start := n[startIdx]
		n = remove(n, startIdx)

		min = startIdx + boardsize
		endIdx := rand.IntN(len(n)-min) + min
		end := n[endIdx-1]
		n = remove(n, endIdx-1)
		l = append(l, Ladder{
			Start: start,
			End:   end,
		})
	}
	return l, n
}

func _GenSnakes(boardsize, total int, n []int) ([]Snake, []int) {
	s := []Snake{}
	for i := 0; i < total; i++ {
		min := 0
		max := len(n) - 1 - boardsize
		startIdx := rand.IntN(max-min) + min
		start := n[startIdx]
		n = remove(n, startIdx)

		min = startIdx + boardsize
		endIdx := rand.IntN(len(n)-min) + min
		end := n[endIdx-1]
		n = remove(n, endIdx-1)
		s = append(s, Snake{
			Start: end,
			End:   start,
		})
	}
	return s, n
}

func remove(n []int, i int) []int {
	return append(n[:i], n[i+1:]...) // this will not change the ordering of elements
	// if ordering is not important then use this
	// n[i] = n[len(n)-1]
	// return n[:len(n)-1]
}
