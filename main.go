package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")

	tb := createTokenBucket(10, 3)
	tb.start()

	for {
		if tb.isAllowed() {
			fmt.Println(time.Now())
		}
	}
}
