package algorithm

import (
	"errors"
	"sync"
)

type Stack struct {
	items []interface{}
	sync.RWMutex
}

func (s *Stack) Len() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.items)
}

func (s *Stack) Push(value interface{}) {
	s.Lock()
	defer s.Unlock()
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (value interface{}, err error) {
	s.Lock()
	defer s.Unlock()
	if len(s.items) <= 0 {
		return nil, errors.New("empty stack")
	}
	value = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value, err
}
