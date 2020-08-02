package main

import "fmt"

func main() {
	// Grokking Algorithms Exercise 4.2
	anArray := []int{3, 314, 32}
	fmt.Println("list has", countList(anArray, 0), "elements")
}

func countList(anArray []int, count int) int {

	if len(anArray) == 0 {
		return count
	} else {
		count += 1
		return countList(anArray[1:], count)
	}
}
