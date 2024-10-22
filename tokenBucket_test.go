package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	tb := createTokenBucket(4, 3)
	tb.start()

	for range 30 {
		if tb.isAllowed() {
			fmt.Println(time.Now())
		} else {
			fmt.Println("Not allowed")
		}
	}
}
