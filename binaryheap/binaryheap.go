// Binhp is a slice-based implementation of a max binary heap data structure.
package binaryheap

import "errors"

type BinaryHeap []int

// Add an element to the heap.
func (h *BinaryHeap) Insert(n int) {
	if len(*h) == 0 {
		*h = BinaryHeap{n}
		return
	}
	*h = append(*h, n)
	(*h).percolateUp()
}

// Return the root of the heap.
func (h *BinaryHeap) Peek() int {
	return (*h)[0]
}

// Delete and return the root of the heap.
func (h *BinaryHeap) Extract() (int, error) {
	if len(*h) == 0 {
		return 0, errors.New("cannot extract from empty heap")
	}
	top := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	(*h).percolateDown(0)
	return top, nil
}

// Return whether n is present in the heap.
func (h *BinaryHeap) Contains(n int) bool {
	for _, v := range *h {
		if n == v {
			return true
		}
	}
	return false
}

// Delete the first occurrence of n from the heap.
// Return true if the deletion was successful.
func (h *BinaryHeap) Delete(n int) bool {
	for i, v := range *h {
		if v == n {
			(*h).swap(i, len(*h)-1)
			*h = (*h)[0 : len(*h)-1]
			(*h).percolateDown(i)
			return true
		}
	}
	return false
}

// Swap the elements at index n with the element at index m.
func (h *BinaryHeap) swap(n, m int) {
	tmp := (*h)[n]
	(*h)[n] = (*h)[m]
	(*h)[m] = tmp
}

// Restore the structure of the heap by cascading up the heap.
func (h *BinaryHeap) percolateUp() {
	var i, j int
	i = len(*h) - 1
	j = parent(i)
	for j >= 0 && (*h)[j] < (*h)[i] {
		h.swap(i, j)
		i = j
		j = parent(i)
	}
}

// Restore the structure of the heap by cascading down the heap.
func (h *BinaryHeap) percolateDown(p int) {
	l, r := children(p)
	if len(*h)-1 < l {
		return
	}
	if len(*h)-1 < r {
		if (*h)[p] < (*h)[l] {
			(*h).swap(p, l)
		}
		return
	}

	var mc int
	if (*h)[r] > (*h)[l] {
		mc = r
	} else {
		mc = l
	}
	if (*h)[p] < (*h)[mc] {
		(*h).swap(p, mc)
		(*h).percolateDown(mc)
	}
}

func children(p int) (int, int) {
	return 2*p + 1, 2*p + 2
}

func parent(i int) int {
	if i%2 == 0 {
		return (i - 2) / 2
	} else {
		return (i - 1) / 2
	}
}

func Heapify(arr []int) BinaryHeap {
	return BinaryHeap{}
}
