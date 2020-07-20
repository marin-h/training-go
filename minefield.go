package main

import (
	"fmt"
)

type Node struct {
	x, y     int
	children []*Node
}

func main() {
	root := Node{1, 1, []*Node{}}
	end := Node{2, 5, []*Node{}}
	draw(&root, end)
}

func draw(current *Node, end Node) {

	if current.x >= 500 || current.y >= 500 {
		//fmt.Println("NO")
		return
	}

	current.children = append(current.children, &Node{current.x + current.y, current.y, []*Node{}})
	current.children = append(current.children, &Node{current.x, current.y + current.x, []*Node{}})

	if match(*current.children[0], end) || match(*current.children[1], end) {
		fmt.Println("YES")
	} else {
		for _, next := range current.children {
			draw(next, end)
		}
	}
}

func match(some Node, goal Node) bool {
	return some.x == goal.x && some.y == goal.y
}
