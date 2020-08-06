package main

import "fmt"

func main() {
	weights := [6]int{1, 5, 2, 7, 9, 3}
	capacity := 10
	fmt.Println(getAnimals(weights[:len(weights)], capacity))
}

func getAnimals(weights []int, capacity int) [2]int {
	for i := 0; i < len(weights)-1; i++ {
		for j := i + 1; j < len(weights); j++ {
			fmt.Println(i, j)
			if weights[i]+weights[j] == capacity {
				return [2]int{weights[i], weights[j]}
			}
		}
	}
	// otherwise return empty truck
	return [2]int{}
}
