package main

import (
	"fmt"

	"advents2023/util"
)

func main() {
	input := util.ReadFile("./day03go1/input.txt")
	fmt.Println(partNumbers(input))
}

func partNumbers(ss []string) int {
	src := convertInput(ss)
	var sum int
	for i := range src {
		sum += scanLine(src, i)
	}
	return sum
}

func scanLine(src [][]int, i int) int {
	var sum, num int
	var touch bool
	for j, c := range src[i] {
		if c >= 0 {
			num = num*10 + c
			touch = touch || isTouching(src, i, j)
			continue
		}
		if touch {
			sum += num
		}
		num = 0
		touch = false
	}
	if touch {
		sum += num
	}
	return sum
}

func isTouching(src [][]int, i int, j int) bool {
	i1 := i - 1
	i2 := i + 1
	if i1 < 0 {
		i1 = 0
	}
	if i2 >= len(src) {
		i2 = len(src) - 1
	}
	j1 := j - 1
	j2 := j + 1
	if j1 < 0 {
		j1 = 0
	}
	if j2 >= len(src[0]) {
		j2 = len(src[0]) - 1
	}

	for i = i1; i <= i2; i++ {
		for j = j1; j <= j2; j++ {
			if src[i][j] == -2 {
				return true
			}
		}
	}
	return false
}

const dot = '.'

func convertInput(ss []string) [][]int {
	res := make([][]int, len(ss))
	for i, s := range ss {
		res[i] = make([]int, len(s))
		for j, c := range s {
			if c == dot {
				res[i][j] = -1
			} else if util.IsInt(c) {
				res[i][j] = util.ToInt(string(c))
			} else {
				res[i][j] = -2
			}
		}
	}
	return res
}
