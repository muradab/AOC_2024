package main

import (
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	x, y int
}

func main() {

	file, err := os.ReadFile("12/day12.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	dirs := []Coordinate{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	visited := make(map[Coordinate]bool)

	bfs := func(start Coordinate) (int, int) {

		visited[start] = true
		x := start.x
		y := start.y
		current_element := lines[x][y]
		queue := make([]Coordinate, 0)
		queue = append(queue, start)
		area := 0
		perimeter := 0
		for len(queue) > 0 {
			fmt.Println(queue)

			area++
			last := queue[len(queue)-1]
			queue = queue[:len(queue)-1]

			outside := func(x, y int) bool {
				return !inbound(x, y, lines) || lines[x][y] != current_element
			}

			for i := 0; i < 4; i++ {
				x := last.x + dirs[i].x
				y := last.y + dirs[i].y
				nx := last.x + dirs[(i+1)%len(dirs)].x
				ny := last.y + dirs[(i+1)%len(dirs)].y
				dx := last.x + dirs[i].x + dirs[(i+1)%len(dirs)].x
				dy := last.y + dirs[i].y + dirs[(i+1)%len(dirs)].y

				if outside(x, y) && outside(nx, ny) {
					perimeter++
				}

				if !outside(x, y) && !outside(nx, ny) {
					if lines[dx][dy] != current_element {
						perimeter++
					}
				}

				if outside(x, y) {
					continue
				}

				if _, ok := visited[Coordinate{x, y}]; ok {
					continue
				}

				queue = append(queue, Coordinate{x, y})
				visited[Coordinate{x, y}] = true
			}
		}
		return area, perimeter
	}

	res := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {

			current := Coordinate{i, j}
			if _, ok := visited[current]; ok {
				continue
			}
			x, y := bfs(current)
			fmt.Println(x, y)

			res += x * y

		}
		// fmt.Println(visited)
	}

	fmt.Println(res)

}
func inbound(x, y int, lines []string) bool {
	return x >= 0 && x < len(lines) && y >= 0 && y < len(lines[0])
}
