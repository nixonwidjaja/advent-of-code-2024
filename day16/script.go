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

var chx = []int{0, 0, 1, -1}
var chy = []int{1, -1, 0, 0}

func partTwo(graph []string) {
	startx, starty, endx, endy := 0, 0, 0, 0
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == 'S' {
				startx, starty = i, j
			}
			if graph[i][j] == 'E' {
				endx, endy = i, j
			}
		}
	}
	visited := make([][][]int, len(graph))
	for i := range visited {
		visited[i] = make([][]int, len(graph[0]))
		for j := range visited[i] {
			visited[i][j] = []int{1e9, 1e9, 1e9, 1e9}
		}
	}
	pq := newHeap()
	pq.push(state{x: startx, y: starty, distance: 0, path: [][]int{{startx, starty}}})
	bestDistance := int(1e9)
	travelled := make([][]bool, len(graph))
	for i := range travelled {
		travelled[i] = make([]bool, len(graph[0]))
	}
	for pq.len() > 0 {
		now := pq.pop()
		if now.x == endx && now.y == endy {
			if bestDistance >= now.distance {
				fmt.Println(now.distance)
				bestDistance = now.distance
				for _, pt := range now.path {
					travelled[pt[0]][pt[1]] = true
				}
			} else {
				break
			}
		}
		newx, newy := now.x+chx[now.dir], now.y+chy[now.dir]
		for i := 0; i < 4; i++ {
			if i == now.dir {
				if graph[newx][newy] != '#' && now.distance+1 <= visited[newx][newy][now.dir] {
					visited[newx][newy][now.dir] = now.distance + 1
					newPath := make([][]int, len(now.path))
					copy(newPath, now.path)
					newPath = append(newPath, []int{newx, newy})
					pq.push(state{x: newx, y: newy, dir: now.dir, distance: now.distance + 1, path: newPath})
				}
			} else {
				if now.distance+1000 <= visited[now.x][now.y][i] {
					visited[now.x][now.y][i] = now.distance + 1000
					pq.push(state{x: now.x, y: now.y, dir: i, distance: now.distance + 1000, path: now.path})
				}
			}
		}
	}
	count := 0
	for i := range travelled {
		for j := range travelled[i] {
			if travelled[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	graph := strings.Split(string(data), "\n")
	partTwo(graph)
}
