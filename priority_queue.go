// taken from https://github.com/oleiade/lane/blob/master/priority_queue.go as I can't go get for some reason
package main

import (
	"golang.org/x/exp/constraints"
)

type PriorityQueue[T any, P constraints.Ordered] struct {
	items      []*priorityQueueItem[T, P]
	itemCount  uint
	comparator func(lhs, rhs P) bool
}

func NewPriorityQueue[T any, P constraints.Ordered](heuristic func(lhs, rhs P) bool) *PriorityQueue[T, P] {
	items := make([]*priorityQueueItem[T, P], 1)
	items[0] = nil

	return &PriorityQueue[T, P]{
		items:      items,
		itemCount:  0,
		comparator: heuristic,
	}
}

func NewMaxPriorityQueue[T any, P constraints.Ordered]() *PriorityQueue[T, P] {
	return NewPriorityQueue[T](Maximum[P])
}

func NewMinPriorityQueue[T any, P constraints.Ordered]() *PriorityQueue[T, P] {
	return NewPriorityQueue[T](Minimum[P])
}

func Maximum[T constraints.Ordered](lhs, rhs T) bool {
	return lhs < rhs
}

func Minimum[T constraints.Ordered](lhs, rhs T) bool {
	return lhs > rhs
}

func (pq *PriorityQueue[T, P]) Push(value T, priority P) {
	item := newPriorityQueueItem(value, priority)

	pq.items = append(pq.items, item)
	pq.itemCount++
	pq.swim(pq.size())
}

func (pq *PriorityQueue[T, P]) Pop() (value T, priority P, ok bool) {

	if pq.size() < 1 {
		ok = false
		return
	}

	max := pq.items[1]
	pq.exch(1, pq.size())
	pq.items = pq.items[0:pq.size()]
	pq.itemCount--
	pq.sink(1)

	value = max.value
	priority = max.priority
	ok = true

	return
}

func (pq *PriorityQueue[T, P]) Head() (value T, priority P, ok bool) {

	if pq.size() < 1 {
		ok = false
		return
	}

	value = pq.items[1].value
	priority = pq.items[1].priority
	ok = true

	return
}

func (pq *PriorityQueue[T, P]) Size() uint {
	return pq.size()
}

func (pq *PriorityQueue[T, P]) Empty() bool {
	return pq.size() == 0
}

func (pq *PriorityQueue[T, P]) swim(k uint) {
	for k > 1 && pq.less(k/2, k) {
		pq.exch(k/2, k)
		k /= 2
	}
}

func (pq *PriorityQueue[T, P]) sink(k uint) {
	for 2*k <= pq.size() {
		j := 2 * k

		if j < pq.size() && pq.less(j, j+1) {
			j++
		}

		if !pq.less(k, j) {
			break
		}

		pq.exch(k, j)
		k = j
	}
}

func (pq *PriorityQueue[T, P]) size() uint {
	return pq.itemCount
}

func (pq *PriorityQueue[T, P]) less(lhs, rhs uint) bool {
	return pq.comparator(pq.items[lhs].priority, pq.items[rhs].priority)
}

func (pq *PriorityQueue[T, P]) exch(lhs, rhs uint) {
	pq.items[lhs], pq.items[rhs] = pq.items[rhs], pq.items[lhs]
}

type priorityQueueItem[T any, P constraints.Ordered] struct {
	value    T
	priority P
}

func newPriorityQueueItem[T any, P constraints.Ordered](value T, priority P) *priorityQueueItem[T, P] {
	return &priorityQueueItem[T, P]{
		value:    value,
		priority: priority,
	}
}
