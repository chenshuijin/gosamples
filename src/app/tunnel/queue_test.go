package main

import (
	"testing"
)

var q = NewSafeQueue()

func TestQueue(t *testing.T) {
	q.Push(ForTmp{})
	_, ok := q.Pop().(ForTmp)
	if !ok {
		t.Errorf("err while pop from queue:%s", ok)
	}
	q.Pop()
}

func BenchmarkQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q.Push(ForTmp{b: i})
	}

	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkQueueParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Push(ForTmp{})
		}
	})
}

type ForTmp struct {
	a string
	b int
	c chan int
	d []string
	e *string
}
