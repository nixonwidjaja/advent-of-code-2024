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

func partOne(locks, keys [][]int) {
	ans := 0
	for _, lock := range locks {
		for _, key := range keys {
			add := 1
			for i := 0; i < 5; i++ {
				if lock[i]+key[i] > 5 {
					add = 0
				}
			}
			ans += add
		}
	}
	fmt.Println(ans)
}

func main() {
	data, err := os.ReadFile("./input.txt")
	check(err)
	dataStr := strings.Split(string(data), "\n")
	locks, keys := make([][]int, 0), make([][]int, 0)
	isLock := false
	for i, x := range dataStr {
		if len(x) == 0 || i%8 == 6 {
			continue
		}
		if i%8 == 0 {
			if x == "#####" {
				isLock = true
				locks = append(locks, make([]int, 5))
			} else {
				isLock = false
				keys = append(keys, make([]int, 5))
			}
			continue
		}
		for j := 0; j < 5; j++ {
			if isLock {
				if x[j] == '#' {
					locks[len(locks)-1][j]++
				}
			} else {
				if x[j] == '#' {
					keys[len(keys)-1][j]++
				}
			}
		}
	}
	partOne(locks, keys)
}
