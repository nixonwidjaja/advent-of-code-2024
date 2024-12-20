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

func try(n int) [][]string {
	res := make([][]string, 0)
	var dfs func(i int, now []string)
	dfs = func(i int, now []string) {
		if i == n {
			copied := make([]string, len(now))
			copy(copied, now)
			res = append(res, copied)
			return
		}
		now = append(now, "+")
		dfs(i+1, now)
		now = now[:len(now)-1]
		now = append(now, "*")
		dfs(i+1, now)
		now = now[:len(now)-1]
		now = append(now, "|")
		dfs(i+1, now)
		now = now[:len(now)-1]
	}
	dfs(0, make([]string, 0))
	return res
}

func partTwo(arr []int, symbols []string) int {
	ans := arr[0]
	for i, symbol := range symbols {
		switch symbol {
		case "+":
			ans += arr[i+1]
		case "*":
			ans *= arr[i+1]
		case "|":
			ans *= int(math.Pow(10, float64(len(strconv.Itoa(arr[i+1])))))
			ans += arr[i+1]
		}
	}
	return ans
}

func solve(targets []int, nums [][]int) {
	ans := 0
	for i := 0; i < len(targets); i++ {
		target, num := targets[i], nums[i]
		symbolArr := try(len(num) - 1)
		for _, symbols := range symbolArr {
			if partTwo(num, symbols) == target {
				ans += target
				break
			}
		}
	}
	fmt.Println(ans)
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	dataStr := strings.Split(string(data), "\n")
	targets := make([]int, len(dataStr))
	nums := make([][]int, len(dataStr))
	for i, x := range dataStr {
		numStr := strings.Fields(x)
		target, err := strconv.Atoi(numStr[0][:len(numStr[0])-1])
		check(err)
		targets[i] = target
		for j, n := range numStr {
			if j == 0 {
				continue
			}
			nInt, err := strconv.Atoi(n)
			check(err)
			nums[i] = append(nums[i], nInt)
		}
	}
	solve(targets, nums)
}
