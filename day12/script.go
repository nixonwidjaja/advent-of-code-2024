package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var chx = []int{0, 0, 1, -1}
var chy = []int{1, -1, 0, 0}

func dfs(x, y int, graph []string, visited [][]bool) (int, int) {
	visited[x][y] = true
	area, perimeter := 1, 0
	for i := 0; i < 4; i++ {
		newx, newy := x+chx[i], y+chy[i]
		if newx >= 0 && newx < len(graph) && newy >= 0 && newy < len(graph[0]) {
			if graph[newx][newy] == graph[x][y] {
				if !visited[newx][newy] {
					newArea, newPerimeter := dfs(newx, newy, graph, visited)
					area += newArea
					perimeter += newPerimeter
				}
			} else {
				perimeter++
			}
		} else {
			perimeter++
		}
	}
	return area, perimeter
}

func partOne(graph []string) {
	n, m := len(graph), len(graph[0])
	visited := make([][]bool, n)
	for i := range graph {
		visited[i] = make([]bool, m)
	}
	ans := 0
	for i := range graph {
		for j := range graph[i] {
			if !visited[i][j] {
				area, perimeter := dfs(i, j, graph, visited)
				ans += area * perimeter
			}
		}
	}
	fmt.Println(ans)
}

func dfsTwo(x, y int, graph []string, visited [][]bool, sides [][][]int) int {
	visited[x][y] = true
	area := 1
	for i := 0; i < 4; i++ {
		newx, newy := x+chx[i], y+chy[i]
		if newx >= 0 && newx < len(graph) && newy >= 0 && newy < len(graph[0]) {
			if graph[newx][newy] == graph[x][y] {
				if !visited[newx][newy] {
					area += dfsTwo(newx, newy, graph, visited, sides)
				}
			} else {
				sides[i] = append(sides[i], []int{x, y})
			}
		} else {
			sides[i] = append(sides[i], []int{x, y})
		}
	}
	return area
}

func countSides(sides [][][]int) int {
	count := 0
	for i, pts := range sides {
		d := make(map[int][]int)
		for _, pt := range pts {
			x, y := pt[0], pt[1]
			if i < 2 {
				d[y] = append(d[y], x)
			} else {
				d[x] = append(d[x], y)
			}
		}
		for _, v := range d {
			sort.Slice(v, func(a, b int) bool {
				return v[a] < v[b]
			})
			count++
			for j, x := range v {
				if j == 0 {
					continue
				}
				if x > v[j-1]+1 {
					count++
				}
			}
		}
	}
	return count
}

func partTwo(graph []string) {
	n, m := len(graph), len(graph[0])
	visited := make([][]bool, n)
	for i := range graph {
		visited[i] = make([]bool, m)
	}
	ans := 0
	for i := range graph {
		for j := range graph[i] {
			if !visited[i][j] {
				sides := make([][][]int, 4)
				area := dfsTwo(i, j, graph, visited, sides)
				ans += area * countSides(sides)
			}
		}
	}
	fmt.Println(ans)
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	graph := strings.Split(string(data), "\n")
	partOne(graph)
	partTwo(graph)
}
