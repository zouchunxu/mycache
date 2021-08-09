package queue

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	q := NewQueue(1)
	q.Push(1)
	q.Push(2)
	v, ok := q.Pop()
	if ok {
		fmt.Println(v.(int))
	}
}
