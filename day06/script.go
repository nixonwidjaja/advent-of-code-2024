package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var chx = []int{0, 1, 0, -1}
var chy = []int{1, 0, -1, 0}

func dfs(x, y, dir int, graph [][]string) {
	graph[x][y] = "X"
	newx, newy := x+chx[dir], y+chy[dir]
	n, m := len(graph), len(graph[0])
	if newx >= 0 && newx < n && newy >= 0 && newy < m {
		if graph[newx][newy] == "#" {
			dfs(x, y, (dir+1)%4, graph)
		} else {
			dfs(newx, newy, dir, graph)
		}
	}
}

func partOne(graph [][]string) {
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == "^" {
				dfs(i, j, 3, graph)
			}
		}
	}
	ans := 0
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == "X" {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func play(x, y, dir int, graph [][]string, visited [][130][4]bool) bool {
	if visited[x][y][dir] {
		return true
	}
	visited[x][y][dir] = true
	newx, newy := x+chx[dir], y+chy[dir]
	n, m := len(graph), len(graph[0])
	if newx >= 0 && newx < n && newy >= 0 && newy < m {
		if graph[newx][newy] == "#" {
			if play(x, y, (dir+1)%4, graph, visited) {
				return true
			}
		} else {
			if play(newx, newy, dir, graph, visited) {
				return true
			}
		}
	}
	return false
}

func partTwo(graph [][]string) {
	startx, starty := 0, 0
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == "^" {
				startx, starty = i, j
			}
		}
	}
	ans := 0
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == "." {
				visited := make([][130][4]bool, 130)
				graph[i][j] = "#"
				if play(startx, starty, 3, graph, visited) {
					ans++
				}
				graph[i][j] = "."
			}
		}
	}
	fmt.Println(ans)
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	dataStr := strings.Split(string(data), "\n")
	graph := make([][]string, len(dataStr))
	for i, x := range dataStr {
		graph[i] = strings.Split(x, "")
	}
	partOne(graph)
	for i, x := range dataStr {
		graph[i] = strings.Split(x, "")
	}
	partTwo(graph)
}
