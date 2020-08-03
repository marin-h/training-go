package main

import "fmt"

func main() {

	var stacks = [][]int{{}} // initialize first stack
	var limit = 2

	add := func(id int) {
		stack := stacks[len(stacks)-1]
		stack = append(stack, id)
		stacks[len(stacks)-1] = stack
		if len(stack) == limit {
			aStack := []int{}
			stacks = append(stacks, aStack)
		}
	}

	remove := func() int {
		removed := -1
		if len(stacks) >= 1 && len(stacks[0]) >= 1 {
			// remove last element of first stack
			removed = stacks[0][len(stacks[0])-1]
			stacks[0] = stacks[0][:len(stacks[0])-1]
			// organize dishes, until all fillable stacks are full
			for i := 0; i < len(stacks)-1; i++ {
				if len(stacks[i+1]) > 0 {
					stacks[i] = append(stacks[i], stacks[i+1][len(stacks[i+1])-1])
					stacks[i+1] = stacks[i+1][:len(stacks[i+1])-1]
				}
			}
			// if last stack got empty, delete it
			if len(stacks[len(stacks)-1]) == 0 {
				stacks = stacks[:len(stacks)-1]
			}
		}
		return removed
	}
	count := func(stackIndex int) int {
		if len(stacks) > stackIndex {
			return len(stacks[stackIndex])
		}
		return -1
	}

	// test the inputs, output should be 2, 1, 2, 7, 2, 2, -1
	add(13)
	add(7)
	add(17)
	fmt.Println(count(0))
	fmt.Println(count(1))
	add(2)
	fmt.Println(count(1))
	add(47)
	fmt.Println(remove())
	fmt.Println(count(0))
	fmt.Println(count(1))
	fmt.Println(count(2))
}
