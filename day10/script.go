package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var chx = []int{0, 0, 1, -1}
var chy = []int{1, -1, 0, 0}

func dfs(x, y int, graph [][]int, visited [][]bool) int {
	count := 0
	if graph[x][y] == 9 && !visited[x][y] {
		visited[x][y] = true
		count = 1
	}
	for i := 0; i < 4; i++ {
		newx, newy := x+chx[i], y+chy[i]
		if newx >= 0 && newx < len(graph) && newy >= 0 && newy < len(graph[0]) && graph[newx][newy] == graph[x][y]+1 {
			count += dfs(newx, newy, graph, visited)
		}
	}
	return count
}

func calculateScore(graph [][]int) {
	score := 0
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == 0 {
				visited := make([][]bool, len(graph))
				for k := range graph {
					visited[k] = make([]bool, len(graph[0]))
				}
				score += dfs(i, j, graph, visited)
			}
		}
	}
	fmt.Println(score)
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	dataStr := strings.Split(string(data), "\n")
	graph := make([][]int, len(dataStr))
	for i, x := range dataStr {
		str := strings.Split(x, "")
		nums := make([]int, 0)
		for _, s := range str {
			num, err := strconv.Atoi(s)
			check(err)
			nums = append(nums, num)
		}
		graph[i] = nums
	}
	calculateScore(graph)
}
