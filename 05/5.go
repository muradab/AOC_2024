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
	partner := make([][]int, 100)

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
		partner[left] = append(partner[left], right)
	}
	ans := 0
	ans2 := 0
	for i := start + 1; i < len(item); i++ {
		disallowed := make(map[int]int)
		// split the string
		value := strings.Split(item[i], ",")
		part := convertToIntArray(value)
		good := true

		for _, num := range part {

			// convert to int

			if _, ok := disallowed[num]; ok {

				good = false
			}
			disallowed[num]++
			for _, v := range pairs[num] {
				disallowed[v]++
			}
		}
		if good {
			// take the middle element
			ans += part[len(part)/2]

		} else {

			news := make(map[int]int)
			for i := 0; i < len(part); i++ {
				for j := 0; j < len(part); j++ {
					if i == j {
						continue
					}
					for _, v := range partner[part[i]] {
						if part[j] == v {
							news[part[i]]++
						}
					}

				}
			}
			for k, v := range news {
				if v == len(part)/2 {
					ans2 += k
				}
			}
			fmt.Println(news)
		}
	}

	fmt.Println(start)

	fmt.Println(ans)
	fmt.Println(ans2)
}

func convertToIntArray(s []string) []int {
	var ans []int
	for _, v := range s {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		ans = append(ans, num)
	}
	return ans
}
