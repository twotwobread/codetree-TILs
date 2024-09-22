package main

import (
	"bufio"
	"bytes"
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

	var n int
	fmt.Fscanf(br, "%d", &n)

	bNum := GetBinaryNum(n)
	fmt.Fprintln(bw, bNum)
}

func GetBinaryNum(num int) string {
	if num == 0 {
		return "0"
	}

	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	for num > 0 {
		remainder := num % 2
		writer.WriteByte('0' + byte(remainder))
		num /= 2
	}
	writer.Flush()

	bBytes := buf.Bytes()
	for i, j := 0, len(bBytes)-1; i < j; i, j = i+1, j-1 {
		bBytes[i], bBytes[j] = bBytes[j], bBytes[i]
	}
	return string(bBytes)
}
