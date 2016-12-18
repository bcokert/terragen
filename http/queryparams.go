package http

import (
	"strconv"
	"strings"
)

// ParseIntArrayParam tries to parse the given query param into an array of integers
// If it can't, it always returns an empty list
func ParseIntArray(v string) []int {
	ints := make([]int, 0, 3)

	if v == "" {
		return []int{}
	}

	for _, value := range strings.Split(v, ",") {
		num, err := strconv.Atoi(value)
		if err == nil {
			ints = append(ints, num)
		} else {
			return []int{}
		}
	}

	return ints
}
