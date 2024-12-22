package main

import hp "container/heap"

type state struct {
	x, y, dir, distance int
	path                [][]int
}

type states []state

func (h states) Len() int           { return len(h) }
func (h states) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h states) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *states) Push(x any) {
	*h = append(*h, x.(state))
}

func (h *states) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type heap struct {
	values *states
}

func newHeap() *heap {
	return &heap{values: &states{}}
}

func (h *heap) push(p state) {
	hp.Push(h.values, p)
}

func (h *heap) pop() state {
	i := hp.Pop(h.values)
	return i.(state)
}

func (h *heap) len() int {
	return len(*h.values)
}
