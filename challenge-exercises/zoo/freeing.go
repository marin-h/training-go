package main

import "fmt"

func main() {
	weights := [6]int{1, 5, 2, 7, 9, 3}
	capacity := 10
	fmt.Println(getAnimals(weights[:len(weights)], capacity))
}

func getAnimals(weights []int, capacity int) [2]int {
	max := 0
	min := 0
	for i := range weights {
		if weights[i] > weights[max] {
			max = i
		} else if weights[i] <= weights[min] {
			min = i
		}
		// check if max and min sum up to equal capacity
		if (i == len(weights)-1) && weights[max]+weights[min] == capacity {
			return [2]int{weights[max], weights[min]}
		}
	}
	// if no match until here, we need to lookup another match for our max value
	for i := range weights {
		if i != max && capacity-weights[max] == weights[i] {
			return [2]int{weights[max], weights[i]}
		}
	}
	// otherwise return empty truck
	return [2]int{}
}
