package main

import (
	"fmt"
	"strconv"
	"strings"

	"advents2023/util"
)

func main() {
	input := util.ReadFile("./day01go2/input.txt")
	fmt.Println(calibrate(input))
}

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func calibrate(ss []string) int {
	var sum int
	for _, s := range ss {
		var nums []int
		for len(s) > 0 {
			d, err := strconv.Atoi(string(s[0]))
			if err == nil {
				nums = append(nums, d)
				s = s[1:]
				continue
			}
			for k, v := range digits {
				if strings.HasPrefix(s, k) {
					nums = append(nums, v)
					s = s[len(k)-2:]
				}
			}
			s = s[1:]
		}
		sum += nums[0]*10 + nums[len(nums)-1]
	}
	return sum
}
