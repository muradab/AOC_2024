package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y int
}

// I complicated it too much 🥲🥲🥲🥲🥲🥲🥲🥲🥲🥲🥲🥲🥲🥲🥲

func main() {

	file, err := os.ReadFile("10/day10.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	grid := make([][]int, len(lines))
	location := make(map[int][]Coordinate)

	for i, line := range lines {
		for j, l := range line {
			num, _ := strconv.Atoi(string(l))
			grid[i] = append(grid[i], num)
			if _, ok := location[num]; !ok {
				location[num] = make([]Coordinate, 0)
			}
			location[num] = append(location[num], Coordinate{i, j})
		}

	}

	res := make(map[Coordinate][]Coordinate)

	dirs := []int{-1, 0, 1, 0, -1}

	for _, v := range location[9] {
		res[v] = make([]Coordinate, 0)
		res[v] = append(res[v], v)

	}

	fmt.Println(res)

	for i := 8; i > -1; i-- {
		// fmt.Println(location[i])

		for _, val := range location[i] {

			for k := 0; k < 4; k++ {
				x := val.x + dirs[k]
				y := val.y + dirs[k+1]
				// fmt.Println(x, y)

				if inbound(x, y, len(lines), len(lines[0])) {
					// fmt.Println(grid[x][y])

					if grid[x][y] == i+1 {
						// senior citizen
						// fmt.Println("senior citizen", grid[x][y], i, x, y)
						senior := Coordinate{x, y}

						res[val] = append(res[val], res[senior]...)

					}

				}
			}
		}

	}
	fmt.Println(res)
	fmt.Println(part2(grid, location))

	ans := 0
	fmt.Println(len(location[0]))

	for _, v := range location[0] {

		unique := make(map[Coordinate]bool)

		for _, k := range res[v] {
			if !unique[k] {

				ans += 1
			}
			unique[k] = true
		}
		fmt.Println(unique, v)

	}

	fmt.Println(ans)

	fmt.Println(res)

}

// places := make([][]Coordinate, 0)

func inbound(x, y int, row, col int) bool {
	return x >= 0 && x < row && y >= 0 && y < col
}

func part2(grid [][]int, locations map[int][]Coordinate) int {

	dirs := []int{-1, 0, 1, 0, -1}
	res := 0
	for _, v := range locations[0] {
		res += dp(v, 0, grid, dirs)

	}

	return res

}

func dp(start Coordinate, height int, grid [][]int, dirs []int) int {

	if height == 9 {
		return 1
	}

	ans := 0

	for i := 0; i < 4; i++ {
		x := start.x + dirs[i]
		y := start.y + dirs[i+1]
		if !inbound(x, y, len(grid), len(grid[0])) {
			continue
		}
		if grid[x][y] == height+1 {
			ans += dp(Coordinate{x, y}, height+1, grid, dirs)
		}
	}
	return ans
}
