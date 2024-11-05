package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	writer  *bufio.Reader
	printer *bufio.Writer
)

func init() {
	writer = bufio.NewReader(os.Stdin)
	printer = bufio.NewWriter(os.Stdout)
}

type Move struct {
	Direction string
	Time      int
}

func main() {
	defer printer.Flush()

	var n, m int
	fmt.Fscanln(writer, &n, &m)

	mA := make([]Move, n)
	for i := 0; i < n; i++ {
		var d string
		var t int
		fmt.Fscanln(writer, &d, &t)
		mA[i] = Move{
			Direction: d,
			Time:      t,
		}
	}

	mB := make([]Move, m)
	for i := 0; i < m; i++ {
		var d string
		var t int
		fmt.Fscanln(writer, &d, &t)
		mB[i] = Move{
			Direction: d,
			Time:      t,
		}
	}

	fpA := WriteFootprints(mA)
	fpB := WriteFootprints(mB)

	t := GetFirstMetTime(fpA, fpB)
	fmt.Fprintln(printer, t)
}

func WriteFootprints(ms []Move) []int {
	// time := make([]int, len(ms))
	var time []int
	curPos := 0
	for _, m := range ms {
		d := 1
		if m.Direction == "L" {
			d *= -1
		}

		for i := 0; i < m.Time; i++ {
			curPos += d
			time = append(time, curPos)
		}
	}
	return time
}

func GetFirstMetTime(aF []int, bF []int) int {
	var loop int
	if len(aF) > len(bF) {
		loop = len(aF)
	} else {
		loop = len(bF)
	}

	for i := 0; i < loop; i++ {
		x, y := i, i
		if i >= len(aF) {
			x = len(aF) - 1
		}
		if i >= len(bF) {
			y = len(bF) - 1
		}

		if aF[x] == bF[y] {
			return i + 1
		}
	}
	return -1
}
