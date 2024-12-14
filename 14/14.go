package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Velocity struct {
	x int
	y int
}

type Robot struct {
	position Position
	velocity Velocity
}

func main() {

	// read the input

	file, err := os.ReadFile("14/14.txt")
	if err != nil {
		panic(err)
	}

	output_file, err := os.Create("14/output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output_file.Close()

	robots := make([]Robot, 0)

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		r := strings.ReplaceAll(line, "p=", "")
		r = strings.ReplaceAll(r, " v=", " ")
		r = strings.ReplaceAll(r, ",", " ")
		item := strings.Split(r, " ")
		robots = append(robots, convertToRobot(item))

	}

	BOARD_HEIGHT := 103
	BOARD_WIDTH := 101

	MIDDLE_ROW := BOARD_WIDTH / 2
	MIDDLE_COL := BOARD_HEIGHT / 2

	// ELAPSED_TIME := 100

	all_quadrants := make([][]int, 0)

	grid := make([][]int, BOARD_WIDTH)
	items := make([][]byte, BOARD_WIDTH)

	// Fill each element with a byte array containing "."
	for i := range grid {
		for j := 0; j < BOARD_HEIGHT; j++ {
			grid[i] = append(grid[i], 0)
			items[i] = append(items[i], '.')
		}

	}

	for i := range robots {
		grid[robots[i].position.x][robots[i].position.y]++
		items[robots[i].position.x][robots[i].position.y] = '#'
	}

	quadrants := make([]int, 4)
	for ELAPSED_TIME := 0; ELAPSED_TIME < BOARD_HEIGHT*BOARD_WIDTH; ELAPSED_TIME++ {
		temp := make([]Robot, 0)

		for r := range robots {
			new_x := robots[r].position.x + robots[r].velocity.x
			new_y := robots[r].position.y + robots[r].velocity.y
			new_x %= BOARD_WIDTH
			new_y %= BOARD_HEIGHT
			new_x += BOARD_WIDTH
			new_y += BOARD_HEIGHT
			new_x %= BOARD_WIDTH
			new_y %= BOARD_HEIGHT
			temp = append(temp, Robot{Position{new_x, new_y}, robots[r].velocity})
			grid[new_x][new_y]++
			items[new_x][new_y] = '#'
			grid[robots[r].position.x][robots[r].position.y]--
			if grid[robots[r].position.x][robots[r].position.y] == 0 {
				items[robots[r].position.x][robots[r].position.y] = '.'

			}

			fmt.Println(new_x, new_y)
			fmt.Println()

			if new_x < MIDDLE_ROW && new_y < MIDDLE_COL {
				quadrants[0]++
			}
			if new_x > MIDDLE_ROW && new_y < MIDDLE_COL {
				quadrants[1]++
			}
			if new_x < MIDDLE_ROW && new_y > MIDDLE_COL {
				quadrants[2]++
			}
			if new_x > MIDDLE_ROW && new_y > MIDDLE_COL {
				quadrants[3]++
			}

		}
		robots = temp
		for i := range items {

			fmt.Fprintln(output_file, string(items[i]))
			// idea is find the time with the maximum row of # I used vs code

		}
		fmt.Fprintln(output_file, ELAPSED_TIME)
		fmt.Println(ELAPSED_TIME)

	}

	fmt.Println(all_quadrants)

	// fmt.Println(res)

}

func convertToRobot(s []string) Robot {

	px, _ := strconv.Atoi(s[0])
	py, _ := strconv.Atoi(s[1])
	vx, _ := strconv.Atoi(s[2])
	vy, _ := strconv.Atoi(s[3])

	return Robot{
		Position{px, py},
		Velocity{vx, vy},
	}
}
