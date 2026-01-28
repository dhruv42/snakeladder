package main

import (
	"fmt"
	"math/rand/v2"
)

type Dice struct {
	Curr int
}

func NewDice() *Dice {
	return &Dice{Curr: 0}
}

func (d *Dice) Roll() int {
	newcurr := rand.IntN(7-1) + 1
	d.set(newcurr)
	fmt.Printf("Rollingggggg the dice........ and it is: %d\n", newcurr)
	return newcurr // because last number is exclusive in the range
}

func (d *Dice) set(n int) {
	d.Curr = n
}

func (d *Dice) get() int {
	return d.Curr
}
