package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	height := 25
	width := 80
	ballX, ballY := width/2, height/2
	leftPaddleY := height/2 - 1
	rightPaddleY := height/2 - 1
	paddleHeight := 3
	SballX, SballY := 1, 1
	game(height, width, ballX, ballY, leftPaddleY, rightPaddleY, paddleHeight, SballX, SballY)
}

func field(height, width, ballX, ballY, leftPaddleY, rightPaddleY, paddleHeight int) {
	borderChar := "#"
	emptyChar := " "
	ballChar := "O"
	paddleChar := "|"

	fmt.Println(strings.Repeat(borderChar, width))

	for y := 0; y < height-2; y++ {
		line := borderChar
		for x := 0; x < width-2; x++ {
			if x == ballX && y == ballY {
				line += ballChar
			} else if x == 1 && y >= leftPaddleY && y < leftPaddleY+paddleHeight {
				line += paddleChar
			} else if x == width-4 && y >= rightPaddleY && y < rightPaddleY+paddleHeight {
				line += paddleChar
			} else {
				line += emptyChar
			}
		}
		line += borderChar
		fmt.Println(line)
	}
	fmt.Println(strings.Repeat(borderChar, width))
}

func game(height, width, ballX, ballY, leftPaddleY, rightPaddleY, paddleHeight, SballX, SballY int) {
	for {
		fmt.Print("\033[H\033[2J")
		field(height, width, ballX, ballY, leftPaddleY, rightPaddleY, paddleHeight)
		ballX += SballX
		ballY += SballY
		if ballY <= 1 || ballY >= height-2 {
			SballY *= -1
		}
		if ballX <= 2 || ballX >= width-3 {
			SballX *= -1
		}
		time.Sleep(50 * time.Millisecond)
	}
}
