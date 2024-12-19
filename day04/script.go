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

func findXmas(arr [][]string, x, y int) int {
	chx := []int{0, 0, 1, 1, 1, -1, -1, -1}
	chy := []int{1, -1, 0, 1, -1, 0, 1, -1}
	xmas := []string{"X", "M", "A", "S"}
	n, m, ans := len(arr), len(arr[0]), 0
	for i := 0; i < 8; i++ {
		for j := 1; j < 4; j++ {
			newx, newy := x+chx[i]*j, y+chy[i]*j
			if newx < 0 || newx >= n || newy < 0 || newy >= m || xmas[j] != arr[newx][newy] {
				break
			}
			if j == 3 {
				ans++
			}
		}
	}
	return ans
}

func findMas(arr [][]string, x, y int) bool {
	chx := []int{1, 1, -1, -1}
	chy := []int{1, -1, 1, -1}
	pos := []string{"MMSS", "MSMS", "SMSM", "SSMM"}
	n, m, res := len(arr), len(arr[0]), ""
	for i := 0; i < 4; i++ {
		newx, newy := x+chx[i], y+chy[i]
		if newx < 0 || newx >= n || newy < 0 || newy >= m {
			return false
		}
		res += arr[newx][newy]
	}
	for _, s := range pos {
		if res == s {
			return true
		}
	}
	return false
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	dataStr := string(data)
	dataArr := strings.Split(dataStr, "\n")
	arr := make([][]string, len(dataArr))
	for i, s := range dataArr {
		arr[i] = strings.Split(s, "")
	}
	n, m, ans, mas := len(arr), len(arr[0]), 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == "X" {
				ans += findXmas(arr, i, j)
			}
			if arr[i][j] == "A" && findMas(arr, i, j) {
				mas++
			}
		}
	}
	fmt.Println(ans, mas)
}
