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

func evolve(secret int) int {
	secret = prune(mix(secret*64, secret))
	secret = prune(mix(secret>>5, secret))
	secret = prune(mix(secret*2048, secret))
	return secret
}

func process(secret int) int {
	for i := 0; i < 2000; i++ {
		secret = evolve(secret)
	}
	return secret
}

func mix(secret, num int) int {
	return secret ^ num
}

func prune(secret int) int {
	return secret % 16777216
}

func partOne(nums []int) {
	ans := 0
	for _, x := range nums {
		ans += process(x)
	}
	fmt.Println(ans)
}

func sequence(secret int) ([]int, []int) {
	prices := make([]int, 2000)
	diff := make([]int, 2000)
	for i := 0; i < 2000; i++ {
		oldSecret := secret % 10
		secret = evolve(secret)
		prices[i] = secret % 10
		diff[i] = secret%10 - oldSecret
	}
	return prices, diff
}

func join(nums []int) string {
	numStr := make([]string, len(nums))
	for i, x := range nums {
		numStr[i] = strconv.Itoa(x)
	}
	return strings.Join(numStr, ",")
}

func accumulate(prices, diff []int) map[string]int {
	d := make(map[string]int)
	for i := range prices {
		if i+3 >= len(prices) {
			break
		}
		priceStr := join(diff[i : i+4])
		if _, ok := d[priceStr]; !ok {
			d[priceStr] = prices[i+3]
		}
	}
	return d
}

func partTwo(nums []int) {
	allDiff := make(map[string]int)
	ans := 0
	for _, secret := range nums {
		prices, diff := sequence(secret)
		d := accumulate(prices, diff)
		for priceStr, price := range d {
			allDiff[priceStr] += price
			if allDiff[priceStr] > ans {
				ans = allDiff[priceStr]
			}
		}
	}
	fmt.Println(ans)
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	dataStr := strings.Split(string(data), "\n")
	nums := make([]int, len(dataStr))
	for i, x := range dataStr {
		nums[i], err = strconv.Atoi(x)
		check(err)
	}
	partOne(nums)
	partTwo(nums)
}
