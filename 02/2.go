package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("02/day2.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	part1 := 0
	part2 := 0

	// split the file into lines
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		// split the line into slices
		parts := strings.Split(line, " ")

		// convert into an int slice
		ints := make([]int, len(parts))

		for i, part := range parts {
			ints[i], _ = strconv.Atoi(part)

		}

		fmt.Println(ints)

		increasing := isIncreasing(ints, true)
		decreasing := isIncreasing(reverse(ints), true)

		if increasing || decreasing {
			part2++
			part1++
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)

}

func isIncreasing(ints []int, firstTime bool) bool {
	fmt.Println(ints)

	for i := 1; i < len(ints); i++ {
		if ints[i] <= ints[i-1] || ints[i]-ints[i-1] > 3 {
			if firstTime {
				return isIncreasing(remove(ints, i), false) || isIncreasing(remove(ints, i-1), false)
			} else {

				return false
			}
		}
	}

	return true
}

func reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// huge pain
func remove(slice []int, s int) []int {
	newSlice := make([]int, len(slice)-1)
	copy(newSlice, slice[:s])
	copy(newSlice[s:], slice[s+1:])
	return newSlice

}
