package main

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

func tabsToSpaces(lines []string) (ret []string) {
	for _, line := range lines {
		line = strings.Replace(line, "\t", "    ", -1)
		ret = append(ret, line)
	}
	return
}

func calculateMaxWidth(lines []string) (max int) {
	for _, line := range lines {
		len := runewidth.StringWidth(line)
		if len > max {
			max = len
		}
	}
	return
}

func normalizeStringsLength(lines []string, width int) (ret []string) {
	for _, line := range lines {
		len := runewidth.StringWidth(line)
		if len < width {
			padding := strings.Repeat(" ", width-len)
			line = line + padding
		}
		ret = append(ret, line)
	}
	return
}

func buildBalloon(lines []string, width int) string {
	count := len(lines)
	var ret []string
	ret = append(ret, " "+strings.Repeat("_", width+2))
	if count == 1 {
		ret = append(ret, fmt.Sprintf("< %s >", lines[0]))
	} else {
		ret = append(ret, fmt.Sprintf("/ %s \\", lines[0]))
		for i := 1; i < count-1; i++ {
			ret = append(ret, fmt.Sprintf("| %s |", lines[i]))
		}
		ret = append(ret, fmt.Sprintf("\\ %s /", lines[count-1]))
	}
	ret = append(ret, " "+strings.Repeat("-", width+2))

	return strings.Join(ret, "\n")
}

func printFigure(name string) {
	switch name {
	case "cow":
		fmt.Println(cow)
	case "stegosaurus":
		fmt.Println(stegosaurus)
	case "gopher":
		fmt.Println(gopher)
	case "cat":
		fmt.Println(cat)
	default:
		fmt.Println(cow)
	}
}

func allFigures(figures []string) string {
	count := len(figures)
	if count == 0 {
		return ""
	}
	if count == 1 {
		return figures[0]
	}

	return fmt.Sprintf("%s and %s", strings.Join(figures[:count-1], ", "), figures[count-1])
}
