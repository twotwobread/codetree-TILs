package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	xMax int = 100
	yMax int = 100
)

var (
	br *bufio.Reader
	bw *bufio.Writer
)

func init() {
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
}

func main() {
	defer bw.Flush()

	// input values
	var n int
	fmt.Fscanln(br, &n)

	cmds := make([][]int, n)
	for i := range cmds {
		cmds[i] = make([]int, 4)
		var x1, y1, x2, y2 int
		fmt.Fscanln(br, &x1, &y1, &x2, &y2)
		cmds[i][0], cmds[i][1], cmds[i][2], cmds[i][3] = x1+xMax, y1+yMax, x2+xMax, y2+yMax
	}

	// initialize grids
	grid := make([][]bool, xMax*2+1)
	for i := range xMax*2 + 1 {
		grid[i] = make([]bool, yMax*2+1)
	}

	for _, cmd := range cmds {
		x1, y1 := cmd[0], cmd[1]
		x2, y2 := cmd[2], cmd[3]

		for x := x1; x < x2; x++ {
			for y := y1; y < y2; y++ {
				grid[x][y] = true
			}
		}
	}

	extent := 0
	for _, g := range grid {
		for _, b := range g {
			if b == true {
				extent += 1
			}
		}
	}

	fmt.Fprintln(bw, extent)
}
