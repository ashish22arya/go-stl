package queue

import ()

// Queue is the container that follows FIFO order.

type (
	queue struct {
		front 	*node
		back	*node
		length	int
	}

	node struct {
		data 	interface{}
		next	*node
	}
)

// NewQueue initialises and returns queue object
func NewQueue() *queue {
	return &queue{nil, nil, 0}
}

// Size returns the size of the current queue
func (q *queue) Size() int {
	return q.length;
}

// IsEmpty returns the boolean whether the queue is empty or not
func (q *queue) IsEmpty() bool {
	return (q.length == 0)
}

// Front returns the front element of the queue
func (q *queue) Front() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.front.data
}

// Back returns the back element of the queue
func (q *queue) Back() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.back.data
}

// Push inserts the new element at the back of the queue
func (q *queue) Push(d interface{}) {
	n := &node{d, nil}
	if q.length == 0 {
		q.front = n
	} else {
		q.back.next = n
	}
	q.back = n
	q.length++
}

// Pop removes the element from the front of the queue
func (q *queue) Pop() {
	if q.length == 0 {
		return
	}

	if q.length == 1 {
		q.front = nil
		q.back = nil
	} else {
		q.front = q.front.next
	}
	q.length--
}