package main

import (
	"fmt"
	"os"
	"strings"
)

func contains(a []string, b string) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.ReadFile("23/23.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	graph := make(map[string][]string)

	// indegree := make(map[string]int)

	for _, line := range lines {
		nodes := strings.Split(line, "-")
		if graph[nodes[0]] == nil {
			graph[nodes[0]] = []string{}
		}
		if graph[nodes[1]] == nil {
			graph[nodes[1]] = []string{}
		}
		graph[nodes[0]] = append(graph[nodes[0]], nodes[1])
		graph[nodes[1]] = append(graph[nodes[1]], nodes[0])

	}
	count := 0
	fmt.Println()
	for k := range graph {
		fmt.Println(graph[k])
		for _, each_first := range graph[k] {
			fmt.Println(graph[each_first])
			for _, each_second := range graph[each_first] {
				fmt.Println(k, each_first, graph[each_second])
				fmt.Println()
				if contains(graph[each_second], k) {
					if k[0] == 't' || each_first[0] == 't' || each_second[0] == 't' {
						count++
					}
					// count++
					// fmt.Println(k, each_first, each_second)
				}
			}

			fmt.Println()
		}

	}

	fmt.Println(count / 6)

	// visited := make(map[string]bool)
	// ans := 0

	// for k := range items {

	// 	if _, ok := visited[k]; !ok {
	// 		t_count := 0
	// 		queue := make([]string, 0)
	// 		queue = append(queue, k)
	// 		queue_size := 0
	// 		visited[k] = true
	// 		for len(queue) > 0 {
	// 			// fmt.Println(queue)
	// 			queue_size++
	// 			current := queue[0]
	// 			queue = queue[1:]
	// 			if current[0] == 't' {
	// 				t_count++
	// 			}
	// 			for _, neigh := range graph[current] {
	// 				if _, ok := visited[neigh]; !ok {
	// 					visited[neigh] = true
	// 					queue = append(queue, neigh)
	// 				}
	// 			}

	// 		}
	// 		fmt.Println()

	// 		for queue_size >= 3 && t_count > 0 {
	// 			fmt.Println(queue_size, t_count)

	// 			left := queue_size - 2
	// 			answer := ((left + 1) * left) / 2
	// 			ans += answer

	// 			t_count--
	// 			queue_size--
	// 		}
	// 	}

	// }

}
