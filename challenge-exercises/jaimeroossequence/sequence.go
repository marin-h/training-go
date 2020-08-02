package main

import "fmt"

func main() {
	j := 6
	jStartSequence := []int{1, 1, 1}

	fmt.Println(sequence(j, jStartSequence))
}

func sequence(jGoal int, jSequence []int) int {

	var jn int

	// use last three numbers to calculate Jn
	tailSequence := jSequence[len(jSequence)-3:]
	for _, number := range tailSequence {
		jn = jn + number
	}
	jSequence = append(jSequence, jn)

	// base case
	if len(jSequence) > jGoal {
		return jSequence[jGoal]
	}
	// recursive case
	return sequence(jGoal, jSequence)
}
