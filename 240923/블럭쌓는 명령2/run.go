package main

import (
	"bufio"
	"fmt"
	"os"
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

	var n, k int
	fmt.Fscanln(br, &n, &k)

	starts := make([]int, k)
	ends := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscanln(br, &starts[i], &ends[i])
	}

	blocks := make([]int, n+1)
	StackBlock(starts, ends, blocks)
	result := GetMaxIndex(blocks)

	fmt.Fprintln(bw, result)
}

func StackBlock(starts []int, ends []int, blocks []int) {
	for i := 0; i < len(starts); i++ {
		for j := starts[i]; j <= ends[i]; j++ {
			blocks[j] += 1
		}
	}
}

func GetMaxIndex(blocks []int) int {
	maxValue := 0
	for _, b := range blocks {
		if maxValue < b {
			maxValue = b
		}
	}
	return maxValue
}
