package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
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

	var a, b int
	fmt.Fscanf(br, "%d %d", &a, &b)
	br.ReadString('\n')
	var n int
	fmt.Fscanf(br, "%d", &n)

	bNum := ConvertNum(n, a, b)
	fmt.Fprintln(bw, bNum)
}

func ConvertNum(num int, from int, to int) string {
	if num == 0 {
		return "0"
	}

	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	decNum := 0
	cnt := 0
	for num > 0 {
		remainder := num % 10
		decNum += int(math.Pow(float64(from), float64(cnt))) * remainder
		cnt += 1
		num /= 10
	}

	for decNum > 0 {
		remainder := decNum % to
		writer.WriteByte('0' + byte(remainder))
		decNum /= to
	}
	writer.Flush()

	bBytes := buf.Bytes()
	for i, j := 0, len(bBytes)-1; i < j; i, j = i+1, j-1 {
		bBytes[i], bBytes[j] = bBytes[j], bBytes[i]
	}
	return string(bBytes)
}
