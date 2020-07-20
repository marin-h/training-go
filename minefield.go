package main

import "fmt"

type Node struct {
	x, y     int
	children []*Node
}

func main() {

	root := Node{1, 1, []*Node{}}

	goals := [2]Node{{2, 5, []*Node{}}, Node{6, 3, []*Node{}}}

	for _, end := range goals {
		result := "NO"
		draw(&root, end, &result)
		fmt.Println(result)
	}
}

func draw(current *Node, end Node, result *string) {

	current.children = append(current.children, &Node{current.x + current.y, current.y, []*Node{}})
	current.children = append(current.children, &Node{current.x, current.y + current.x, []*Node{}})

	for _, next := range current.children {
		if match(*next, end) {
			*result = "YES"
			return
		} else if near(*next, end) {
			draw(next, end, result)
		}
	}
}

func match(some Node, goal Node) bool {
	return some.x == goal.x && some.y == goal.y
}

func near(next Node, end Node) bool {
	// check if end node is still in current path, since we only move forward
	// check next node not overflowing 500 position for x and y
	return next.x <= end.x && next.y <= end.y && next.x <= 500 && next.y <= 500
}
