package lib

import (
	"fmt"
	"strings"
)

type Node struct {
	value interface{}
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Insert(d interface{}) error {
	if d == nil {
		return fmt.Errorf("cannot insert nil value")
	}

	newNode := &Node{value: d, next: nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}

	return nil
}

func (l *List) String() string {
	var sb strings.Builder

	for p := l.head; p != nil; p = p.next {
		sb.WriteString(fmt.Sprintf("-> %v ", p.value))
	}

	return sb.String()
}

func (l *List) Len() int {
	var count int

	for p := l.head; p != nil; p = p.next {
		count++
	}

	return count
}
