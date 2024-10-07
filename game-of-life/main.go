package main

import (
	"fmt"
	"math/rand"
	"time"
)

const width = 20
const height = 20

var board [][]uint8
var buffer [][]uint8

// TODO: draw the board. Could use
func update_board_old() error { // [][]uint8

	return nil
}

func board_init() {
	board = make([][]uint8, width)
	buffer = make([][]uint8, width)
	for i := range board{
		board[i] = make([]uint8, height)
		buffer[i] = make([]uint8, height)
	}
}

func board_init_random(b [][]uint8){
	for x := 1; x < width-1; x++ {
		for y := 1 ; y < height-1; y++ {
			if rand.Float32() < 0.5 {
				b[x][y] = 1
			} else {
				b[x][y] = 0
			}
		}
	}
}

func update_board(b [][]uint8, buffer [][]uint8) error {
	for x := 1; x < len(b) - 1; x++{
		for y := 1; y < len(b[0]) - 1; y++{
			n := b[x-1][y-1] + b[x-1][y] + b[x][y-1] + b[x-1][y+1] + b[x+1][y+1] + b[x+1][y] + b[x][y+1]

			if b[x][y] == 0 && n == 3 {
				buffer[x][y] = 1
			} else if n < 2 || n > 3 {
				buffer[x][y] =  0
			} else {
				buffer[x][y] = b[x][y]
			}
		}
	}
	// Make buffer the current grid
	for x := 0; x < len(b); x++ {
		for y := 0; y < len(b[0]); y++ {
			b[x][y] = buffer[x][y]
		}
	}


	return nil	
}
func show_board(b [][]uint8){

	for _, s := range b{
		fmt.Println(s)
	}
	fmt.Println("\n")
}


func main() {

	board_init()
	board_init_random(board)
	for true{
		update_board(board, buffer)
		show_board(board)
		time.Sleep(1 * time.Second)
	}
	
}
