package medium

import (
	"sort"
)

func Change(change uint16, coins []uint16) []uint16 {
	result := [][]uint16{}

	// to get maximum first for least coins to use
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	for o, oc := range coins {
		if oc > change {
			continue
		}

		for i := o + 1; i < len(coins); i++ {
			ic := coins[i]
			if ic > change {
				continue
			}

			if oc+ic == change {
				combination := []uint16{oc, ic}
				result = append(result, combination)
				break
			}

			for j := i + 1; j < len(coins); j++ {
				jc := coins[j]
				if jc > change {
					continue
				}

				if oc+ic+jc == change {
					combination := []uint16{oc, ic, jc}
					result = append(result, combination)
					break
				}
			}
		}
	}

	sort.Slice(result[0], func(i, j int) bool {
		return result[0][i] < result[0][j]
	})

	return result[0]
}
