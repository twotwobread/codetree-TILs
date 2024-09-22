package main

import (
	"bufio"
	"fmt"
	"os"
)

const OneWeekDays = 7

var (
	br   *bufio.Reader
	bw   *bufio.Writer
	DAYS = []int{
		0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
	}
	WEEK_DAYS = map[string]int{
		"Mon": 0,
		"Tue": 1,
		"Wed": 2,
		"Thu": 3,
		"Fri": 4,
		"Sat": 5,
		"Sun": 6,
	}
)

func init() {
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
}

func main() {
	defer bw.Flush()
	var m1, d1, m2, d2 int
	fmt.Fscanf(br, "%d %d %d %d", &m1, &d1, &m2, &d2)
	br.ReadString('\n')
	var dow string
	fmt.Fscanln(br, &dow)

	dowNum := WEEK_DAYS[dow]
	days1 := GetDaysFrom0101(m1, d1)
	days2 := GetDaysFrom0101(m2, d2)

	diff := days2 - days1
	cnt := int(diff / OneWeekDays)
	if diff%OneWeekDays >= dowNum {
		cnt += 1
	}
	fmt.Fprintln(bw, cnt)
}

func GetDaysFrom0101(m int, d int) int {
	total := 0
	for i := 1; i <= m; i++ {
		if i == m {
			total += d
		} else {
			total += DAYS[i]
		}
	}
	return total
}
