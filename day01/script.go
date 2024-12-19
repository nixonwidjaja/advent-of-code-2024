package main

import (
	"fmt"
	"os"
	"sort"
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

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	arr := strings.Fields(string(data))
	left := make([]int, len(arr)/2)
	right := make([]int, len(arr)/2)
	for i, x := range arr {
		num, err := strconv.Atoi(x)
		check(err)
		if i % 2 == 0 {
			left[i/2] =num
		} else {
			right[i/2] = num
		}
	}
	sort.Ints(left)
	sort.Ints(right)
	ans := 0
	for i:=0;i<len(left);i++ {
		ans += abs(right[i] - left[i])
	}
	fmt.Println(ans)

	count := make(map[int]int)
	for _, x := range right {
		count[x] ++
	}
	sim := 0
	for _, x := range left {
		sim += x * count[x]
	}
	fmt.Println(sim)
}