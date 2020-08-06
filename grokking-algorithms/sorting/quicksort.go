package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := generateSlice(10)
	fmt.Println(qsort(arr))
}

func qsort(a []int) []int {
	// base case: array is 1 or 0 elem
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// choose a pivot randomly
	pivotIndex := rand.Int() % len(a)

	// swap the pivot with last value at the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// if element is smaller than pivot move to left side
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	// move the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// shoot recursive calls
	qsort(a[:left])
	qsort(a[left+1:])

	return a
}

// generate random values for slice
func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
