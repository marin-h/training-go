package main

import "fmt"

func main() {
	bookIDList := []int{2, 1, 4, 3, 15, 20, 27, 31}
	fmt.Println(findBookPosition(bookIDList, 25))
}

func findBookPosition(bookIDList []int, id int) int {

	left := 0
	right := len(bookIDList)
	position := (left + right) / 2

	for len(bookIDList[left:right]) > 1 {
		position = (left + right) / 2
		if bookIDList[position] == id {
			return position
		}
		if bookIDList[position] > id {
			if bookIDList[position+1] == id {
				position = position + 1
			}
			right = position
		} else if bookIDList[position] < id {
			if bookIDList[position-1] == id {
				position = position - 1
			}
			left = position
		}
	}
	return -1
}
