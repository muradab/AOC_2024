package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction struct {
	x  int
	y  int
	rx int
	ry int
}

type Position struct {
	x int
	y int
}

func main() {

	file, err := os.ReadFile("15/15.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n\n")

	// starting := Position{0, 0}
	movements := strings.ReplaceAll(lines[1], "\n", "")

	Part2(strings.Split(lines[0], "\n"), movements)

	// grid := make([][]byte, 0)
	// for j, line := range strings.Split(lines[0], "\n") {
	// 	grid = append(grid, []byte(line))
	// 	for i, c := range line {
	// 		if c == '@' {
	// 			starting.x = j
	// 			starting.y = i

	// 		}

	// 	}
	// }

	// fmt.Println(starting)

	// movements := strings.ReplaceAll(lines[1], "\n", "")

	// dirs := map[byte]Direction{
	// 	'^': {-1, 0, 1, 0},
	// 	'v': {1, 0, -1, 0},
	// 	'>': {0, 1, 0, -1},
	// 	'<': {0, -1, 0, 1},
	// }

	// for i := 0; i < len(movements); i++ {
	// 	// for i := 0; i < len(grid); i++ {
	// 	// 	for j := 0; j < len(grid[i]); j++ {
	// 	// 		fmt.Print(string(grid[i][j]))
	// 	// 	}
	// 	// 	fmt.Println()
	// 	// }
	// 	dir := dirs[movements[i]]
	// 	fmt.Println(string(movements[i]), starting)
	// 	temp := Position{starting.x + dir.x, starting.y + dir.y}

	// 	for grid[temp.x][temp.y] == 'O' {
	// 		temp.x += dir.x
	// 		temp.y += dir.y
	// 		fmt.Print(temp.x, temp.y)

	// 	}
	// 	fmt.Println()
	// 	if grid[temp.x][temp.y] == '#' {
	// 		continue
	// 	}
	// 	if grid[temp.x][temp.y] == '.' {
	// 		// shift items from starting to temp
	// 		for grid[temp.x][temp.y] != '@' {
	// 			grid[temp.x][temp.y] = 'O'
	// 			temp.x += dir.rx
	// 			temp.y += dir.ry

	// 		}
	// 		grid[temp.x+dir.x][temp.y+dir.y] = '@'
	// 		grid[temp.x][temp.y] = '.'

	// 		starting = Position{temp.x + dir.x, temp.y + dir.y}
	// 	}
	// 	// for i := 0; i < len(grid); i++ {
	// 	// 	for j := 0; j < len(grid[i]); j++ {
	// 	// 		fmt.Print(string(grid[i][j]))
	// 	// 	}
	// 	// 	fmt.Println()
	// 	// }

	// }

	// cost := 0

	// for i := 0; i < len(grid); i++ {
	// 	for j := 0; j < len(grid[i]); j++ {

	// 		if grid[i][j] == 'O' {

	// 			cost += i*100 + j
	// 		}
	// 	}

	// }
	// fmt.Println(cost)

}

func Part2(input []string, movements string) {
	dirs := map[byte]Direction{
		'^': {-1, 0, 1, 0},
		'v': {1, 0, -1, 0},
		'>': {0, 1, 0, -1},
		'<': {0, -1, 0, 1},
	}
	starting := Position{0, 0}

	grid := make([][]byte, len(input))
	fmt.Println(grid)

	for i := 0; i < len(input); i++ {

		for j := 0; j < len(input[i]); j++ {

			if input[i][j] == 'O' {
				grid[i] = append(grid[i], '[')
				grid[i] = append(grid[i], ']')
			}
			if input[i][j] == '#' {
				grid[i] = append(grid[i], '#')
				grid[i] = append(grid[i], '#')
			}
			if input[i][j] == '.' {
				grid[i] = append(grid[i], '.')
				grid[i] = append(grid[i], '.')
			}
			if input[i][j] == '@' {
				starting.x = i
				starting.y = j * 2
				grid[i] = append(grid[i], '@')
				grid[i] = append(grid[i], '.')
			}
		}
	}

	for i := 0; i < len(movements); i++ {
		// for i := 0; i < len(grid); i++ {
		// 	for j := 0; j < len(grid[i]); j++ {
		// 		fmt.Print(string(grid[i][j]))
		// 	}
		// 	fmt.Println()
		// }
		dir := dirs[movements[i]]
		fmt.Println(string(movements[i]), starting)
		temp := Position{starting.x + dir.x, starting.y + dir.y}

		for grid[temp.x][temp.y] == 'O' {
			temp.x += dir.x
			temp.y += dir.y
			fmt.Print(temp.x, temp.y)

		}
		fmt.Println()
		if grid[temp.x][temp.y] == '#' {
			continue
		}
		if grid[temp.x][temp.y] == '.' {
			// shift items from starting to temp
			for grid[temp.x][temp.y] != '@' {
				grid[temp.x][temp.y] = 'O'
				temp.x += dir.rx
				temp.y += dir.ry

			}
			grid[temp.x+dir.x][temp.y+dir.y] = '@'
			grid[temp.x][temp.y] = '.'

			starting = Position{temp.x + dir.x, temp.y + dir.y}
		}
		// for i := 0; i < len(grid); i++ {
		// 	for j := 0; j < len(grid[i]); j++ {
		// 		fmt.Print(string(grid[i][j]))
		// 	}
		// 	fmt.Println()
		// }

	}

	fmt.Println(starting)

}
