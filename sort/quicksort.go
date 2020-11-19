package sort

import "math/rand"

func QuickSort(input []QSData, p, r int) []QSData {
	if p < r {
		q := partition(input, p, r)
		QuickSort(input, p, q)
		QuickSort(input, q+1, r)
	}
	return input
}

func partition(input []QSData, p, r int) int {
	randIdx := p + (rand.Int() % (r - p + 1))
	input[p], input[randIdx] = input[randIdx], input[p]
	bearingItem := input[p]
	i := p - 1
	j := r + 1
	for {
		for {
			j--
			if input[j].LessThan(bearingItem) || input[j].EqualsTo(bearingItem) {
				break
			}
		}
		for {
			i++
			if bearingItem.LessThan(input[i]) || bearingItem.EqualsTo(input[i]) {
				break
			}
		}
		if i < j {
			input[i], input[j] = input[j], input[i]
		} else {
			return j
		}
	}
}
