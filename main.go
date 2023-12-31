package main

import (
	"fmt"
	"math/rand"
	"time"
)

const rangeInt = 10_000_000

func main() {
	priorityQueue := NewMaxPriorityQueue[int, int]()

	priorityQueue.Push(1, 1)
	priorityQueue.Push(2, 1)
	priorityQueue.Push(3, 1)
	priorityQueue.Push(4, 5)
	priorityQueue.Push(5, 9)

	fmt.Println("PriorityQueue Enqueued: 1:1, 2:1, 3:1, 4:5, 5:9")
	fmt.Print("PriorityQueue Dequeued:")

	for priorityQueue.size() > 0 {
		val, priority, _ := priorityQueue.Pop()
		fmt.Printf(" %d:%d;", val, priority)
	}

	pairs := make([]*Pair, rangeInt)
	for i := range pairs {
		pairs[i] = &Pair{
			val:      rand.Intn(rangeInt),
			priority: rand.Intn(9) + 1,
		}
	}

	start := time.Now()
	for _, pair := range pairs {
		priorityQueue.Push(pair.val, pair.index)
	}

	duration := time.Since(start)
	fmt.Printf("\nPriorityQueue enqueue time: %s\n", duration)

	count := 0
	start = time.Now()
	for priorityQueue.size() > 0 {
		priorityQueue.Pop()
		count++
	}
	duration = time.Since(start)

	fmt.Printf("PriorityQueue dequeued items: %d\n", count)
	fmt.Printf("PriorityQueue dequeue time: %s\n", duration)
	fmt.Println("-----")
}
