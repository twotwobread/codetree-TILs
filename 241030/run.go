package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader  *bufio.Reader
	scanner *bufio.Scanner
	printer *bufio.Writer
)

func init() {
	reader = bufio.NewReader(os.Stdin)
	scanner = bufio.NewScanner(reader)
	printer = bufio.NewWriter(os.Stdout)
}

func main() {
	defer printer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscanln(reader, &x)
		nums[i] = x
	}

	cValue := 1
	mValue := 1
	pValue := nums[0]
	for i := 1; i < n; i++ {
		if pValue == nums[i] {
			cValue++
		} else {
			if mValue < cValue {
				mValue = cValue
			}
			cValue = 1
		}
		pValue = nums[i]
	}

	if mValue < cValue {
		mValue = cValue
	}

	fmt.Fprintf(printer, "%d", mValue)
}
