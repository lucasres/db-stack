package internal

type Stack struct {
	Values []string
}

func (s *Stack) Push(value string) {
	s.Values = append(s.Values, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Values) == 0
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	// get last index
	index := s.GetLastIndex()
	// get element
	element := s.Values[index]
	// remove fro values
	s.Values = s.Values[:index]
	return element, true
}

func (s *Stack) GetLastIndex() int {
	return len(s.Values) - 1
}

func (s *Stack) Get() string {
	if len(s.Values) == 0 {
		return ""
	}

	index := s.GetLastIndex()

	return s.Values[index]
}

func NewStack() *Stack {
	return &Stack{
		Values: make([]string, 0),
	}
}
