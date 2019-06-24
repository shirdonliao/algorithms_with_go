package stack

import (
"sync"
"errors"
)

type Stack struct {
	lock  sync.Mutex // 保证线程安全
	items []int
}

//
// 新建栈结构
//
func NewStack() *Stack {
	return &Stack{
		lock:  sync.Mutex{},
		items: make([]int, 0),
	}
}

//
// 入栈
//
func (s *Stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, v)
}

//
// 出栈
//
func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := len(s.items)
	if n == 0 {
		return 0, errors.New("[ERROR]: stack is empty")
	}

	top := s.items[n-1]
	s.items = s.items[n-1:] // 弹出栈顶元素
	return top, nil
}