package main

import "fmt"

func main() {
	flightLength := 311
	movieLengths := []int{50, 300, 121, 190, 110}
	fmt.Println(getMovies(flightLength, movieLengths))
}

func getMovies(flightLength int, movieLengths []int) bool {

	// find if two movies together sum exactly to flightlength in O(n) runtime
	movieMap := map[int]int{}
	for i, length := range movieLengths {
		if length < flightLength {
			movieMap[length] = i
		}
	}
	for length, _ := range movieMap {
		for i := 0; i <= 20; i++ {
			if _, matches := movieMap[flightLength-length-i]; matches {
				return true
			}
		}
	}
	return false
}
