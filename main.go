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
	var showCount bool
	flag.StringVar(&figure, "f", "cow", "the figure name. Valid values are "+allFigures(animals))
	flag.BoolVar(&showCount, "c", false, "show the count of available animals and exit")
	flag.Parse()

	// 如果用户指定了 -c 标志，只输出动物数量并退出
	if showCount {
		fmt.Println(len(animals))
		os.Exit(0)
	}

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
