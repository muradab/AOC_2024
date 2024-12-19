package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.ReadFile("19/19.txt")

	if err != nil {
		panic(err)
	}
	input := strings.Split(string(file), "\n\n")

	words := strings.Split(input[0], ", ")
	combinations := strings.Split(input[1], "\n")

	wordMap := make(map[string]bool)
	for _, w := range words {
		wordMap[w] = true
	}

	count := 0

	for c := range combinations {
		memo := make(map[int]bool)
		if res := dp(0, memo, combinations[c], wordMap); res {
			count++
		}
	}
	sum := 0
	for c := range combinations {
		memo := make(map[int]int)
		sum += dp2(0, memo, combinations[c], wordMap)

	}

	fmt.Println(combinations)
	fmt.Println(words)
	fmt.Println(count)
	fmt.Println(sum)

}

func dp(index int, memo map[int]bool, s string, words map[string]bool) bool {
	if index == len(s) {
		return true
	}
	if _, ok := memo[index]; ok {
		return memo[index]
	}

	ans := false
	for i := index; i < min(len(s), index+10); i++ {
		if _, ok := words[s[index:i+1]]; ok {
			ans = ans || dp(i+1, memo, s, words)
		}
	}

	memo[index] = ans
	return ans

}
func dp2(index int, memo map[int]int, s string, words map[string]bool) int {
	if index == len(s) {
		return 1
	}
	if _, ok := memo[index]; ok {
		return memo[index]
	}

	ans := 0
	for i := index; i < min(len(s), index+10); i++ {
		if _, ok := words[s[index:i+1]]; ok {
			ans += dp2(i+1, memo, s, words)
		}
	}

	memo[index] = ans
	return ans

}
