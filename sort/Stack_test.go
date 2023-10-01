package sort

import "testing"

func TestStack(t *testing.T) {
	s := Stack{}
	s.push(5)
	s.push(7)
	s.push(9)

	pp := s.pop()
	if pp != 9 {
		t.Fatalf("expecting 9 to be popped, but obtained %d", pp)
	}
	if s.length != 2 {
		t.Fatalf("expecting stack lenght to be 2, but obtained %d", s.length)
	}

	s.push(11)

	pp = s.pop()
	if pp != 11 {
		t.Fatalf("expecting 11 to be popped, but obtained %d", pp)
	}
	pp = s.pop()
	if pp != 7 {
		t.Fatalf("expecting 7 to be popped, but obtained %d", pp)
	}
	if s.peek() != 5 {
		t.Fatalf("expecting stack peek to be 5, but obtained %d", s.length)
	}
	pp = s.pop()
	if pp != 5 {
		t.Fatalf("expecting 5 to be popped, but obtained %d", pp)
	}
	pp = s.pop()
	if pp != -1 {
		t.Fatalf("expecting -1 to be popped, but obtained %d", pp)
	}
	if s.length != 0 {
		t.Fatalf("expecting stack lenght to be 0, but obtained %d", s.length)
	}

	s.push(69)
	if s.peek() != 69 {
		t.Fatalf("expecting stack peek to be 69, but obtained %d", s.length)
	}
	if s.length != 1 {
		t.Fatalf("expecting stack lenght to be 1, but obtained %d", s.length)
	}
}
