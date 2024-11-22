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
		fmt.Fscan(reader, &nums[i])
	}

	oddCnt, evenCnt := GetCnt(n, nums)
	result := GetMaxPairCnt(oddCnt, evenCnt)
	fmt.Fprintln(printer, result)
}

func GetCnt(n int, nums []int) (int, int) {
	oI, eI := 0, 0
	for i := 0; i < n; i++ {
		if nums[i]%2 == 0 {
			eI++
		} else {
			oI++
		}
	}

	return oI, eI
}

func GetMaxPairCnt(odd, even int) int {
	loop := 0
	rOdd, rEven := 0, 0
	for {
		if odd == 0 && even == 0 {
			break
		}

		if loop%2 == 0 { // 짝수 차례
			if even > 0 {
				even -= 1
				rEven += 1
			} else if odd >= 2 {
				odd -= 2
				rEven += 1
			} else { // 홀수 1개 남음. -> 2번전의 짝과 1번 전의 홀 그리고 지금 홀 더해서 짝 만들기 -> 그룹 개수가 하나 적어짐.
				return rOdd + rEven - 1
			}
		} else { // 홀수 차례
			if odd > 0 {
				odd -= 1
				rOdd += 1
			} else { // 짝수만 남음 -> 이전 홀과 지금 짝을 합침 -> 현재 그룹 개수가 답
				return rOdd + rEven
			}
		}
		loop += 1
	}
	return rOdd + rEven
}
