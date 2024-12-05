package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// read an input
	file, err := os.ReadFile("05/day5.txt")

	if err != nil {
		panic(err)
	}

	pairs := make([][]int, 100)

	item := (strings.Split(string(file), "\n"))
	start := 0
	for i, v := range item {

		part := strings.Split(v, "|")
		start = i
		if len(part) != 2 {
			break
		}
		// convert to int
		left, err := strconv.Atoi(part[0])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(part[1])
		if err != nil {
			panic(err)
		}
		pairs[right] = append(pairs[right], left)
	}
	ans := 0

	for i := start + 1; i < len(item); i++ {
		disallowed := make(map[int]int)
		// split the string
		part := strings.Split(item[i], ",")
		good := true
		fmt.Println(part)

		for _, v := range part {
			fmt.Println(disallowed)
			// convert to int
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			if _, ok := disallowed[num]; ok {
				fmt.Println("been here", num)
				good = false
			}
			disallowed[num]++
			for _, v := range pairs[num] {
				disallowed[v]++
			}
		}
		if good {
			// take the middle element
			item, _ := strconv.Atoi(part[len(part)/2])
			ans += item
		}

	}
	fmt.Println(start)

	fmt.Println(ans)
}
