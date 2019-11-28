package main

import (
	"sync"
)

type Queue interface {
	Pop() interface{}
	Push(interface{}) error
}

type element struct {
	next  *element
	value interface{}
}

type SafeQueue struct {
	head, tail *element
	mux        sync.Mutex
	size       int32
	length     int32
}

func NewSafeQueue() *SafeQueue {
	s := &SafeQueue{
		head: nil,
		tail: nil,
		mux:  sync.Mutex{},
		size: 100,
	}
	return s
}

func (self *SafeQueue) Pop() interface{} {
	self.mux.Lock()
	defer self.mux.Unlock()
	if self.length <= 0 {
		return nil
	}
	if self.head == nil {
		return nil
	}

	tmp := self.head
	self.head = self.head.next
	tmp.next = nil
	self.length--
	return tmp.value
}

func (self *SafeQueue) Push(item interface{}) error {
	self.mux.Lock()
	defer self.mux.Unlock()
	ele := element{
		value: item,
	}
	if self.head == nil {
		self.head = &ele
		self.tail = &ele

	} else if self.tail == nil {
		self.head = &ele
		self.tail = &ele
	} else {
		self.tail.next = &ele
		self.tail = self.tail.next
	}
	self.length++
	return nil
}
