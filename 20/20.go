package main

import (
	"fmt"

	"github.com/muradab/advent_of_code/utils"
)

func main() {

	type position struct {
		x, y int
	}
	grid, err := utils.ParseGrid("20/20.txt")

	if err != nil {
		panic(err)
	}

	ROWS := len(grid)
	COLS := len(grid[0])

	start := position{x: 3, y: 1}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if grid[i][j] == 'S' {
				start.x = i
				start.y = j
				break
			}
		}
	}

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
	end := position{x: 0, y: 0}

	for heap.Len() > 0 {

		state, _ := heap.Pop()

		visited[position{state.x, state.y}] = state.cost

		if grid[state.x][state.y] == 'E' {
			end.x = state.x
			end.y = state.y
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

			newCost := state.cost + 1

			heap.Push(State{x: x, y: y, cost: newCost, direction: dir})

		}
	}

	queue := make([]State, 0)
	queue = append(queue, State{start.x, start.y, 0, dirs[0]})
	another_visited := make(map[position]int)
	another_count := -1

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		another_visited[position{current.x, current.y}] = current.cost
		another_count++

		for _, dir := range dirs {

			first := position{x: current.x + dir.x, y: current.y + dir.y}
			// fmt.Println(first, current.x, current.y)
			for _, dir := range dirs {
				second := position{x: first.x + dir.x, y: first.y + dir.y}

				if second.x < 0 || second.x >= ROWS || second.y < 0 || second.y >= COLS {
					continue
				}
				// fmt.Println(second)
				if grid[second.x][second.y] == '.' {
					new_cost := current.cost + 2
					if new_cost < visited[second] {
						fmt.Println(visited[second] - new_cost)
					}
					continue

				}
				for _, dir := range dirs {
					third := position{x: second.x + dir.x, y: second.y + dir.y}
					// fmt.Println(third, second)
					if third.x < 0 || third.x >= ROWS || third.y < 0 || third.y >= COLS {
						continue
					}
					if grid[third.x][third.y] == '#' {
						continue
					}
					new_cost := current.cost + 3
					if new_cost < visited[third] {
						fmt.Println(visited[third] - new_cost)
					}

				}

			}

			x := current.x + dir.x
			y := current.y + dir.y
			if x < 0 || x >= ROWS || y < 0 || y >= COLS {
				continue
			}
			if _, ok := another_visited[position{x, y}]; ok {
				continue
			}
			if grid[x][y] == '#' {
				continue
			}

			newCost := current.cost + 1
			heap.Push(State{x: x, y: y, cost: newCost, direction: dir})
			queue = append(queue, State{x: x, y: y, cost: newCost, direction: dir})
		}
	}

	fmt.Println("another", another_count)
}
