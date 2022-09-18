package main

import (
	"fmt"
	"math"
)

// Kadane's Algo
// This algorithm can be used to find largest sum of subarray.
// The idea is to look for positive contiguous subarray.
// As soon as the sum hits negative, reject it.
// Time complexity O(n)
// Space complexity O(1)
func kadane(array []int) (int, []int) {
	subarray := []int{}
	maxValue := math.MinInt
	currentMax := 0
	for _, val := range array {
		if currentMax < 0 {
			currentMax = 0
			subarray = []int{}
		}
		currentMax = currentMax + val
		subarray = append(subarray, val)
		if maxValue < currentMax {
			maxValue = currentMax
		}

	}
	if maxValue < 0 {
		subarray = []int{maxValue}
	}
	return maxValue, subarray
}

func main() {
	array := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	sum, subarray := kadane(array)
	fmt.Println(sum, subarray)

	array = []int{-3, -2, -4}
	sum, subarray = kadane(array)
	fmt.Println(sum, subarray)
}
