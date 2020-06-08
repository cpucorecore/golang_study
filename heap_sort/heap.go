package heap_sort

func LEFT(i int) int {
	return 2 * i
}

func RIGHT(i int) int {
	return 2*i + 1
}

func MaxHeapify(a []int, i int) {
	largest := i

	l := LEFT(i)
	r := RIGHT(i)

	if l <= (len(a)-1) && a[l] > a[i] {
		largest = l
	}

	if r <= (len(a)-1) && a[r] > a[largest] {
		largest = r
	}

	if largest != i {
		a[largest], a[i] = a[i], a[largest]
		MaxHeapify(a, largest)
	}
}

func BuildMaxHeap(a []int) {
	size := len(a) - 1
	for i := size / 2; i >= 1; i-- {
		MaxHeapify(a, i)
	}
}
