package main

import "fmt"

func main() {
	bookIDList := []int{2, 1, 4, 3, 15, 20, 27, 31}
	searchIdList := []int{2, 1, 4, 3, 15, 20, 27, 31}
	for _, id := range searchIdList {
		fmt.Println(findBookPosition(bookIDList, id))
	}
}

func findBookPosition(bookIDList []int, id int) int {

	left := 0
	right := len(bookIDList)
	position := (left + right) / 2

	for bookIDList[position] != id {
		if bookIDList[position] > id {
			if bookIDList[position+1] == id {
				return position + 1
			}
			right = position
		} else if bookIDList[position] < id {
			if bookIDList[position-1] == id {
				return position - 1
			}
			left = position
		}
		position = (left + right) / 2
	}
	return position
}
