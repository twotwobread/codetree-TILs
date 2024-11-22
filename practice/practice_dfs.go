package main

import (
	"fmt"
)

type Person struct {
	X     int
	Y     int
	Value int
}

func main() {
	var n int
	fmt.Scanln(&n)

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	people := make([]Person, 3)
	for i := 0; i < 3; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		people[i] = Person{
			X:     x - 1,
			Y:     y - 1,
			Value: 0,
		}
	}

	new := CheckSameTown(n, grid)
	for i := 0; i < 3; i++ {
		x, y := people[i].X, people[i].Y
		people[i].Value = new[x][y]
	}

	if people[0].Value == people[1].Value && people[1].Value == people[2].Value {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", new[i][j])
		}
		fmt.Println()
	}
}

func CheckSameTown(n int, grid [][]int) [][]int {
	check := make([][]int, n)
	for i := 0; i < n; i++ {
		check[i] = make([]int, n)
	}

	cnt := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if check[i][j] == 0 && grid[i][j] == 1 {
				PrintTown(cnt, n, i, j, check, grid)
				cnt += 1
			}
		}
	}
	return check
}

func PrintTown(cnt, n, x, y int, check, grid [][]int) {
	check[x][y] = cnt

	dirX := []int{1, 0, -1, 0}
	dirY := []int{0, -1, 0, 1}
	for i := 0; i < 4; i++ {
		nX := dirX[i] + x
		nY := dirY[i] + y
		if !isOutOfRange(n, nX, nY) && check[nX][nY] == 0 && grid[nX][nY] != 0 {
			PrintTown(cnt, n, nX, nY, check, grid)
		}
	}
}

func isOutOfRange(n, x, y int) bool {
	return x < 0 || x >= n || y < 0 || y >= n
}
