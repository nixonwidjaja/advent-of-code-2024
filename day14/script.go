package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partOne(P, V [][]int) {
	counts := make([]int, 4)
	n, m := 103, 101
	for i := 0; i < len(P); i++ {
		x, y, vx, vy := P[i][0], P[i][1], V[i][0], V[i][1]
		x = ((x+vx*100)%n + n) % n
		y = ((y+vy*100)%m + m) % m
		if x < n/2 && y < m/2 {
			counts[0]++
		} else if x > n/2 && y < m/2 {
			counts[1]++
		} else if x < n/2 && y > m/2 {
			counts[2]++
		} else if x > n/2 && y > m/2 {
			counts[3]++
		}
	}
	ans := 1
	for _, x := range counts {
		ans *= x
	}
	fmt.Println(ans)
}

var chx = []int{0, 0, 1, -1}
var chy = []int{1, -1, 0, 0}

func partTwo(P, V [][]int) {
	count, t := 0, 0
	n, m := 103, 101
	for count < int(math.Sqrt(float64(len(P)))) {
		graph := make([][]bool, n)
		for i := 0; i < n; i++ {
			graph[i] = make([]bool, m)
		}
		for i, p := range P {
			x, y, vx, vy := p[0], p[1], V[i][0], V[i][1]
			P[i][0] = ((x+vx)%n + n) % n
			P[i][1] = ((y+vy)%m + m) % m
			graph[P[i][0]][P[i][1]] = true
		}
		count = 0
		for i := 1; i < n-1; i++ {
			for j := 1; j < m-1; j++ {
				surrounded := true
				for k := 0; k < 4; k++ {
					newx, newy := i+chx[k], j+chy[k]
					if !graph[newx][newy] {
						surrounded = false
						break
					}
				}
				if surrounded {
					count++
				}
			}
		}
		t++
	}
	fmt.Println(t)
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	dataStr := strings.Split(string(data), "\n")
	P, V := make([][]int, len(dataStr)), make([][]int, len(dataStr))
	for i, x := range dataStr {
		splitted := strings.Fields(x)
		pStr := strings.Split(splitted[0], ",")
		px, err := strconv.Atoi(pStr[0][2:])
		check(err)
		py, err := strconv.Atoi(pStr[1])
		check(err)
		P[i] = []int{py, px}
		vStr := strings.Split(splitted[1], ",")
		vx, err := strconv.Atoi(vStr[0][2:])
		check(err)
		vy, err := strconv.Atoi(vStr[1])
		check(err)
		V[i] = []int{vy, vx}
	}
	partOne(P, V)
	partTwo(P, V)
}
