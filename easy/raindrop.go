package easy

import (
	"strconv"
)

func Raindrop(num *int, res *string) string {
	isFactor := false

	if (*num % 3) == 0 {
		isFactor = true
		*res += "Pling"
	}

	if (*num % 5) == 0 {
		isFactor = true
		*res += "Plang"
	}

	if (*num % 7) == 0 {
		isFactor = true
		*res += "Plong"
	}

	if !isFactor {
		*res = strconv.Itoa(*num)
	}

	return ""
}
