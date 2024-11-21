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

	var n int
	fmt.Fscanln(reader, &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &nums[i])
	}

	result := GetCount(n, nums)
	fmt.Fprintln(printer, result)
}

func GetCount(n int, nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := total / n

	result := 0
	for _, num := range nums {
		if sum < num {
			result += (num - sum)
		}
	}
	return result
}
