package main

func moveRight(x, y int, graph [][]string) (int, int) {
	newx, newy := x, y+1
	switch graph[newx][newy] {
	case "#":
		return x, y
	case ".":
		graph[newx][newy] = "@"
		graph[x][y] = "."
		return newx, newy
	default:
		for curr := newy; graph[newx][curr] != "#"; curr++ {
			if graph[newx][curr] == "." {
				graph[newx][curr] = "O"
				graph[newx][newy] = "@"
				graph[x][y] = "."
				return newx, newy
			}
		}
		return x, y
	}
}

func moveLeft(x, y int, graph [][]string) (int, int) {
	newx, newy := x, y-1
	switch graph[newx][newy] {
	case "#":
		return x, y
	case ".":
		graph[newx][newy] = "@"
		graph[x][y] = "."
		return newx, newy
	default:
		for curr := newy; graph[newx][curr] != "#"; curr-- {
			if graph[newx][curr] == "." {
				graph[newx][curr] = "O"
				graph[newx][newy] = "@"
				graph[x][y] = "."
				return newx, newy
			}
		}
		return x, y
	}
}

func moveUp(x, y int, graph [][]string) (int, int) {
	newx, newy := x-1, y
	switch graph[newx][newy] {
	case "#":
		return x, y
	case ".":
		graph[newx][newy] = "@"
		graph[x][y] = "."
		return newx, newy
	default:
		for curr := newx; graph[curr][newy] != "#"; curr-- {
			if graph[curr][newy] == "." {
				graph[curr][newy] = "O"
				graph[newx][newy] = "@"
				graph[x][y] = "."
				return newx, newy
			}
		}
		return x, y
	}
}

func moveDown(x, y int, graph [][]string) (int, int) {
	newx, newy := x+1, y
	switch graph[newx][newy] {
	case "#":
		return x, y
	case ".":
		graph[newx][newy] = "@"
		graph[x][y] = "."
		return newx, newy
	default:
		for curr := newx; graph[curr][newy] != "#"; curr++ {
			if graph[curr][newy] == "." {
				graph[curr][newy] = "O"
				graph[newx][newy] = "@"
				graph[x][y] = "."
				return newx, newy
			}
		}
		return x, y
	}
}

func makeMove(x, y int, graph [][]string, move string) (int, int) {
	switch move {
	case ">":
		return moveRight(x, y, graph)
	case "<":
		return moveLeft(x, y, graph)
	case "^":
		return moveUp(x, y, graph)
	case "v":
		return moveDown(x, y, graph)
	}
	return 0, 0
}

func partOne(graph [][]string, moves []string) {
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == "@" {
				x, y := i, j
				for _, move := range moves {
					x, y = makeMove(x, y, graph, move)
				}
				return
			}
		}
	}
}
