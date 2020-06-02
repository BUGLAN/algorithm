package algorithm

type Stack struct {
	items []interface{}
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) Push(value interface{}) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() interface{} {
	value := s.items[s.Len()-1]
	s.items = s.items[:s.Len()-1]
	return value
}
