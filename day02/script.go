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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func eval(arr []int) bool {
	incr := arr[1] > arr[0]
	for i, x := range arr {
		if i == 0 {
			continue
		}
		if (x <= arr[i-1] && incr) || (x >= arr[i-1] && !incr) {
			return false
		}
		if abs(x-arr[i-1]) > 3 {
			return false
		}
	}
	return true
}

func dampen(arr []int) bool {
	for i := range arr {
		newArr := make([]int, 0)
		for j := range arr {
			if j == i {
				continue
			}
			newArr = append(newArr, arr[j])
		}
		if eval(newArr) {
			return true
		}
	}
	return false
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	lines := strings.Split(string(data), "\n")
	ans, ansDamp := 0, 0
	for _, line := range lines {
		strs := strings.Fields(line)
		nums := make([]int, len(strs))
		for i, x := range strs {
			num, err := strconv.Atoi(x)
			check(err)
			nums[i] = num
		}
		if eval(nums) {
			ans++
			ansDamp++
		} else if dampen(nums) {
			ansDamp++
		}
	}
	fmt.Println(ans)
	fmt.Println(ansDamp)
}