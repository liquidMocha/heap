package main

type Heap struct {
	data []int
}

func (heap Heap) parent(i int) *int {
	return &heap.data[(i-1)/2]
}

func (heap Heap) left(i int) *int {
	return &heap.data[i*2+1]
}

func (heap Heap) right(i int) *int {
	return &heap.data[(i+1)*2]
}

func (heap Heap) hasLeft(i int) bool {
	return len(heap.data) > (i*2 + 1)
}

func (heap Heap) hasRight(i int) bool {
	return len(heap.data) > (i+1)*2
}

func (heap Heap) maxHeapify(index int) Heap {
	current := heap.data[index]
	if heap.hasLeft(index) {
		left := *heap.left(index)

		if heap.hasRight(index) {
			right := *heap.right(index)

			if left > current && left > right {
				heap.data[index] = left
				*heap.left(index) = current

				return heap.maxHeapify(index*2 + 1)
			}

			if right > current && right > left {
				heap.data[index] = right
				*heap.right(index) = current

				return heap.maxHeapify((index + 1) * 2)
			}
		} else {
			if left > current {
				heap.data[index] = left
				*heap.left(index) = current

				return heap.maxHeapify(index*2 + 1)
			}
		}
	}

	return heap
}
