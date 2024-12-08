package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x, y int
}

func main() {
	file, err := os.ReadFile("08/day8.txt")
	if err != nil {
		panic(err)
	}
	grid := strings.Split(string(file), "\n")

	fmt.Println(part1(grid))
	fmt.Println(part2(grid))

}

func part1(grid []string) int {
	row := len(grid)
	col := len(grid[0])
	count := 0
	visited := make(map[int]bool, row)
	pairs := make(map[string][]Position)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] != '.' {
				pairs[string(grid[i][j])] = append(pairs[string(grid[i][j])], Position{i, j})
			}
		}
	}

	for _, v := range pairs {
		for i := 0; i < len(v); i++ {

			for j := i + 1; j < len(v); j++ {

				row_diff := v[j].x - v[i].x
				col_diff := v[j].y - v[i].y
				next_candidate_row_1 := v[i].x - row_diff
				next_candidate_row_2 := v[j].x + row_diff
				next_candidate_col_1 := v[i].y - col_diff
				next_candidate_col_2 := v[j].y + col_diff
				if inbound(next_candidate_row_1, next_candidate_col_1, row, col) {

					if _, ok := visited[next_candidate_row_1*col+next_candidate_col_1]; !ok {
						count++
					}
					visited[next_candidate_row_1*col+next_candidate_col_1] = true
				}
				if inbound(next_candidate_row_2, next_candidate_col_2, row, col) {

					if _, ok := visited[next_candidate_row_2*col+next_candidate_col_2]; !ok {
						count++
					}
					visited[next_candidate_row_2*col+next_candidate_col_2] = true
				}

			}
		}

	}
	return count
}
func part2(grid []string) int {

	row := len(grid)
	col := len(grid[0])
	count := 0
	visited := make(map[int]bool, row)
	pairs := make(map[string][]Position)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] != '.' {
				pairs[string(grid[i][j])] = append(pairs[string(grid[i][j])], Position{i, j})
			}
		}
	}
	for _, v := range pairs {
		for i := 0; i < len(v); i++ {

			for j := i + 1; j < len(v); j++ {

				row_diff := v[j].x - v[i].x
				col_diff := v[j].y - v[i].y

				next_candidate_row_1 := v[i].x
				next_candidate_row_2 := v[j].x
				next_candidate_col_1 := v[i].y
				next_candidate_col_2 := v[j].y

				for inbound(next_candidate_row_1, next_candidate_col_1, row, col) {
					if _, ok := visited[next_candidate_row_1*col+next_candidate_col_1]; !ok {
						count++
					}
					visited[next_candidate_row_1*col+next_candidate_col_1] = true
					next_candidate_col_1 -= col_diff
					next_candidate_row_1 -= row_diff
				}
				for inbound(next_candidate_row_2, next_candidate_col_2, row, col) {
					if _, ok := visited[next_candidate_row_2*col+next_candidate_col_2]; !ok {
						count++
					}
					visited[next_candidate_row_2*col+next_candidate_col_2] = true
					next_candidate_row_2 += row_diff
					next_candidate_col_2 += col_diff
				}
			}
		}
	}

	return count
}

func inbound(x, y int, row, col int) bool {
	return x >= 0 && x < row && y >= 0 && y < col
}
