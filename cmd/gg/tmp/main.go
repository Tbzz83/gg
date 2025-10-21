package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	text := "hello world"
	cursor := 5

	blockStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("12")).
		Foreground(lipgloss.Color("0"))

	for {
		fmt.Print("\033[H\033[2J") // clear screen

		// Draw text with block cursor
		var builder strings.Builder
		builder.WriteString(text[:cursor])
		if cursor < len(text) {
			builder.WriteString(blockStyle.Render(string(text[cursor])))
			builder.WriteString(text[cursor+1:])
		} else {
			builder.WriteString(blockStyle.Render(" "))
		}

		fmt.Println(builder.String())
		time.Sleep(200 * time.Millisecond)
	}
}

