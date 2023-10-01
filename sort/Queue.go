package sort

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	length int
	head   *Node
	tail   *Node
}

func (q *Queue) enqueue(item int) {
	node := &Node{value: item}
	q.length++
	if q.length == 1 {
		q.head = node
		q.tail = node
		return
	}
	q.tail.next = node
	q.tail = node
}

func (q *Queue) deque() int {
	if q.head == nil {
		return -1
	}
	q.length--

	h := q.head
	q.head = q.head.next
	h.next = nil

	if q.length == 0 {
		q.tail = nil
	}

	return h.value
}

func (q *Queue) peek() int {
	if q.head == nil {
		return -1
	}
	return q.head.value
}
