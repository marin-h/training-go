package main

import (
	"fmt"
)

func main() {
	// Grokking Algorithms Exercise 4.4
	anArray := []float64{4, 8, 9, 15, 66, 73, 102}
	var elem float64 = 8

	if position := binarySearch(0, len(anArray), anArray, elem); position >= 0 {
		fmt.Println("number position is", position)
	} else {
		fmt.Println("number not found")
	}
}

func binarySearch(start int, end int, anArray []float64, elem float64) int {

	position := (start + end) / 2

	// base case: list is reduced to 1 elem
	if len(anArray[start:end]) == 1 {
		if anArray[position] == elem {
			return position
		} else {
			return -1
		}
	} else {
		// recursive case: list is still searchable
		if anArray[position] > elem {
			end = position - 1
		} else if anArray[position] < elem {
			start = position + 1
		} else {
			start = position
			end = position + 1
		}
		return binarySearch(start, end, anArray, elem)
	}
}
