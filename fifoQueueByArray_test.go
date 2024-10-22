package main

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	queue := createFifoQueue(10)

	{
		queue.Put(1)
		v, e := queue.Pop()

		if v != 1 || e != nil {
			t.Error("Wrong value 1")
		}

		if !queue.IsEmpty() {
			t.Error("Queue is supposed to be empty")
		}
	}

	{
		for i := range 13 {
			if queue.Put(i) {
				fmt.Printf("put ok %d\n", i)
			} else {
				fmt.Printf("put failed %d\n", i)
			}
		}

		for i := range 10 {
			v, e := queue.Pop()
			if e != nil || v != i {
				t.Errorf("Wrong value, expected %d, got %d", i, v)
			}
		}

		if !queue.IsEmpty() {
			t.Error("Queue is supposed to be empty")
		}

		for i := range 13 {
			if queue.Put(i) {
				fmt.Printf("put ok %d\n", i)
			} else {
				fmt.Printf("put failed %d\n", i)
			}
		}

		for i := range 10 {
			v, e := queue.Pop()
			if e != nil || v != i {
				t.Errorf("Wrong value, expected %d, got %d", i, v)
			}
		}

		if !queue.IsEmpty() {
			t.Error("Queue is supposed to be empty")
		}
	}
}
