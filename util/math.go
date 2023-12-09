package util

import "strconv"

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func IsInt(c rune) bool {
	return c >= 48 && c <= 57
}
