/*
	- 한 라인에 공백으로 띄워진 여러 숫자 배열에 저장하는 연습
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	a := make([]int, n)
	scanner.Scan()
	aStr := strings.Split(scanner.Text(), " ")
	for i := 0; i < n; i++ {
		fmt.Sscanf(aStr[i], "%d", &a[i])
	}

	b := make([]int, n)
	scanner.Scan()
	bStr := strings.Split(scanner.Text(), " ")
	for i := 0; i < n; i++ {
		fmt.Sscanf(bStr[i], "%d", &b[i])
	}

	var builder strings.Builder
	for i := n; i > 0; i-- {
		if i%2 != 0 {

			builder.WriteString(fmt.Sprintf("%d", b[i-1]))
			builder.WriteString(" ")
		}
	}
	builder.WriteString("\n")

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			builder.WriteString(strconv.Itoa(a[i-1]))
			builder.WriteString(" ")
		}
	}
	builder.WriteString("\n")

	result := builder.String()
	fmt.Println(result)
}
