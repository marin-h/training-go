package main

import "fmt"

type Node struct {
	x, y     int
	children []*Node
}

func main() {
	root := Node{1, 1, []*Node{}}
	//end := Node{2, 5, []*Node{}} // YES
	end := Node{6, 3, []*Node{}} //	NO
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

	for _, next := range current.children {
		if match(*next, end) {
			*result = "YES"
			return
		} else if next.x <= end.x && next.y <= end.y {
			draw(next, end, result)
		}
	}
}

func match(some Node, goal Node) bool {
	return some.x == goal.x && some.y == goal.y
}
