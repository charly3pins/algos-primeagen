package sort

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{}
	q.enqueue(5)
	q.enqueue(7)
	q.enqueue(9)

	dq := q.deque()
	if dq != 5 {
		t.Fatalf("expecting 5 to be dequeued, but obtained %d", dq)
	}
	if q.length != 2 {
		t.Fatalf("expecting queue lenght to be 2, but obtained %d", q.length)
	}

	q.enqueue(11)

	dq = q.deque()
	if dq != 7 {
		t.Fatalf("expecting 7 to be dequeued, but obtained %d", dq)
	}
	dq = q.deque()
	if dq != 9 {
		t.Fatalf("expecting 9 to be dequeued, but obtained %d", dq)
	}
	if q.peek() != 11 {
		t.Fatalf("expecting queue peek to be 11, but obtained %d", q.length)
	}
	dq = q.deque()
	if dq != 11 {
		t.Fatalf("expecting 11 to be dequeued, but obtained %d", dq)
	}
	dq = q.deque()
	if dq != -1 {
		t.Fatalf("expecting -1 to be dequeued, but obtained %d", dq)
	}
	if q.length != 0 {
		t.Fatalf("expecting queue lenght to be 0, but obtained %d", q.length)
	}

	q.enqueue(69)
	if q.peek() != 69 {
		t.Fatalf("expecting queue peek to be 69, but obtained %d", q.length)
	}
	if q.length != 1 {
		t.Fatalf("expecting queue lenght to be 1, but obtained %d", q.length)
	}
}
