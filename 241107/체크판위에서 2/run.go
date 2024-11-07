package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	White = "W"
	Black = "B"
)

var (
	printer *bufio.Writer
	reader  *bufio.Reader
)

func init() {
	printer = bufio.NewWriter(os.Stdout)
	reader = bufio.NewReader(os.Stdin)
}

type Point struct {
	X     int
	Y     int
	Value string
}

func main() {
	defer printer.Flush()

	var r, c int
	fmt.Fscanln(reader, &r, &c)

	arr := make([][]string, r)
	for i := 0; i < r; i++ {
		arr[i] = make([]string, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			var x string
			fmt.Fscan(reader, &x)
			arr[i][j] = x
		}
	}

	result := []Point{
		{X: 0, Y: 0, Value: arr[0][0]},
	}
	noc := getNumberOfCases(0, 0, 0, arr, result)
	fmt.Fprintln(printer, noc)
}

func getNumberOfCases(n, r, c int, arr [][]string, result []Point) int {
	maxR := len(arr)
	maxC := len(arr[0])
	if r > 0 && c > 0 && arr[maxR-1][maxC-1] != arr[r][c] && n == 2 {
		return 1
	}

	cnt := 0
	for i := r + 1; i < maxR-1; i++ {
		for j := c + 1; j < maxC-1; j++ {
			if arr[r][c] != arr[i][j] {
				result = append(result, Point{X: i, Y: j, Value: arr[i][j]})
				cnt += getNumberOfCases(n+1, i, j, arr, result)
				result = result[:len(result)-1]
			}
		}
	}
	return cnt
}
