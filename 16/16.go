package main

import (
	"fmt"
	"strconv"

	"github.com/muradab/advent_of_code/utils"
)

func main() {

	type position struct {
		x, y int
	}
	grid, err := utils.ParseGrid("16/16.txt")

	if err != nil {
		panic(err)
	}

	ROWS := len(grid)
	COLS := len(grid[0])

	start := position{x: ROWS - 2, y: 1}

	dirs := []position{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	type State struct {
		x, y      int
		cost      int
		direction position
	}

	heap := utils.NewHeap(func(a, b State) bool {
		return a.cost < b.cost
	})

	heap.Push(State{x: start.x, y: start.y, cost: 0, direction: dirs[0]})
	visited := make(map[position]int)
	prev := make(map[position][]position)

	for heap.Len() > 0 {

		state, _ := heap.Pop()

		visited[position{state.x, state.y}] = state.cost
		fmt.Println(state.x, state.y)

		if state.x == 1 && state.y == COLS-2 {
			fmt.Println(state.cost)
			break

			// break
		}

		for _, dir := range dirs {
			x := state.x + dir.x
			y := state.y + dir.y
			if x < 0 || x >= ROWS || y < 0 || y >= COLS {
				continue
			}
			if _, ok := visited[position{x, y}]; ok {
				continue
			}
			if grid[x][y] == '#' {
				continue
			}
			if x == 8 && y == 5 {
				fmt.Println("print here", state.x, state.y)
			}
			newPosition := position{x: x, y: y}
			slice := prev[newPosition]
			slice = append(slice, position{x: state.x, y: state.y})
			prev[newPosition] = slice
			newCost := state.cost + 1
			if state.direction.x != dir.x || state.direction.y != dir.y {
				newCost += 1000
			}
			heap.Push(State{x: x, y: y, cost: newCost, direction: dir})

		}
	}

	queue := make([]position, 0)
	queue = append(queue, position{x: 1, y: COLS - 2})
	seen := make(map[position]bool)
	fmt.Println(prev[position{x: 10, y: 1}], "megenteta")
	fmt.Println(prev[position{x: 11, y: 2}], "megenteta")
	count := 0

	for len(queue) > 0 {
		fmt.Println(queue)

		state := queue[0]
		queue = queue[1:]
		count++

		for _, pair := range prev[state] {

			queue = append(queue, pair)
			seen[pair] = true

		}
	}
	ir := make([][]string, ROWS)

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if v, ok := visited[position{x: i, y: j}]; ok {
				ir[i] = append(ir[i], strconv.Itoa(v))
			} else {
				ir[i] = append(ir[i], string(grid[i][j]))

			}
		}
	}
	fmt.Println(count)
	for _, row := range ir {
		fmt.Println(fmt.Sprintf(fmt.Sprintf("%%-%ds", 4), row))
	}

}
