package main

import (
	"fmt"
	"strings"
)

func main() {
	height := 25
	width := 80
	field(height, width)
}

func field(height, width int) {
	borderChar := "#"
	emptyChar := " "
	fmt.Println(strings.Repeat(borderChar, width))
	for i := 0; i < height-2; i++ {
		fmt.Println(borderChar + strings.Repeat(emptyChar, width-2) + borderChar)
	}
	fmt.Println(strings.Repeat(borderChar, width))
}

func ball() {

}
