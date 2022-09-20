package easy

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	str := strconv.Itoa(x)
	x1 := 0
	y := len(str) - 1
	for x1 < y {
		if str[x1] != str[y] {
			return false
		}
		x1++
		y--
	}
	return true
}
