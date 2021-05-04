package binaryheap

import (
	"reflect"
	"testing"
)

func TestBinaryHeapInsert(t *testing.T) {
	tests := []struct {
		name string
		proc func() BinaryHeap
		want BinaryHeap
	}{
		{
			name: "insert into empty heap",
			proc: func() BinaryHeap {
				h := BinaryHeap{}
				h.Insert(0)
				return h
			},
			want: BinaryHeap{0},
		},
		{
			name: "insert two elems in order",
			proc: func() BinaryHeap {
				h := BinaryHeap{}
				h.Insert(1)
				h.Insert(0)
				return h
			},
			want: BinaryHeap{1, 0},
		},
		{
			name: "insert two elems in reverse order",
			proc: func() BinaryHeap {
				h := BinaryHeap{}
				h.Insert(0)
				h.Insert(1)
				return h
			},
			want: BinaryHeap{1, 0},
		},
		{
			name: "insert 10 elements in random order",
			proc: func() BinaryHeap {
				h := BinaryHeap{}
				h.Insert(0)
				h.Insert(1)
				h.Insert(4)
				h.Insert(9)
				h.Insert(0)
				h.Insert(8)
				h.Insert(6)
				h.Insert(9)
				h.Insert(1)
				h.Insert(1)
				h.Insert(3)
				h.Insert(2)
				return h
			},
			want: BinaryHeap{9, 9, 8, 4, 3, 2, 6, 0, 1, 0, 1, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := test.proc()
			if !reflect.DeepEqual(h, test.want) {
				t.Errorf("testcase %s\n\tgot %v\n\twant %v", test.name, h, test.want)
			}
		})
	}
}

func TestBinaryHeapExtract(t *testing.T) {
	tests := []struct {
		name string
		proc func() []int
		want []int
	}{
		{
			name: "extract 0 from single root",
			proc: func() []int {
				h := BinaryHeap{}
				h.Insert(0)
				n, _ := h.Extract()
				return []int{n}
			},
			want: []int{0},
		},
		{
			name: "extract 1 from single root",
			proc: func() []int {
				h := BinaryHeap{}
				h.Insert(1)
				n, _ := h.Extract()
				return []int{n}
			},
			want: []int{1},
		},
		{
			name: "extract twice",
			proc: func() []int {
				h := BinaryHeap{}
				h.Insert(1)
				h.Insert(0)
				n, _ := h.Extract()
				retval := []int{n}
				n, _ = h.Extract()
				retval = append(retval, n)
				return retval
			},
			want: []int{1, 0},
		},
		{
			name: "extract four",
			proc: func() []int {
				h := BinaryHeap{}
				h.Insert(1)
				h.Insert(0)
				h.Insert(3)
				h.Insert(2)
				l := 4
				retval := make([]int, 0)
				for i := 0; i < l; i++ {
					n, _ := h.Extract()
					retval = append(retval, n)
				}
				return retval
			},
			want: []int{3, 2, 1, 0},
		},
		{
			name: "extract many",
			proc: func() []int {
				h := BinaryHeap{}
				h.Insert(1)
				h.Insert(0)
				h.Insert(3)
				h.Insert(2)
				h.Insert(1)
				h.Insert(0)
				h.Insert(3)
				h.Insert(2)
				h.Insert(1)
				h.Insert(0)
				h.Insert(3)
				h.Insert(2)
				l := 12
				retval := make([]int, 0)
				for i := 0; i < l; i++ {
					n, _ := h.Extract()
					retval = append(retval, n)
				}
				return retval
			},
			want: []int{3, 3, 3, 2, 2, 2, 1, 1, 1, 0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.proc()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("testcase %s\n\tgot  %d\n\twant %d", test.name, got, test.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name string
		proc func() BinaryHeap
		want BinaryHeap
	}{
		{
			name: "delete the root",
			proc: func() BinaryHeap {
				h := BinaryHeap{3}
				h.Delete(3)
				return h
			},
			want: BinaryHeap{},
		},
		{
			name: "delete a leaf",
			proc: func() BinaryHeap {
				h := BinaryHeap{9, 9, 8, 4, 3, 2, 6, 0, 1, 0, 1, 1}
				h.Delete(6)
				return h
			},
			want: BinaryHeap{9, 9, 8, 4, 3, 2, 1, 0, 1, 0, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.proc()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("testcase %s\n\tgot  %v\n\twant %v", test.name, got, test.want)
			}
		})
	}
}
