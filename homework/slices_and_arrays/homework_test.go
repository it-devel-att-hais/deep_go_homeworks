package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values   []int
	size     int
	front    int
	rear     int
	overflow bool
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
		size:   size,
		front:  -1,
		rear:   -1,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.front == -1 {
		q.front = 0
	}
	if q.Full() {
		return false
	}

	if q.rear+1 >= q.size {
		q.rear = (q.rear + 1) % q.size
		q.overflow = true
	} else {
		q.rear++
	}
	q.values[q.rear] = value
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
		q.overflow = false
		return true
	}
	q.front++
	if q.front >= q.size {
		q.front = q.front % q.size
	}
	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.front]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.rear]
}

func (q *CircularQueue) Empty() bool {
	return q.rear == -1 && q.front == -1
}

func (q *CircularQueue) Full() bool {
	if q.front == 0 && q.rear == q.size-1 {
		return true
	}
	if q.overflow && q.front == q.rear-1 {
		return true
	}
	return false
}

func (q *CircularQueue) Print() {
	front := q.front
	if front == -1 {
		fmt.Println("[]")
		return
	}

	values := make([]int, 0, q.size)

	fmt.Println(q.front, q.rear, q.values)
	for {
		if front == q.size {
			front = front % q.size
		}
		values = append(values, q.values[front])
		if front == q.rear {
			break
		}
		front++
		if front == q.size {
			front = front % q.size
			values = append(values, q.values[front])
		}
		if front == q.rear {
			break
		}
	}
	fmt.Println(values)
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
