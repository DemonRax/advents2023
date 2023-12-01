package main

import (
	"fmt"
	"strconv"

	"advents2023/util"
)

func main() {
	input := util.ReadFile("./day01go1/input.txt")
	fmt.Println(calibrate(input))
}

func calibrate(ss []string) int {
	var first, last, sum int
	for _, s := range ss {
		for _, c := range s {
			d, err := strconv.Atoi(string(c))
			if err != nil {
				continue
			}
			last = d
			if first == 0 {
				first = d
			}
		}
		sum += first*10 + last
		first = 0
	}
	return sum
}
