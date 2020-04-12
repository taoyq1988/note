package stack

type Stack struct {
	data  []interface{}
	index int
}

func NewStack() *Stack {
	return &Stack{
		data:  make([]interface{}, 0),
		index: -1,
	}
}

func (s *Stack) Empty() bool {
	return s.index == -1
}

func (s *Stack) Top() interface{} {
	if s.index == -1 {
		return nil
	}
	return s.data[s.index]
}

func (s *Stack) Pop() (d interface{}) {
	if s.index == -1 {
		return nil
	}
	d = s.data[s.index]
	s.data = s.data[:s.index]
	s.index--
	return
}

func (s *Stack) Push(d interface{}) {
	s.index++
	s.data = append(s.data, d)
}
