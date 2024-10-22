package main

import (
	"sync"
	"time"
)

type TokenBucket struct {
	size       int
	refillRate int
	tokenCount int
	m          sync.RWMutex
}

func createTokenBucket(bucketSize, refillRate int) *TokenBucket {
	tb := TokenBucket{
		size:       bucketSize,
		refillRate: refillRate,
		tokenCount: bucketSize,
	}

	return &tb
}

func (tb *TokenBucket) start() {
	go func() {
		for range time.Tick(time.Second * 1) {
			tb.m.Lock()
			tb.tokenCount += tb.refillRate
			if tb.tokenCount > tb.size {
				tb.tokenCount = tb.size
			}
			tb.m.Unlock()
		}
	}()
}

func (tb *TokenBucket) isAllowed() bool {
	tb.m.Lock()
	defer tb.m.Unlock()

	if tb.tokenCount <= 0 {
		return false
	}

	tb.tokenCount -= 1
	return true
}
