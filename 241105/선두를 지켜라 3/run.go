package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MaxMoveCnt      = 1000
	MaxMoveDuration = 1000
)

var (
	printer *bufio.Writer
	reader  *bufio.Reader
)

func init() {
	printer = bufio.NewWriter(os.Stdout)
	reader = bufio.NewReader(os.Stdin)
}

type Move struct {
	Velocity int
	Duration int
}

func main() {
	defer printer.Flush()

	var moveHoursA, moveHoursB int
	fmt.Fscanln(reader, &moveHoursA, &moveHoursB)

	movesA := readMoves(moveHoursA)
	movesB := readMoves(moveHoursB)

	footprintsA, timeA := writeFootprints(movesA)
	footprintsB, timeB := writeFootprints(movesB)
	maxTime := max(timeA, timeB)

	result := getChangeHeadCnt(maxTime, footprintsA, footprintsB)
	fmt.Fprintln(printer, result)
}

func readMoves(cnt int) []Move {
	moves := make([]Move, cnt)
	for i := 0; i < cnt; i++ {
		var v, t int
		fmt.Fscanln(reader, &v, &t)
		moves[i] = Move{
			Velocity: v,
			Duration: t,
		}
	}
	return moves
}

func writeFootprints(ms []Move) ([]int, int) {
	time, pos := 0, 0
	footprints := make([]int, MaxMoveCnt*MaxMoveDuration)
	for _, m := range ms {
		for i := 0; i < m.Duration; i++ {
			pos += m.Velocity
			footprints[time] = pos
			time += 1
		}
	}

	return footprints, time
}

func getChangeHeadCnt(t int, fpa, fpb []int) int {
	cnt := 0
	head := -1
	for i := 0; i < t; i++ {
		if fpa[i] > fpb[i] {
			if head != 0 {
				cnt += 1
			}
			head = 0
		} else if fpa[i] < fpb[i] {
			if head != 1 {
				cnt += 1
			}
			head = 1
		} else {
			if head != 2 {
				cnt += 1
			}
			head = 2
		}
	}
	return cnt
}
