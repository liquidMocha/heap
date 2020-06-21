package main

type Heap struct {
	data []int
}

func (heap Heap) parent(i int) int {
	return (i - 1) / 2
}

func (heap Heap) left(i int) int {
	return i*2 + 1
}

func (heap Heap) right(i int) int {
	return (i + 1) * 2
}

func (heap Heap) hasLeft(i int) bool {
	return len(heap.data) > (i*2 + 1)
}

func (heap Heap) hasRight(i int) bool {
	return len(heap.data) > (i+1)*2
}

func (heap Heap) maxHeapify(index int) {
	indexOfLargest := index

	if heap.hasLeft(index) && heap.data[heap.left(index)] > heap.data[indexOfLargest] {
		indexOfLargest = heap.left(index)
	}

	if heap.hasRight(index) && heap.data[heap.right(index)] > heap.data[indexOfLargest] {
		indexOfLargest = heap.right(index)
	}

	if indexOfLargest != index {
		temp := heap.data[index]
		heap.data[index] = heap.data[indexOfLargest]
		heap.data[indexOfLargest] = temp
		heap.maxHeapify(indexOfLargest)
	}
}

func (heap Heap) buildMaxHeap() {
	for i := (len(heap.data) - 1 - 1) / 2; i >= 0; i-- {
		heap.maxHeapify(i)
	}
}

func (heap Heap) sort() []int {
	var result []int

	for len(heap.data) != 0 {
		heap.buildMaxHeap()

		result = append(result, heap.data[0])
		heap.data = heap.data[1:]
	}

	return result
}
