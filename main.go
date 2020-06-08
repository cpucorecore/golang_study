package main

import (
	"fmt"
	"golang_study/heap_sort"
)

func main() {
	a := []int{0, 4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heap_sort.BuildMaxHeap(a)
	fmt.Print(a)
}
