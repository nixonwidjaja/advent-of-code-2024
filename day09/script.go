package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partOne(arr []int) {
	n := len(arr)
	L, R := 0, n-1
	for L < R {
		for L < n && arr[L] != -1 {
			L++
		}
		for R >= 0 && arr[R] == -1 {
			R--
		}
		if L >= R {
			break
		}
		arr[L] = arr[R]
		arr[R] = -1
		L++
		R--
	}
	ans := 0
	for i, x := range arr {
		if x == -1 {
			break
		}
		ans += i * x
	}
	R = n - 1
	for arr[R] == -1 {
		R--
	}
	fmt.Println(ans)
}

func partTwo(arr []int, data string) {
	free := make([][]int, 0) // (idx, size)
	mem := make([]int, 0)
	idx := 0
	for i, x := range data {
		num := int(x - '0')
		if i%2 == 1 && num > 0 {
			free = append(free, []int{idx, num})
		}
		mem = append(mem, idx)
		idx += num
	}
	for i := len(data) - 1; i >= 0; i -= 2 {
		index, size := mem[i], int(data[i]-'0')
		for j, x := range free {
			freeIdx, freeSize := x[0], x[1]
			if freeIdx >= index {
				break
			}
			if freeSize == size {
				free[j][1] = 0
			} else if freeSize > size {
				free[j] = []int{freeIdx + size, freeSize - size}
			} else {
				continue
			}
			for k := freeIdx; k < freeIdx+size; k++ {
				arr[k] = arr[index]
			}
			for k := index; k < index+size; k++ {
				arr[k] = -1
			}
			break
		}
	}
	ans := 0
	for i, x := range arr {
		if x == -1 {
			continue
		}
		ans += i * x
	}
	fmt.Println(ans)
}

func main() {
	file, err := os.ReadFile("./input")
	check(err)
	data := string(file)
	id := 0
	arr := make([]int, 0)
	for i, x := range data {
		num := int(x - '0')
		for j := 0; j < num; j++ {
			if i%2 == 0 {
				arr = append(arr, id)
			} else {
				arr = append(arr, -1)
			}
		}
		if i%2 == 0 {
			id++
		}
	}
	arrOne := make([]int, len(arr))
	copy(arrOne, arr)
	partOne(arrOne)
	partTwo(arr, data)
}
