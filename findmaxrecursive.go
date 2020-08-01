package main

import (
	"fmt"
	"math"
)

func main() {
	// Grokking Algorithms Exercise 4.3
	anArray := []float64{3, 41, 314, 32}
	fmt.Println("max number is", findMax(anArray, 0))
}

func findMax(anArray []float64, max float64) float64 {

	if len(anArray) == 0 {
		return max
	} else {
		return findMax(anArray[2:len(anArray)], math.Max(anArray[0], anArray[1]))
	}
}
