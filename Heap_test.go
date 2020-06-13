package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestHeap(t *testing.T) {
	t.Run("should get parent", func(t *testing.T) {
		heap := Heap{data: []int{1, 2, 3, 4, 5, 6, 7}}
		assert.Equal(t, 1, heap.parent(1))
		assert.Equal(t, 1, heap.parent(2))
		assert.Equal(t, 2, heap.parent(3))
		assert.Equal(t, 2, heap.parent(4))
		assert.Equal(t, 3, heap.parent(5))
		assert.Equal(t, 3, heap.parent(6))
	})

	t.Run("should get left child", func(t *testing.T) {
		heap := Heap{data: []int{1, 2, 3, 4, 5, 6, 7, 8}}
		assert.Equal(t, 2, heap.left(0))
		assert.Equal(t, 4, heap.left(1))
		assert.Equal(t, 6, heap.left(2))
	})

	t.Run("should get right child", func(t *testing.T) {
		heap := Heap{data: []int{1, 2, 3, 4, 5, 6, 7, 8}}
		actual := heap.right(2)

		assert.Equal(t, 7, actual)
	})

	t.Run("should max heapify", func(t *testing.T) {
		heap := Heap{data: []int{0, 8, 5, 4, 2, 1, 3}}
		actual := heap.maxHeapify(0)

		expected := Heap{[]int{8, 4, 5, 0, 2, 1, 3}}
		assert.Equal(t, expected, actual)
	})
}
