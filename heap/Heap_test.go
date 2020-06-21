package heap

import (
	"testing"
	"time"
)
import "github.com/stretchr/testify/assert"
import "math/rand"

func TestHeap(t *testing.T) {
	t.Run("should get parent", func(t *testing.T) {
		heap := Heap{data: []int{1, 2, 3, 4, 5, 6, 7}}
		assert.Equal(t, 1, heap.data[heap.parent(1)])
		assert.Equal(t, 1, heap.data[heap.parent(2)])
		assert.Equal(t, 2, heap.data[heap.parent(3)])
		assert.Equal(t, 2, heap.data[heap.parent(4)])
		assert.Equal(t, 3, heap.data[heap.parent(5)])
		assert.Equal(t, 3, heap.data[heap.parent(6)])
	})

	t.Run("should get left child", func(t *testing.T) {
		heap := Heap{data: []int{1, 2, 3, 4, 5, 6, 7, 8}}
		assert.Equal(t, 2, heap.data[heap.left(0)])
		assert.Equal(t, 4, heap.data[heap.left(1)])
		assert.Equal(t, 6, heap.data[heap.left(2)])
	})

	t.Run("should get right child", func(t *testing.T) {
		heap := Heap{data: []int{1, 2, 3, 4, 5, 6, 7, 8}}
		actual := heap.right(2)

		assert.Equal(t, 7, heap.data[actual])
	})

	t.Run("should max heapify", func(t *testing.T) {
		heap := Heap{data: []int{0, 8, 5, 4, 2, 1, 3}}
		heap.maxHeapify(0)

		expected := Heap{[]int{8, 4, 5, 0, 2, 1, 3}}
		assert.Equal(t, expected, heap)
	})

	t.Run("should build max heap", func(t *testing.T) {
		heap := Heap{}
		rand.Seed(time.Now().UnixNano())

		for i := 0; i < 50; i++ {
			randomNumber := rand.Intn(100)
			heap.data = append(heap.data, randomNumber)
		}

		heap.buildMaxHeap()

		for i := range heap.data {
			if heap.hasLeft(i) {
				assert.True(t, heap.data[heap.left(i)] <= heap.data[i])
			}
			if heap.hasRight(i) {
				assert.True(t, heap.data[heap.right(i)] <= heap.data[i])
			}
		}
	})

	t.Run("should heap sort", func(t *testing.T) {
		heap := Heap{}
		rand.Seed(time.Now().UnixNano())

		for i := 0; i < 50; i++ {
			randomNumber := rand.Intn(100)
			heap.data = append(heap.data, randomNumber)
		}

		actual := heap.sort()

		assert.True(t, len(actual) == len(heap.data))
		for i := range actual {
			if i != len(actual)-1 {
				assert.True(t, actual[i] >= actual[i+1])
			}
		}
	})
}
