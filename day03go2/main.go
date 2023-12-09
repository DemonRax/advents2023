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
	gears := findGears(ss)
	for i, s := range ss {
		nums := scanLine(s)
		for _, coords := range nums {
			xs, ys := gearsTouching(gears, coords, i)
			for k := range xs {
				gears[xs[k]][ys[k]] = append(gears[xs[k]][ys[k]], convertNum(s, coords))
			}
		}
	}
	var sum int
	for _, gr := range gears {
		for _, g := range gr {
			if len(g) == 2 {
				sum += g[0] * g[1]
			}
		}
	}
	return sum
}

func convertNum(s string, coords []int) int {
	return util.ToInt(s[coords[0] : coords[len(coords)-1]+1])
}

func gearsTouching(gears [][][]int, coords []int, ii int) ([]int, []int) {
	xs := make([]int, 0, len(coords))
	ys := make([]int, 0, len(coords))
	i1 := ii - 1
	i2 := ii + 1
	if i1 < 0 {
		i1 = 0
	}
	if i2 >= len(gears) {
		i2 = len(gears) - 1
	}
	j1 := coords[0] - 1
	j2 := coords[len(coords)-1] + 1
	if j1 < 0 {
		j1 = 0
	}
	if j2 >= len(gears[0]) {
		j2 = len(gears[0]) - 1
	}
	for i := i1; i <= i2; i++ {
		for j := j1; j <= j2; j++ {
			if gears[i][j] != nil {
				xs = append(xs, i)
				ys = append(ys, j)
			}
		}
	}
	return xs, ys
}

func scanLine(s string) [][]int {
	res := make([][]int, 0, len(s)/2)
	coords := make([]int, 0, len(s))
	for j, c := range s {
		if util.IsInt(c) {
			coords = append(coords, j)
			continue
		} else {
			if len(coords) > 0 {
				res = append(res, coords)
				coords = []int{}
			}
		}
	}
	if len(coords) > 0 {
		res = append(res, coords)
	}
	return res
}

const gear = '*'

func findGears(ss []string) [][][]int {
	res := make([][][]int, len(ss))
	for i, s := range ss {
		res[i] = make([][]int, len(s))
		for j, c := range s {
			if c == gear {
				res[i][j] = []int{}
			}
		}
	}
	return res
}
