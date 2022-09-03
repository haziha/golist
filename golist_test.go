package golist

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	l := New[int]()
	for i := 0; i < 5; i++ {
		l.PushBack(i)

		l.PushFront(i * i)
	}

	fmt.Println(l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	l.Remove(l.Front())
	l.Remove(l.Back())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
