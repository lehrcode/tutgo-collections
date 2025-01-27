package main

import (
	"fmt"
	"github.com/lehrcode/tutgo-collections"
	"github.com/lehrcode/tutgo-collections/linkedlist"
	"github.com/lehrcode/tutgo-collections/slicelist"
)

func main() {
	var queue collections.Queue[int] = slicelist.New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	fmt.Println(queue)
	var next = queue.Deque()
	fmt.Println(next)
	fmt.Println(queue)

	var queue2 collections.Queue[int] = linkedlist.New[int]()
	queue2.Enqueue(1)
	queue2.Enqueue(2)
	queue2.Enqueue(3)
	queue2.Enqueue(4)
	fmt.Println(queue)
	var next2 = queue.Deque()
	fmt.Println(next2)
	fmt.Println(queue2)
}
