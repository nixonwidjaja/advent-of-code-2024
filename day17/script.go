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

func combo(A, B, C, operand int) int {
	if operand < 4 {
		return operand
	}
	switch operand {
	case 4:
		return A
	case 5:
		return B
	case 6:
		return C
	}
	return 0
}

func partOne(A, B, C int, nums []int) {
	outputs := make([]int, 0)
	for i := 0; i < len(nums); i += 2 {
		comboOperand := combo(A, B, C, nums[i+1])
		switch nums[i] {
		case 0:
			A = A >> comboOperand
		case 1:
			B = B ^ nums[i+1]
		case 2:
			B = comboOperand % 8
		case 3:
			if A == 0 {
				continue
			}
			i = nums[i+1] - 2
		case 4:
			B = B ^ C
		case 5:
			outputs = append(outputs, comboOperand%8)
		case 6:
			B = A >> comboOperand
		case 7:
			C = A >> comboOperand
		}
	}
	fmt.Println(join(outputs))
}

func join(outputs []int) string {
	outputStr := make([]string, len(outputs))
	for i, x := range outputs {
		outputStr[i] = strconv.Itoa(x)
	}
	return strings.Join(outputStr, ",")
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)
	dataStr := strings.Split(string(data), "\n")
	A, err := strconv.Atoi(strings.Fields(dataStr[0])[2])
	check(err)
	B, err := strconv.Atoi(strings.Fields(dataStr[1])[2])
	check(err)
	C, err := strconv.Atoi(strings.Fields(dataStr[2])[2])
	check(err)
	numsStr := strings.Split(strings.Fields(dataStr[4])[1], ",")
	nums := make([]int, len(numsStr))
	for i, x := range numsStr {
		nums[i], err = strconv.Atoi(x)
		check(err)
	}
	partOne(A, B, C, nums)
}
