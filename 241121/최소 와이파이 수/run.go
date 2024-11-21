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

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	people := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &people[i])
	}

	result := GetMinWifiCnt(n, m, people)
	fmt.Fprintln(printer, result)
}

func GetMinWifiCnt(n, m int, people []int) int {
	cnt := 0
	// for i, p := range people {
	for i := 0; i < n; i++ {
		if people[i] == 1 {
			cnt += 1
			i += 2 * m
		}
	}
	return cnt
}

// func CheckUsable(now, max, r int, check []bool) {
// 	for i := now; i <= now+max*2; i++ {
// 		if IsOutOfRange(r, i) {
// 			break
// 		}

// 		check[i] = true
// 	}
// }

// func IsOutOfRange(r, now int) bool {
// 	return now < 0 || now >= r
// }
