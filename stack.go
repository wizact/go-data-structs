package stack

import "fmt"

type Item struct {
	Value string
}

func (c *Item) String() string {
	return fmt.Sprint(c.Value)
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	items []*Item
	count int
}

// Push adds an item to the stack.
func (s *Stack) Push(n *Item) {
	s.items = append(s.items, n)
	s.count++
}

// Peek returns the last item but not removing it from the stack.
func (s *Stack) Peek() *Item {
	if s.count == 0 {
		return nil
	}
	itemToPeek := s.items[s.count]
	return itemToPeek
}

// Pop removes and returns an item from the stack in last to first order.
func (s *Stack) Pop() *Item {
	if s.count == 0 {
		return nil
	}
	s.count--
	itemToPop := s.items[s.count]
	s.items = s.items[:s.count]
	return itemToPop
}
