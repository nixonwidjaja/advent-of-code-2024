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

func partOne(arr []int) {
	for k := 0; k < 25; k++ {
		nums := make([]int, 0)
		for _, x := range arr {
			if x == 0 {
				nums = append(nums, 1)
			} else if s := strconv.Itoa(x); len(s)%2 == 0 {
				l, err := strconv.Atoi(s[:len(s)/2])
				check(err)
				r, err := strconv.Atoi(s[len(s)/2:])
				check(err)
				nums = append(nums, l, r)
			} else {
				nums = append(nums, x*2024)
			}
		}
		arr = nums
	}
	fmt.Println(len(arr))
}

func partTwo(arr []int) {
	d := make(map[int]int)
	for _, x := range arr {
		d[x]++
	}
	for k := 0; k < 75; k++ {
		dNew := make(map[int]int)
		for x, count := range d {
			if x == 0 {
				dNew[1] += count
			} else if s := strconv.Itoa(x); len(s)%2 == 0 {
				l, err := strconv.Atoi(s[:len(s)/2])
				check(err)
				r, err := strconv.Atoi(s[len(s)/2:])
				check(err)
				dNew[l] += count
				dNew[r] += count
			} else {
				dNew[x*2024] += count
			}
		}
		d = dNew
	}
	ans := 0
	for _, count := range d {
		ans += count
	}
	fmt.Println(ans)
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	dataStr := strings.Fields(string(data))
	nums := make([]int, len(dataStr))
	for i, x := range dataStr {
		num, err := strconv.Atoi(x)
		check(err)
		nums[i] = num
	}
	numsOne := make([]int, len(nums))
	copy(numsOne, nums)
	partOne(numsOne)
	partTwo(nums)
}
