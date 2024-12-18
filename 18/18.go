package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("18/18.txt")

	if err != nil {
		panic(err)
	}
	ROWS := 71
	COLS := 71

	grid := make([][]byte, ROWS)
	for i := 0; i < ROWS; i++ {
		grid[i] = make([]byte, COLS)
	}

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			grid[i][j] = '.'
		}
	}

	for k, val := range strings.Split(string(file), "\n") {
		nums := strings.Split(val, ",")
		r, c := convertToNum(nums[0]), convertToNum(nums[1])
		grid[c][r] = '#'

		// for i := 0; i < ROWS; i++ {
		// 	fmt.Println(string(grid[i]))
		// 	//fmt.Println(grid[i])
		// }
		good := false

		dirs := []int{-1, 0, 1, 0, -1}
		type state struct {
			x, y int
			cost int
		}

		type coordinate struct {
			x, y int
		}
		var queue []state
		queue = append(queue, state{0, 0, 0})
		visited := make(map[coordinate]bool)
		visited[coordinate{0, 0}] = true

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			if current.x == ROWS-1 && current.y == COLS-1 {
				fmt.Println(current.cost)
				good = true
				break
			}

			for i := 0; i < 4; i++ {
				x := current.x + dirs[i]
				y := current.y + dirs[i+1]
				if x < 0 || x >= ROWS || y < 0 || y >= COLS {
					continue
				}
				if grid[x][y] == '#' {
					continue
				}
				if visited[coordinate{x, y}] {
					continue
				}
				cost := current.cost + 1
				visited[coordinate{x, y}] = true
				queue = append(queue, state{x, y, cost})
			}
		}

		if !good {
			fmt.Println(r, c, k)
			break
		}

	}

}

func convertToNum(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
