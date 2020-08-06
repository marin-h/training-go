package main

import "fmt"

func main() {
	graph := map[string][]string{}
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	fmt.Println(bfs(graph))
}

func bfs(graph map[string][]string) string {

	visited := map[string]bool{}
	queue := append([]string{}, graph["you"]...)

	for len(queue) > 0 {

		person := queue[0]
		queue = queue[1:]

		if visited[person] {
			fmt.Println(person, "already visited")
			continue
		}
		if matches(person) {
			return "found " + person
		} else {
			fmt.Println("visiting", person, "-> adding their friends to queue")
			visited[person] = true
			queue = append(queue, graph[person]...)
		}
	}
	return "not found"
}

func matches(name string) bool {
	return name[0:1] == "t"
}
