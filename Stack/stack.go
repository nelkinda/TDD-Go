package Stack

type Stack struct {
	data []interface{}
}

func (s Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(i int) {
	s.data = append(s.data, i)
}

func (s *Stack) GetSize() int {
	return len(s.data)
}

func (s *Stack) Peek() interface{} {
	return s.data[len(s.data) - 1]
}

func (s *Stack) Pop() interface{} {
	el := s.Peek()
	s.data = s.data[:len(s.data) - 1]
	return el
}
