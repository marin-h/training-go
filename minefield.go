package main

import "fmt"

type Node struct {
	x, y     int
	children []*Node
}

type Path struct {
	root, end Node
}

func mainMainfield() {

	root := Node{1, 1, []*Node{}}
	goal1 := Node{2, 5, []*Node{}}
	goal2 := Node{6, 3, []*Node{}}
	paths := [2]Path{{root, goal1}, {root, goal2}}

	for _, path := range paths {
		if draw(&path.root, path.end) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func draw(current *Node, end Node) bool {

	current.children = append(current.children, &Node{current.x + current.y, current.y, []*Node{}})
	current.children = append(current.children, &Node{current.x, current.y + current.x, []*Node{}})
	for _, next := range current.children {
		if match(*next, end) {
			return true
		} else if near(*next, end) {
			return draw(next, end)
		}
	}
	return false
}

func match(some Node, goal Node) bool {
	return some.x == goal.x && some.y == goal.y
}

func near(next Node, end Node) bool {
	// check if end node is still in current path, since we only move forward
	// check next node not overflowing 500 position for x and y
	return next.x <= end.x && next.y <= end.y && next.x <= 500 && next.y <= 500
}
