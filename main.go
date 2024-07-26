package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var figure string
	flag.StringVar(&figure, "f", "cow", "the figure name. Valid values are "+allFigures(animals))
	flag.Parse()

	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	var lines []string
	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(input))
	}

	lines = tabsToSpaces(lines)
	width := calculateMaxWidth(lines)
	lines = normalizeStringsLength(lines, width)
	balloon := buildBalloon(lines, width)
	fmt.Println(balloon)
	printFigure(figure)
	fmt.Println("")
}
