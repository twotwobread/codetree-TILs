package main

import (
	"bufio"
	"fmt"
	"os"
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

	var n, s int
	fmt.Fscanln(reader, &n, &s)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	diff := getClosestDiff(s, nums)
	fmt.Fprintln(printer, diff)
}

func getClosestDiff(goal int, nums []int) int {
	size := len(nums)
	sum := 0
	for i := 0; i < size; i++ {
		sum += nums[i]
	}

	diff := goal
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			nSum := sum - nums[i] - nums[j]
			dTmp := getDiff(goal, nSum)
			if diff > dTmp {
				diff = dTmp
			}
		}
	}
	return diff
}

func getDiff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
