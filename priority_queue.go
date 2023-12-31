package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

type PriorityQueue []*Pair

func (piq PriorityQueue) Len() int {
	return len(piq)
}

func (piq PriorityQueue) Less(i, j int) bool {
	return piq[i].priority > piq[j].priority
}

func (piq PriorityQueue) Swap(i, j int) {
	piq[i], piq[j] = piq[j], piq[i]
	piq[i].index = i
	piq[j].index = j
}

func (piq *PriorityQueue) Push(x interface{}) {
	n := len(*piq)
	item := x.(*Pair)
	item.index = n
	*piq = append(*piq, item)
}

func (piq *PriorityQueue) Pop() interface{} {
	old := *piq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*piq = old[0 : n-1]
	return item
}

func main() {
	pq := make(PriorityQueue, 5)

	pq[0] = &Pair{val: 1, priority: 1, index: 0}
	pq[1] = &Pair{val: 2, priority: 1, index: 1}
	pq[2] = &Pair{val: 3, priority: 1, index: 2}
	pq[3] = &Pair{val: 4, priority: 5, index: 3}
	pq[4] = &Pair{val: 5, priority: 9, index: 4}

	heap.Init(&pq)

	fmt.Println("PriorityQueue Enqueued: 1:1, 2:1, 3:1, 4:5, 5:9")
	fmt.Print("PriorityQueue Dequeued:")

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Pair)
		fmt.Printf(" %d:%d;", item.val, item.priority)
	}

	pairs := make([]*Pair, RangeInt)
	for i := range pairs {
		pairs[i] = &Pair{
			val:      rand.Intn(RangeInt),
			priority: rand.Intn(9) + 1,
		}
	}

	start := time.Now()
	for _, pair := range pairs {
		heap.Push(&pq, pair)
	}
	duration := time.Since(start)
	fmt.Printf("\nPriorityQueue enqueue time: %s\n", duration)

	count := 0
	start = time.Now()
	for pq.Len() > 0 {
		heap.Pop(&pq)
		count++
	}
	duration = time.Since(start)

	fmt.Printf("PriorityQueue dequeued items: %d\n", count)
	fmt.Printf("PriorityQueue dequeue time: %s\n", duration)
	fmt.Println("-----")
}
