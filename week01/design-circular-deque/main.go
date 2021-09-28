package main

type MyCircularDeque struct {
	cache    []int
	capacity int
	length   int
	front    int
	rear     int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		cache:    make([]int, k),
		capacity: k,
		front:    1,
	}
}

func (q *MyCircularDeque) InsertFront(value int) bool {
	if q.IsFull() {
		return false
	}

	q.length++
	q.front--
	if q.front == -1 {
		q.front = q.capacity - 1
	}
	q.cache[q.front] = value
	return true
}

func (q *MyCircularDeque) InsertLast(value int) bool {
	if q.IsFull() {
		return false
	}

	q.length++
	q.rear++
	if q.rear == q.capacity {
		q.rear = 0
	}
	q.cache[q.rear] = value
	return true
}

func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}

	q.length--
	q.front++
	if q.front == q.capacity {
		q.front = 0
	}
	return true
}

func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}

	q.length--
	q.rear--
	if q.rear == -1 {
		q.rear = q.capacity - 1
	}
	return true
}

func (q *MyCircularDeque) GetFront() int {
	if q.IsEmpty() {
		return -1
	}

	return q.cache[q.front]
}

func (q *MyCircularDeque) GetRear() int {
	if q.IsEmpty() {
		return -1
	}

	return q.cache[q.rear]
}

func (q *MyCircularDeque) IsEmpty() bool {
	return q.length == 0
}

func (q *MyCircularDeque) IsFull() bool {
	return q.length == q.capacity
}
