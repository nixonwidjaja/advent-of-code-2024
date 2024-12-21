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

func brute(a, b, prize []int) int {
	for j := 0; j <= 100; j++ {
		for k := 0; k <= 100; k++ {
			testX, testY := a[0]*j+b[0]*k, a[1]*j+b[1]*k
			if testX == prize[0] && testY == prize[1] {
				return 3*j + k
			}
		}
	}
	return 0
}

func partOne(A, B, prizes [][]int) {
	ans := 0
	for i := 0; i < len(A); i++ {
		a, b, prize := A[i], B[i], prizes[i]
		ans += brute(a, b, prize)
	}
	fmt.Println(ans)
}

func solve(a, b, prize []int) int {
	prize[0] += 10000000000000
	prize[1] += 10000000000000
	x := (prize[0]*b[1] - prize[1]*b[0]) / (a[0]*b[1] - a[1]*b[0])
	y := (prize[0]*a[1] - prize[1]*a[0]) / (b[0]*a[1] - b[1]*a[0])
	xMod := (prize[0]*b[1] - prize[1]*b[0]) % (a[0]*b[1] - a[1]*b[0])
	yMod := (prize[0]*a[1] - prize[1]*a[0]) % (b[0]*a[1] - b[1]*a[0])
	if x >= 0 && y >= 0 && xMod == 0 && yMod == 0 {
		return 3*x + y
	}
	return 0
}

func partTwo(A, B, prizes [][]int) {
	ans := 0
	for i := 0; i < len(A); i++ {
		a, b, prize := A[i], B[i], prizes[i]
		ans += solve(a, b, prize)
	}
	fmt.Println(ans)
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	dataStr := strings.Split(string(data), "\n")
	prompts := make([][]string, len(dataStr))
	for i, x := range dataStr {
		prompts[i] = strings.Fields(x)
	}
	A := make([][]int, 0)
	B := make([][]int, 0)
	prizes := make([][]int, 0)
	for i := 0; i < len(prompts); i++ {
		prompt := prompts[i]
		switch i % 4 {
		case 0:
			x, err := strconv.Atoi(prompt[2][2 : len(prompt[2])-1])
			check(err)
			y, err := strconv.Atoi(prompt[3][2:])
			check(err)
			A = append(A, []int{x, y})
		case 1:
			x, err := strconv.Atoi(prompt[2][2 : len(prompt[2])-1])
			check(err)
			y, err := strconv.Atoi(prompt[3][2:])
			check(err)
			B = append(B, []int{x, y})
		case 2:
			x, err := strconv.Atoi(prompt[1][2 : len(prompt[1])-1])
			check(err)
			y, err := strconv.Atoi(prompt[2][2:])
			check(err)
			prizes = append(prizes, []int{x, y})
		}
	}
	partOne(A, B, prizes)
	partTwo(A, B, prizes)
}
