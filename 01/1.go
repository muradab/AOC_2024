package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	// Open the file
	file, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0
	left := make([]int, 0)
	right := make([]int, 0)
	right_count := make(map[int]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		first := parts[0]
		second := parts[1]
		firstInt, err := strconv.ParseInt(first, 10, 32)
		if err != nil {
			fmt.Println("Error parsing first part:", err)
			return
		}

		secondInt, err := strconv.ParseInt(second, 10, 32)
		if err != nil {
			fmt.Println("Error parsing first part:", err)
			return
		}
		right_count[int(secondInt)]++
		left = append(left, int(firstInt))
		right = append(right, int(secondInt))

	}

	// sort the slices
	slices.Sort(left)
	slices.Sort(right)
	fmt.Println(len(left))
	fmt.Println(right_count)
	for i := 0; i < len(left); i++ {

		fmt.Println(left[i], right[i])
		total += left[i] * right_count[left[i]]

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Total:", total)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
