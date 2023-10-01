package sort

import "math"

type NodeStack struct {
	value int
	prev  *NodeStack
}

type Stack struct {
	length int
	head   *NodeStack
}

func (s *Stack) push(item int) {
	node := &NodeStack{value: item}
	s.length++
	if s.length == 1 {
		s.head = node
		return
	}
	node.prev = s.head
	s.head = node
}

func (s *Stack) pop() int {
	s.length = int(math.Max(0, float64(s.length-1)))
	if s.length == 0 {
		head := s.head
		s.head = nil
		if head == nil {
			return -1
		}
		return head.value
	}

	head := s.head
	s.head = head.prev

	return head.value
}

func (s *Stack) peek() int {
	if s.head == nil {
		return -1
	}
	return s.head.value
}
