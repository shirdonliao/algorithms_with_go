package stack

import "testing"

func TestStack(t *testing.T)  {

	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	top, err :=  s.Pop()
	if err != nil {
		t.Errorf("[ERROR]: stack.Pop() error. Got error: %v", err)
	}
	if top != 3{
		t.Errorf("[ERROR]: stack.Pop() failed. Got error: %v, expected 3", top)
	}
}