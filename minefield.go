package main

import "fmt"

type Node struct {
	x, y     int
	children []*Node
}

func main() {
	root := Node{1, 1, []*Node{}}
	end := Node{2, 5, []*Node{}} // NO
	//	end := Node{6, 3, []*Node{}}	YES
	result := "NO"
	draw(&root, end, &result)
	fmt.Println(result)
}

func draw(current *Node, end Node, result *string) {

	if current.x >= 500 || current.y >= 500 {
		return
	}

	current.children = append(current.children, &Node{current.x + current.y, current.y, []*Node{}})
	current.children = append(current.children, &Node{current.x, current.y + current.x, []*Node{}})

	if match(*current.children[0], end) || match(*current.children[1], end) {
		*result = "YES"
	} else {
		for _, next := range current.children {
			draw(next, end, result)
		}
	}
}

func match(some Node, goal Node) bool {
	return some.x == goal.x && some.y == goal.y
}
