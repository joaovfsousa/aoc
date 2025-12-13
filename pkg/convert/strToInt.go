package convert

import (
	"fmt"
	"strconv"
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert %v to int", s))
	}
	return i
}
