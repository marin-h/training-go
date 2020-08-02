package main

import "fmt"

func main() {
	// Grokking Algorithms Exercise 4.1
	inputArray := []int{2, 4, 5, 0, 10}
	fmt.Println(sumRecursive(inputArray))
}

func sumRecursive(anArray []int) int {
	if len(anArray) == 0 {
		return 0
	} else {
		return anArray[0] + sumRecursive(anArray[1:len(anArray)])
	}
}
