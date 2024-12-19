package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	dataStr := string(data)
	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	all := re.FindAllString(dataStr, -1)
	ans := 0
	for _, s := range all {
		comma := strings.Index(s, ",")
		a, err := strconv.Atoi(s[4:comma])
		check(err)
		b, err := strconv.Atoi(s[comma+1 : len(s)-1])
		check(err)
		ans += a * b
	}
	fmt.Println(ans)

	allIndex := re.FindAllStringIndex(dataStr, -1)
	doRe := regexp.MustCompile(`do\(\)`)
	allDo := doRe.FindAllStringIndex(dataStr, -1)
	dontRe := regexp.MustCompile(`don't\(\)`)
	allDont := dontRe.FindAllStringIndex(dataStr, -1)

	events := make([][]int, 0)
	for _, x := range allIndex {
		events = append(events, []int{x[0], 0})
	}
	for _, x := range allDo {
		events = append(events, []int{x[0], 1})
	}
	for _, x := range allDont {
		events = append(events, []int{x[0], -1})
	}
	sort.Slice(events, func(a, b int) bool {
		return events[a][0] < events[b][0]
	})
	i, enabled, ansEnabled := 0, true, 0
	for _, x := range events {
		switch x[1] {
		case 0:
			if enabled {
				s := all[i]
				comma := strings.Index(s, ",")
				a, err := strconv.Atoi(s[4:comma])
				check(err)
				b, err := strconv.Atoi(s[comma+1 : len(s)-1])
				check(err)
				ansEnabled += a * b
			}
			i++
		case 1:
			enabled = true
		case -1:
			enabled = false
		}
	}
	fmt.Println(ansEnabled)
}
