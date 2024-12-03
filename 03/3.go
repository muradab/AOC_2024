package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("03/day3.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	fmt.Println(len(lines))
	part1 := 0

	for _, line := range lines {
		// create a stack of strings
		stack := make([]int, 0)
		ok := false
		start := 0

		for i, val := range line {
			item := string(val)
			fmt.Println(item)
			if ok {

				if start+1 == i && item == "(" {
					stack = append(stack, 0)
				} else if start+1 == i && item != "(" {
					ok = false
					stack = []int{}
				} else if number, err := strconv.Atoi(string(item)); err == nil {
					lastItem := stack[len(stack)-1]
					newNum := lastItem*10 + number
					stack[len(stack)-1] = newNum
				} else if item == ")" && len(stack) == 2 {
					part1 += stack[0] * stack[1]
					ok = false
					stack = []int{}
				} else if item == "," && len(stack) == 1 {
					stack = append(stack, 0)
				} else {
					ok = false
					stack = []int{}
				}
			}
			if item == "l" && i > 2 {
				if string(line[i-2]) == "m" && string(line[i-1]) == "u" {
					ok = true
					start = i

				}
			}

		}
	}
	fmt.Println(part1)
}
