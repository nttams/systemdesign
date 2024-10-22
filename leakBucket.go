package main

import (
	"sync"
)

// This rate limiter is basically a FIFO queue, wrapped and protected by a mutex
// For now, assuming the incomming request has type of int
type LeakBucket struct {
	queue FifoQueue
	mu *sync.RWMutex
}
