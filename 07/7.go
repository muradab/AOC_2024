package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Key struct {
	index int
	total int
}

func main() {
	file, err := os.ReadFile("07/day7.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func convertToIntArray(s string) []int {
	nums := strings.Split(s, " ")
	ints := make([]int, len(nums))
	for i, num := range nums {
		ints[i], _ = strconv.Atoi(num)
	}
	return ints
}

func checkForPart1(current int, required int, index int, array []int, memo map[Key]bool) bool {
	if current == required {
		return true
	}
	if index >= len(array) {
		return false
	}
	key := Key{index, current}
	if _, ok := memo[key]; ok {
		return memo[key]
	}

	memo[key] = checkForPart1(current+array[index], required, index+1, array, memo) ||
		checkForPart1(current*array[index], required, index+1, array, memo)

	return memo[key]

}

func checkForPart2(current int, required int, index int, array []int) bool {

	if index == len(array) {
		return current == required
	}
	if current > required {
		return false
	}

	currStr := strconv.Itoa(current)
	nextStr := strconv.Itoa(array[index])
	next, _ := strconv.Atoi(currStr + nextStr)

	return checkForPart2(current+array[index], required, index+1, array) ||
		checkForPart2(current*array[index], required, index+1, array) ||
		checkForPart2(next, required, index+1, array)
}

func part1(input []string) int {
	sum := 0

	for _, line := range input {
		parts := strings.Split(line, ": ")
		a, b := parts[0], parts[1]
		ints := convertToIntArray(b)
		required := convertToIntArray(a)
		memo := make(map[Key]bool)
		if checkForPart1(0, required[0], 0, ints, memo) {
			sum += required[0]
		}
	}
	return sum

}
func part2(input []string) int {
	sum := 0

	for _, line := range input {
		parts := strings.Split(line, ": ")
		a, b := parts[0], parts[1]
		ints := convertToIntArray(b)
		required := convertToIntArray(a)
		if checkForPart2(0, required[0], 0, ints) {
			sum += required[0]
		}

	}

	return sum

}
