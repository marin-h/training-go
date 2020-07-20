package main

import (
	"fmt"
	"math"
)

func mainKinderSquare() {
	arr := []int{3, 1, 4}
	kinderSquare(arr)
}

func kinderSquare(arr []int) {
	for _, value := range arr {
		if condition(value) {
			fmt.Println(calculate(value))
		}
	}
}

func calculate(value int) int {
	return int(math.Pow(float64(value), 2) + 10)
}

func condition(value int) bool {
	remain := calculate(value) % 10
	return remain != 5 && remain != 6
}
