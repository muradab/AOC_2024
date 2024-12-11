package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("11/day11.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), " ")

	nums := make([]int, 0)
	for _, v := range lines {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	// ans := 0

	ans := 0

	memo := make(map[State]int)

	for i, num := range nums {
		fmt.Println(i)
		ans += dfs(num, 75, memo)

	}

	fmt.Println(ans)

}

type State struct {
	num, height int
}

func dfs(num, height int, memo map[State]int) int {

	if height == 0 {
		return 1
	}
	state := State{num, height}

	if val, ok := memo[state]; ok {
		return val
	}
	digit_length := len(strconv.Itoa(num))

	if num == 0 {
		memo[state] = dfs(num+1, height-1, memo)
	} else if digit_length%2 == 1 {
		return dfs(num*2024, height-1, memo)
	} else {

		num_str := strconv.Itoa(num)
		left := num_str[0 : digit_length/2]
		right := num_str[digit_length/2:]

		left_num, _ := strconv.Atoi(left)
		right_num, _ := strconv.Atoi(right)

		memo[state] = dfs(left_num, height-1, memo) + dfs(right_num, height-1, memo)
	}
	return memo[state]
}
