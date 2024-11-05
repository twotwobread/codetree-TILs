package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MaxMoveCount    = 1000
	MaxMoveDuration = 1000
)

var (
	reader *bufio.Reader
	writer *bufio.Writer
)

func init() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

type Move struct {
	Direction string
	Duration  int // Time을 Duration으로 변경
}

func main() {
	defer writer.Flush()

	var numMovesA, numMovesB int
	fmt.Fscanln(reader, &numMovesA, &numMovesB)

	movesA := readMoves(numMovesA)
	movesB := readMoves(numMovesB)

	// Write the footprints for each player
	footprintsA, timeA := writeFootprints(movesA)
	footprintsB, timeB := writeFootprints(movesB)

	// Find the first time the footprints meet
	maxTime := max(timeA, timeB)
	firstMeetTime := getFirstMeetTime(maxTime, footprintsA, footprintsB)
	fmt.Fprintln(writer, firstMeetTime)
}

func readMoves(numMoves int) []Move { // move input function로 빼내
	moves := make([]Move, numMoves)
	for i := 0; i < numMoves; i++ {
		var direction string
		var duration int
		fmt.Fscanln(reader, &direction, &duration)
		moves[i] = Move{
			Direction: direction,
			Duration:  duration,
		}
	}
	return moves
}

func writeFootprints(moves []Move) ([]int, int) {
	// footprints := make([]int, 0, len(moves)*10)
	footprints := make([]int, MaxMoveCount*MaxMoveDuration)
	time := 0
	currentPosition := 0
	for _, move := range moves {
		direction := 1
		if move.Direction == "L" {
			direction = -1
		}
		for i := 0; i < move.Duration; i++ {
			currentPosition += direction
			// footprints = append(footprints, currentPosition) // append 시간복잡도가 O(N)임.
			// Go 배열은 고정된 크기를 가진 array니까 새로운 배열을 만들고 데이터를 복사하는 로직에서 O(N)이 걸림.
			footprints[time] = currentPosition
			time += 1
		}
	}
	// 그럼 해당 함수의 시간복잡도는 O(N^3)임. -> 여기서 시간 효율성에 문제가 생김. -> append 쓰면 안됨.
	return footprints, time
}

func getFirstMeetTime(time int, footprintsA, footprintsB []int) int {
	for i := 0; i < time; i++ {
		indexA, indexB := i, i
		if i >= len(footprintsA) {
			indexA = len(footprintsA) - 1
		}
		if i >= len(footprintsB) {
			indexB = len(footprintsB) - 1
		}
		if footprintsA[indexA] == footprintsB[indexB] {
			return i + 1
		}
	}
	return -1
}

func max(a, b int) int { // max 함수로 빼내기
	if a > b {
		return a
	}
	return b
}
