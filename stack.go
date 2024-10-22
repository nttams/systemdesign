package main

// import (
// 	"errors"
// 	"sync"
// )

// type FifoQueueByArray struct {
// 	size         int
// 	currentIndex int
// 	data         []int
// 	m            sync.RWMutex
// }

// func createFifoQueueByArray(size int) *FifoQueueByArray {
// 	queue := FifoQueueByArray{
// 		size: size,
// 		currentIndex: 0,
// 		data: make([]int, size),
// 	}

// 	return &queue
// }

// func (queue *FifoQueueByArray) IsFull() bool {
// 	return queue.currentIndex >= queue.size
// }

// func (queue *FifoQueueByArray) IsEmpty() bool {
// 	return queue.currentIndex <= 0
// }

// func (queue *FifoQueueByArray) Put(value int) bool{
// 	queue.m.Lock()
// 	defer queue.m.Unlock()

// 	if queue.IsFull() {
// 		return false
// 	}

// 	queue.data[queue.currentIndex] = value
// 	queue.currentIndex += 1
// 	return true
// }

// func (queue *FifoQueueByArray) Pop() (int, error) {
// 	queue.m.Lock()
// 	defer queue.m.Unlock()

// 	if queue.IsEmpty() {
// 		return 0, errors.New("empty list")
// 	}

// 	queue.currentIndex -= 1
// 	return queue.data[queue.currentIndex], nil
// }
