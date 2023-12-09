package main

import (
	"fmt"
	"strings"

	"advents2023/util"
)

func main() {
	input := util.ReadFile("./day04go1/input.txt")
	fmt.Println(scratchCards(input))
}

func scratchCards(ss []string) int {
	var sum int
	copies := make([]int, len(ss))
	for i, s := range ss {
		wins, nums := scanLine(s)
		var matches int
		for _, win := range wins {
			for _, num := range nums {
				if win == num {
					matches++
				}
			}
		}
		sum += copies[i] + 1
		if matches > 0 {
			for j := 0; j < matches; j++ {
				copies[i+j+1] = copies[i+j+1] + copies[i] + 1
			}
		}
		matches = 0
	}
	return sum
}

func scanLine(s string) ([]int, []int) {
	prefix := strings.Split(s, ":")
	all := strings.Split(strings.TrimSpace(prefix[1]), " | ")
	return convertInts(strings.Split(all[0], " ")), convertInts(strings.Split(all[1], " "))
}

func convertInts(in []string) []int {
	res := make([]int, 0, len(in))
	for _, s := range in {
		if s == "" {
			continue
		}
		res = append(res, util.ToInt(s))
	}
	return res
}
