package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.ReadFile("09/day9.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	number := string(file)
	item := make([]string, 0)
	spaces := make([][]int, 0)
	num_sizes := make([][]int, 0)

	for i := 0; i < len(number); i++ {
		num, _ := strconv.Atoi(string(number[i]))

		if i%2 == 0 {
			num_sizes = append(num_sizes, []int{num, sum})

			sum += num
			for j := 0; j < num; j++ {

				item = append(item, strconv.Itoa(i/2))
			}
		} else {
			spaces = append(spaces, []int{sum, sum + num})
			sum += num
			for j := 0; j < num; j++ {
				item = append(item, ".")
			}
		}
	}

	// fmt.Println(sum)
	// fmt.Println(spaces)
	// fmt.Println(item)
	// fmt.Println(num_sizes)
	// fmt.Println(part1(item))
	fmt.Println(part2(spaces, num_sizes, item))
	// fmt.Println(part1(item))

}

func part1(item []string) int {
	left := 0
	right := len(item) - 1

	for left < right {
		if item[left] != "." {
			left++
		} else if item[right] == "." {
			right--
		} else {
			item[left], item[right] = item[right], item[left]
			left++
			right--
		}
	}

	ans := 0

	for i, v := range item {
		num, err := strconv.Atoi(v)
		if err != nil {
			break
		}
		ans += num * i

	}
	return ans
}

func part2(spaces, num_sizes [][]int, item []string) int {

	for i := len(num_sizes) - 1; i >= 0; i-- {

		limit := num_sizes[i][1]
		for j := 0; j < len(spaces); j++ {
			if spaces[j][0] < limit && spaces[j][1]-spaces[j][0] >= num_sizes[i][0] {

				for k := 0; k < num_sizes[i][0]; k++ {
					item[spaces[j][0]+k] = strconv.Itoa(i)
				}
				spaces[j][0] = spaces[j][0] + num_sizes[i][0]

				for k := 0; k < num_sizes[i][0]; k++ {
					item[num_sizes[i][1]+k] = "."
				}

				break
			}

		}

	}

	ans := 0

	for i := 0; i < len(item); i++ {
		num, err := strconv.Atoi(item[i])
		if err != nil {
			continue
		}
		ans += num * i
	}

	return ans
}
