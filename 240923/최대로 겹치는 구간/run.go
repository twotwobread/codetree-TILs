package main

import "fmt"

const (
	MaxRange = 100
	MinRange = -100
	Offset   = 100
)

type Line struct {
	start int
	end   int
}

func main() {
	var n int
	fmt.Scanln(&n)

	lines := make([]Line, n)
	for i := range n {
		var x1, x2 int
		fmt.Scanf("%d %d", &x1, &x2)
		lines[i] = Line{start: x1, end: x2}
	}

	printedLines := PrintLines(lines)
	result := FindMaxStackedLineCnt(printedLines)
	fmt.Println(result)
}

func PrintLines(lines []Line) []int {
	printed := make([]int, MaxRange+Offset)
	for _, l := range lines {
		for s := l.start; s < l.end; s++ {
			printed[s+Offset] += 1
		}
	}
	return printed
}

func FindMaxStackedLineCnt(lines []int) int {
	maxValue := -1
	for _, n := range lines {
		if maxValue < n {
			maxValue = n
		}
	}
	return maxValue
}
