package easy

import (
	"fmt"
	"math"
	"strconv"
)

func LargestSeriesProduct(series string, span int) int {
	largest := math.MinInt

	for i := range series {
		product := 1
		if i+(span-1) < len(series) {
			wordSpan := series[i : i+span]
			for wi := 0; wi < len(wordSpan); wi++ {
				num, err := strconv.Atoi(string(wordSpan[wi]))
				if err == nil {
					product *= num
				} else {
					fmt.Print("didnt able to convert to number")
				}
			}
			if product > largest {
				largest = product
			}
		}
	}

	return largest
}
