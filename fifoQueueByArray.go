package main

import (
	"errors"
	"sync"
)

// Main idea: tail chases head
// Problem: but it will chase forever, how to loop around?
type FifoQueue struct {
	size int
	head int
	tail int
	data []int
	m    sync.RWMutex
}

func createFifoQueue(size int) *FifoQueue {
	queue := FifoQueue{
		size: size,
		head: 0,
		tail: 0,
		data: make([]int, size),
	}

	return &queue
}

func (queue *FifoQueue) IsFull() bool {
	return queue.head-queue.tail >= queue.size
}

func (queue *FifoQueue) IsEmpty() bool {
	return queue.head == queue.tail
}

func (queue *FifoQueue) Put(value int) bool {
	queue.m.Lock()
	defer queue.m.Unlock()

	if queue.IsFull() {
		return false
	}

	queue.data[queue.head%queue.size] = value
	queue.head += 1
	return true
}

func (queue *FifoQueue) Pop() (int, error) {
	queue.m.Lock()
	defer queue.m.Unlock()

	if queue.IsEmpty() {
		return 0, errors.New("empty list")
	}

	v := queue.data[queue.tail%queue.size]
	queue.tail += 1
	return v, nil
}
