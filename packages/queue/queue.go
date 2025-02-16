package queue

import (
	"errors"
)

type Queue[T any] struct {
	theArray    []T
	front       int
	back        int
	currentSize int
}

func NewQueue[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		theArray:    make([]T, capacity),
		front:       0,
		currentSize: 0,
		back:        -1,
	}
}

func (q *Queue[T]) MakeEmpty() {
	q.currentSize = 0
	q.front = 0
	q.back = -1
}

func (q *Queue[T]) IsEmpty() bool {
	return q.currentSize == 0
}

func (q *Queue[T]) IsFull() bool {
	return q.currentSize == len(q.theArray)
}

func (q *Queue[T]) Enqueue(element T) {
	if q.IsFull() {
		q.theArray = append(q.theArray, element)
	}
	q.back = q.increment(q.back)
	q.theArray[q.back] = element
	q.currentSize++
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zerovalue T
		return zerovalue, errors.New("Queue empty")
	}

	q.currentSize--
	element := q.theArray[q.front]
	q.front++
	return element, nil
}

func (q *Queue[T]) increment(x int) int {
	if x++; x == len(q.theArray) {
		x = 0
	}
	return x
}
