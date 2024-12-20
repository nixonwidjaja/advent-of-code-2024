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

func findAntinodeOne(a, b []int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{2*a[0] - b[0], 2*a[1] - b[1]})
	ans = append(ans, []int{2*b[0] - a[0], 2*b[1] - a[1]})
	return ans
}

func findAntinodeTwo(a, b []int, n, m int) [][]int {
	ans := make([][]int, 0)
	for i := 0; ; i++ {
		x, y := a[0]+i*(a[0]-b[0]), a[1]+i*(a[1]-b[1])
		if x < 0 || x >= n || y < 0 || y >= m {
			break
		}
		ans = append(ans, []int{x, y})
	}
	for i := -1; ; i-- {
		x, y := a[0]+i*(a[0]-b[0]), a[1]+i*(a[1]-b[1])
		if x < 0 || x >= n || y < 0 || y >= m {
			break
		}
		ans = append(ans, []int{x, y})
	}
	return ans
}

func solve(graph []string) {
	n, m := len(graph), len(graph[0])
	locations := make(map[byte][][]int)
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] != '.' {
				locations[graph[i][j]] = append(locations[graph[i][j]], []int{i, j})
			}
		}
	}
	antinodes := make([][]bool, n)
	for i := range antinodes {
		antinodes[i] = make([]bool, m)
	}
	for _, v := range locations {
		for i, a := range v {
			for j, b := range v {
				if i == j {
					continue
				}
				antinode := findAntinodeTwo(a, b, n, m)
				for _, pt := range antinode {
					x, y := pt[0], pt[1]
					if x >= 0 && x < n && y >= 0 && y < m {
						antinodes[x][y] = true
					}
				}
			}
		}
	}
	count := 0
	for i := range antinodes {
		for j := range antinodes[i] {
			if antinodes[i][j] {
				count += 1
			}
		}
	}
	fmt.Println(count)
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	graph := strings.Split(string(data), "\n")
	solve(graph)
}
