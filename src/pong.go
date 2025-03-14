package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/mattn/go-tty"
)

func main() {
	height := 25
	width := 80
	ballX, ballY := width/2, height/2
	leftPaddleY := height/2 - 1
	rightPaddleY := height/2 - 1
	paddleHeight := 3
	SballX, SballY := 1, 1
	player1Score, player2Score := 0, 0

	game(height, width, ballX, ballY, leftPaddleY, rightPaddleY, paddleHeight, SballX, SballY, &player1Score, &player2Score)
}

func field(height, width, ballX, ballY, leftPaddleY, rightPaddleY, paddleHeight, player1Score, player2Score int) {
	borderChar := "#"
	emptyChar := " "
	ballChar := "O"
	paddleChar := "|"

	fmt.Printf("Player 1: %d %s Player 2: %d\n", player1Score, strings.Repeat(" ", width-27), player2Score)
	fmt.Println(strings.Repeat(borderChar, width))

	for y := 0; y < height-3; y++ {
		line := borderChar
		for x := 0; x < width-2; x++ {
			if x == ballX && y == ballY {
				line += ballChar
			} else if x == 3 && y >= leftPaddleY && y < leftPaddleY+paddleHeight {
				line += paddleChar
			} else if x == width-6 && y >= rightPaddleY && y < rightPaddleY+paddleHeight {
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

func game(height, width, ballX, ballY, leftPaddleY, rightPaddleY, paddleHeight, SballX, SballY int, player1Score, player2Score *int) {
	tty, _ := tty.Open()
	defer tty.Close()

	for {
		fmt.Print("\033[H\033[2J")
		field(height, width, ballX, ballY, leftPaddleY, rightPaddleY, paddleHeight, *player1Score, *player2Score)
		ballX += SballX
		ballY += SballY
		if ballY <= 1 || ballY >= height-3 {
			SballY *= -1
		}
		if ballX == 3 && ballY >= leftPaddleY && ballY < leftPaddleY+paddleHeight {
			SballX *= -1
		}
		if ballX == width-6 && ballY >= rightPaddleY && ballY < rightPaddleY+paddleHeight {
			SballX *= -1
		}
		if ballX <= 1 {
			*player2Score++
			ballX, ballY, SballX, SballY = width/2, height/2, 1, 1
		} else if ballX >= width-2 {
			*player1Score++
			ballX, ballY, SballX, SballY = width/2, height/2, -1, -1
		}
		char, err := tty.ReadRune()
		if *player1Score == 11 {
			fmt.Println("Игрок 1 победил!")
			char = 'q'
		} else if *player2Score == 11 {
			fmt.Println("Игрок 2 победил!")
			char = 'q'
		}
		if err == nil {
			switch char {
			case 'q':
				return
			case 'a':
				if leftPaddleY > 1 {
					leftPaddleY--
				}
			case 'z':
				if leftPaddleY < height-paddleHeight-3 {
					leftPaddleY++
				}
			case 'k':
				if rightPaddleY > 1 {
					rightPaddleY--
				}
			case 'm':
				if rightPaddleY < height-paddleHeight-3 {
					rightPaddleY++
				}
			}
		}

		time.Sleep(50 * time.Millisecond)
	}
}
