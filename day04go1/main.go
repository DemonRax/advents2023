package main

import (
	"fmt"
	"strings"

	"advents2023/util"
)

func main() {
	input := util.ReadFile("./day04go1/input.txt")
	fmt.Println(winningPoints(input))
}

func winningPoints(ss []string) int {
	var sum int
	for _, s := range ss {
		wins, nums := scanLine(s)
		var matches int
		for _, win := range wins {
			for _, num := range nums {
				if win == num {
					matches++
				}
			}
		}
		if matches > 0 {
			sum += util.Pow(2, matches-1)
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
