package ds

type Stack struct {
    list LinkedList
}

func (s *Stack) Push(value string) {  // add new head node
	s.list.InsertAt(0, value)
}

func (s *Stack) Pop() (string, bool) { // remove the head
	if s.list.IsEmpty() {
		return "", false
	}
	stackHead := s.list.Head.data
	s.list.RemoveAt(0)
	return stackHead, true
}
