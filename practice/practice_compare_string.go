package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var result []string
	bN := []string{"0", "1"}

	for _, b := range bN {
		result = GetBinNums(result, n, b)
	}

	for _, bin := range result {
		fmt.Println(bin)
	}
}

func GetBinNums(result []string, l int, b string) []string {
	if len(b) >= l {
		bS := ""
		cnt := 1
		for _, c := range b {
			if bS == string(c) {
				cnt += 1
				if cnt >= 3 {
					return result
				}
			} else {
				bS = string(c)
				cnt = 1
			}
		}
		result = append(result, b)
		return result
	}

	result = GetBinNums(result, l, b+"0")
	result = GetBinNums(result, l, b+"1")
	return result
}
