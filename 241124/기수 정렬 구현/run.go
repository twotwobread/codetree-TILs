package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	maxNum := 0
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		if maxNum < nums[i] {
			maxNum = nums[i]
		}
	}

	sorted := SortImprovedRadix(nums, maxNum)
	for _, num := range sorted {
		fmt.Printf("%d ", num)
	}
	// maxRadix := 0
	// nums := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&nums[i])
	// 	radix := GetRadixSize(nums[i])
	// 	if maxRadix < radix {
	// 		maxRadix = radix
	// 	}
	// }

	// sorted := SortByRadix(maxRadix, nums)
	// for i := 0; i < 10; i++ {
	// 	for _, num := range sorted[i] {
	// 		fmt.Printf("%d ", num)
	// 	}
	// }
}

func SortImprovedRadix(nums []int, maxNum int) []int {
	exp := 1
	counts := make([]int, 10) // 배열 재사용을 통해서 메모리 할당 비용을 최소화
	result := make([]int, len(nums))

	for maxNum/exp > 0 {
		for i := 0; i < 10; i++ {
			counts[i] = 0
		}

		// counting
		for _, num := range nums {
			rNum := (num / exp) % 10
			counts[rNum]++
		}

		// prefix sum
		for i := 1; i < 10; i++ {
			counts[i] += counts[i-1]
		}

		// reverse iterate for stability sorting
		for i := len(nums) - 1; i >= 0; i-- {
			rNum := (nums[i] / exp) % 10
			counts[rNum]--
			result[counts[rNum]] = nums[i]
		}

		exp *= 10
		nums, result = result, nums // 기존에 copy를 이용했는데 그 비용 최소화
	}

	return nums
}

// func GetRadixSize(num int) int {
// 	cnt := 0
// 	for num > 0 {
// 		num /= 10
// 		cnt += 1
// 	}
// 	return cnt
// }

// 전통적인 radix
// 1. append를 계속 처리해야해서 O(k*N*N) 시간복잡도가 나옴. -> 생각보다 되게 느림.
// 2. 메모리적으로 효율도 안좋음 N*10*2 정도 사용하게 됨.
// func SortByRadix(radix int, nums []int) [][]int {
// 	devideNum := 10
// 	before := make([][]int, 10)
// 	sorted := make([][]int, 10)
// 	for _, num := range nums {
// 		rValue := num % devideNum
// 		before[rValue] = append(before[rValue], num)
// 	}

// 	for i := 2; i <= radix; i++ {
// 		for _, values := range before {
// 			for _, v := range values {
// 				nowDevide := int(math.Pow(float64(devideNum), float64(i)))
// 				rValue := v % nowDevide
// 				rValue = rValue / (nowDevide / devideNum)
// 				sorted[rValue] = append(sorted[rValue], v)
// 			}
// 		}

// 		for i := 0; i < 10; i++ {
// 			before[i] = sorted[i]
// 			sorted[i] = []int{}
// 		}
// 	}

// 	return before
// }
