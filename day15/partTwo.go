package main

func partTwo(graph [][]string, moves []string) {
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
