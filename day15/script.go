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

func main() {
	data, err := os.ReadFile("./test")
	check(err)
	dataStr := strings.Split(string(data), "\n")
	graph := make([][]string, len(dataStr)-2)
	for i, x := range dataStr {
		if i == len(dataStr)-2 {
			break
		}
		graph[i] = strings.Split(x, "")
	}
	moves := strings.Split(dataStr[len(dataStr)-1], "")
	partOne(graph, moves)
	ans := 0
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == "O" {
				ans += 100*i + j
			}
		}
	}
	fmt.Println(ans)
	for i, x := range dataStr {
		if i == len(dataStr)-2 {
			break
		}
		graph[i] = strings.Split(x, "")
	}
	partTwo(graph, moves)
}
