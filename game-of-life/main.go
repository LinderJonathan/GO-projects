package main

import (
	"fmt"
	"math/rand"
)

const width = 300
const height = 300

var board [][]uint16 = [][]uint16{}
var buffer [][]uint16 = [][]uint16{}

// TODO: draw the board. Could use
func update_board() error { // [][]uint8

	return nil
}
func board_init(m [][]uint16) [][]uint16 {
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if rand.Float32() < 0.5 {
				m[x][y] = 1
			} else {
				m[x][y] = 0
			}

		}
	}
	return m
}

var temp = board_init(board)

func main() {
	fmt.Println("hello, world")
}
