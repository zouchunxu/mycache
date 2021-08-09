package queue

import (
	"runtime"
	"sync/atomic"
)

type Queue struct {
	total  int64
	back   int64
	front  int64
	values []interface{}
}

func NewQueue(total int64) *Queue {
	q := &Queue{}
	q.total = total
	q.values = make([]interface{}, total)
	return q
}

func (q *Queue) Push(val interface{}) {
	for {
		back := atomic.LoadInt64(&q.back) % q.total
		pos := (back + 1) % q.total
		if !atomic.CompareAndSwapInt64(&q.back, back, pos) {
			runtime.Gosched()
			continue
		}
		q.values[back] = val
		return
	}
}

func (q *Queue) Pop() (interface{}, bool) {
	if q.front == q.back {
		return nil, false
	}
	for {
		front := atomic.LoadInt64(&q.front) % q.total
		back := atomic.LoadInt64(&q.back) % q.total
		if front == back {
			return nil, false
		}
		pos := (front + 1) % q.total
		if !atomic.CompareAndSwapInt64(&q.front, front, pos) {
			runtime.Gosched()
			continue
		}
		return q.values[front], true
	}
}
