package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	printer *bufio.Writer
	reader  *bufio.Reader
)

func init() {
	printer = bufio.NewWriter(os.Stdout)
	reader = bufio.NewReader(os.Stdin)
}

func main() {
	defer printer.Flush()

	var size, cmdCnt int
	fmt.Fscanln(reader, &size, &cmdCnt)

	var dirStr string
	fmt.Fscanln(reader, &dirStr)
	dirs := strings.Split(dirStr, "")

	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Fscan(reader, &grid[i][j])
		}
	}

	total := GetSumMoveNums(grid, dirs)
	fmt.Fprintln(printer, total)
}

func GetSumMoveNums(grid [][]int, dirs []string) int {
	mid := int(len(grid) / 2)
	dir, cR, cC := 0, mid, mid
	total := 0
	for _, cmd := range dirs {
		nDir, nR, nC := rotate(cmd, dir, cR, cC)
		if !isOutOfRange(len(grid), nR, nC) && isMove(nR, nC, cR, cC) {
			total += grid[cR][cC]
			cR, cC = nR, nC
		}
		dir = nDir
	}
	return total + grid[cR][cC]
}

func isMove(nR, nC, cR, cC int) bool {
	return nR != cR || nC != cC
}

func rotate(cmd string, dir, r, c int) (int, int, int) {
	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}

	if cmd == "L" {
		return (dir + 3) % 4, r, c
	} else if cmd == "R" {
		return (dir + 1) % 4, r, c
	} else {
		return dir, r + dy[dir], c + dx[dir]
	}
}

func isOutOfRange(size, r, c int) bool {
	return r < 0 || r >= size || c < 0 || c >= size
}
