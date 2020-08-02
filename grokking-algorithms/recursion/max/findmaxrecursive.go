package main

import (
	"fmt"
	"math"
)

func main() {
	// Grokking Algorithms Exercise 4.3
	anArray := []float64{3, 314, 41, 32}
	fmt.Println("max number is", findMax(anArray, 0))
}

func findMax(anArray []float64, max float64) float64 {

	if len(anArray) == 0 {
		return max
	} else {
		return findMax(anArray[1:], math.Max(anArray[0], max))
	}
}
