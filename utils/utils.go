package utils

import (
	"os"
	"strings"
)

type Position struct {
	X, Y int
}

func ParseGrid(filePath string) ([]string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(file), "\n"), nil
}

func Inbounds(x, y int, rows, cols int) bool {
	return x >= 0 && x < cols && y >= 0 && y < rows
}

type Heap[T any] struct {
	data     []T
	lessFunc func(a, b T) bool
}

type UnionFind[T comparable] struct {
	parent map[T]T
	rank   map[T]int
}

func NewUnionFind[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
	}
}

func (uf *UnionFind[T]) Find(x T) T {
	if _, exists := uf.parent[x]; !exists {
		uf.parent[x] = x
		uf.rank[x] = 1
	}

	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind[T]) Union(x, y T) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

func (uf *UnionFind[T]) Connected(x, y T) bool {
	return uf.Find(x) == uf.Find(y)
}

func NewHeap[T any](lessFunc func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data:     []T{},
		lessFunc: lessFunc,
	}
}

func (h *Heap[T]) Push(item T) {
	h.data = append(h.data, item)
	h.upHeap(len(h.data) - 1)
}

func (h *Heap[T]) Pop() (T, bool) {
	if len(h.data) == 0 {
		var zero T
		return zero, false
	}

	top := h.data[0]
	last := len(h.data) - 1

	h.swap(0, last)
	h.data = h.data[:last]
	h.downHeap(0)

	return top, true
}

func (h *Heap[T]) Peek() (T, bool) {
	if len(h.data) == 0 {
		var zero T
		return zero, false
	}
	return h.data[0], true
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) upHeap(index int) {
	for {
		parent := (index - 1) / 2
		if index == 0 || !h.lessFunc(h.data[index], h.data[parent]) {
			break
		}
		h.swap(index, parent)
		index = parent
	}
}

func (h *Heap[T]) downHeap(index int) {
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < len(h.data) && h.lessFunc(h.data[left], h.data[smallest]) {
			smallest = left
		}
		if right < len(h.data) && h.lessFunc(h.data[right], h.data[smallest]) {
			smallest = right
		}
		if smallest == index {
			break
		}

		h.swap(index, smallest)
		index = smallest
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
