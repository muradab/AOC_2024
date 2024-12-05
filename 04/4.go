package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// read the file
	file, err := os.ReadFile("04/day4.txt")
	if err != nil {
		panic(err)
	}

	// create a grid

	rows := len(strings.Split(string(file), "\n"))
	grid := make([]string, rows)
	for i := 0; i < rows; i++ {
		grid[i] = strings.Split(string(file), "\n")[i]
	}
	part1 := 0
	part2 := 0
	// dirs := []int{-1, 0, 1}
	loop := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] == 'A' {
				loop += 1
				if findMandS(1, 1, i, j, grid) && findMandS(-1, 1, i, j, grid) {
					part2 += 1
				}

			}
			// for k := 0; k < 3; k++ {
			// 	for l := 0; l < 3; l++ {
			// 		loop += 1

			// 		if match(dirs[k], dirs[l], i, j, grid) {
			// 			part1 += 1
			// 		}

			// 	}
			// }
			//fmt.Println(grid[i])
		}
	}
	fmt.Println(grid)
	fmt.Println(part1)
	fmt.Println(loop)
	fmt.Println(part2)

}

func inbound(x, y int, grid []string) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

// func match(left, right, x, y int, grid []string) bool {

// 	word := "XMAS"

// 	for i := 0; i < 4; i++ {
// 		if !inbound(x, y, grid) {
// 			return false
// 		}
// 		if grid[x][y] != word[i] {
// 			return false
// 		}
// 		x += left
// 		y += right

// 	}
// 	return true
// }

func findMandS(left, right int, x, y int, grid []string) bool {

	m := false
	s := false

	if !inbound(x+left, y+right, grid) {
		return false
	}
	if grid[x+left][y+right] == 'M' {
		m = true
	}
	if grid[x+left][y+right] == 'S' {
		s = true
	}

	left *= -1
	right *= -1
	if !inbound(x+left, y+right, grid) {
		return false
	}
	if grid[x+left][y+right] == 'M' {
		m = true
	}
	if grid[x+left][y+right] == 'S' {
		s = true
	}

	return m && s

}
