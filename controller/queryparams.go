package controller

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// ParseFloatArrayParam tries to parse the given query param into an array of floats
// It lets the user handle the case where it's empty/not present
func ParseFloatArrayParam(queryParams url.Values, key string) ([]float64, error) {
	values := queryParams.Get(key)
	floats := make([]float64, 0, 3)

	if values == "" {
		return floats, nil
	}

	for _, value := range strings.Split(values, ",") {
		num, err := strconv.ParseFloat(value, 64)
		if err == nil {
			floats = append(floats, num)
		} else {
			return []float64{}, fmt.Errorf("GetFloatArrayParam: failed to parse a float '%s' in the array: (%s)", value, err.Error())
		}
	}

	return floats, nil
}
