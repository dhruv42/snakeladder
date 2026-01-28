package game

import (
	"fmt"
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

	return &Board{
		Size:         size,
		TotalSquares: totalSquares,
		Ladders:      _GenLadders(size, totalSquares, totalLadders, numbers),
		Snakes:       _GenSnakes(totalSquares, totalSnakes),
	}
}

func _GenLadders(boardsize, totalSquares int, total int, n []int) []Ladder {
	l := []Ladder{}
	min := 1

	for i := 0; i < total; i++ {
		startIdx := rand.IntN((len(n)-1)/2-min) + min
		start := n[startIdx]
		remove(n, startIdx)

		endIdx := rand.IntN((len(n)-1)-start) + start
		end := n[endIdx]
		remove(n, endIdx)
		l = append(l, Ladder{
			Start: start,
			End:   end,
		})
	}

	// for i := 0; i < total; i++ {
	// 	start := rand.IntN((totalSquares-boardsize)-min) + min

	// 	for _, ok := excludeNumbersMap[start]; ok; {
	// 		start = rand.IntN((totalSquares-boardsize)-min) + min
	// 		fmt.Println("ladders start loop -----", start)
	// 	}
	// 	excludeNumbersMap[start] = true

	// 	end := rand.IntN(totalSquares-(start+1)) + (start + 1)
	// 	// loop until we get the number which is not been used already
	// 	for _, ok := excludeNumbersMap[end]; ok; {
	// 		end = rand.IntN(totalSquares-(start+1)) + (start + 1)
	// 		fmt.Println("ladders end loop -----", end)
	// 	}
	// 	excludeNumbersMap[end] = true
	// 	fmt.Println("ladders ---", start, end)
	// 	l = append(l, Ladder{
	// 		Start: start,
	// 		End:   end,
	// 	})
	// }
	fmt.Println(excludeNumbersMap)
	return l
}

func _GenSnakes(totalSquares int, total int) []Snake {
	s := []Snake{}
	min := 1

	// excludeNumbersMap := make(map[int]bool)

	for i := 0; i < total; i++ {
		start := rand.IntN(totalSquares-min) + min
		for _, ok := excludeNumbersMap[start]; ok; {
			start = rand.IntN(totalSquares-min) + min
		}
		excludeNumbersMap[start] = true

		end := rand.IntN(totalSquares-start) + start
		// loop until we get the number which is not been used already
		for _, ok := excludeNumbersMap[end]; ok || start == end; {
			end = rand.IntN(totalSquares-start) + start
		}
		excludeNumbersMap[end] = true
		// swapping because snake's start is head and end is tail, start will be always
		// greater than the tail
		start, end = end, start

		fmt.Println("snakes ---", start, end)
		s = append(s, Snake{
			Start: start,
			End:   end,
		})
	}
	return s
}

func remove(n []int, i int) []int {
	n[i] = n[len(n)-1]
	return n[:len(n)-1]
}
