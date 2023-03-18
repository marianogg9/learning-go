package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	const ( // set constant values for rows and columns
		width     = 50
		height    = 10
		cellEmpty = ' '
		cellBall  = '⚽'
		maxFrames = 1200
		speed     = time.Second / 20
	)

	var (
		cell   rune   // to use as a placeholder for the ball (unicode)
		px, py int    // vertical/horizontal position of the ball
		vx, vy = 1, 1 // velocity of the ball: 1 means the ball moves in a positive direction 1 step at a time (vertically y or horizontall x), -1 means it moves negatively.
	)

	board := make([][]bool, width) // define the board (two dimensional boolean slice)
	for row := range board {
		board[row] = make([]bool, height) // initialize the board
	}

	screen.Clear() // clear the screen to start the animation

	for i := 0; i < maxFrames; i++ { // create this animation maxFrames amount of times = 60 seconds, 20 frames per second
		px += vx // the position will be defined by its velocity - if velocity is 1, the position will increase by 1. If velocity is -1, then the position decreases.
		py += vy

		if px <= 0 || px >= width-1 { // if the ball hits its horizontal edges, revert its horizontal velocity
			vx *= -1
		}

		if py <= 0 || py >= height-1 { // if the ball hits its vertical edges, revert its vertical velocity
			vy *= -1
		}

		for y := range board[0] { // reinitialize the board every time by reseting all values to false (no content)
			for x := range board {
				board[x][y] = false
			}
		}

		board[px][py] = true
		buf := make([]rune, 0, width*height) // declare the buffer -> it will look like: buf = [' ', ' ', 'ball', ' ',..,'\n',..,' ']
		buf = buf[:0]                        // empty the buffer each time

		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				// fmt.Print(string(cell), " ") // print the ball // given print is an expensive operation, using a buffer to accumulate values.
				buf = append(buf, cell, ' ')
			}
			// fmt.Println()
			buf = append(buf, '\n')
		}
		screen.MoveTopLeft()   // clear the screen everytime so it prints each new board in the same position, making the animation work correctly
		fmt.Print(string(buf)) // print the board with the ball
		time.Sleep(speed)      // make the animation slower
	}
}
