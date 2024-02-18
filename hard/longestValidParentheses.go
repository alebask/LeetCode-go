package hard

type stack struct {
	values []int
}

func (s *stack) Len() int {
	return len(s.values)
}
func (s *stack) Push(vals ...int) {
	s.values = append(s.values, vals...)
}
func (s *stack) Pop() int {
	v := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return v
}
func (s *stack) Top() int {
	return s.values[len(s.values)-1]
}
func newStack() *stack {
	return &stack{values: make([]int, 0)}
}

func longestValidParentheses(s string) int {

	st := newStack()

	//-1 corresponds to '(', 0 corresponds to ')'
	for _, r := range s {
		if r == '(' {
			st.Push(-1)
			continue
		}

		if st.Len() > 0 && st.Top() == -1 {
			st.Pop()
			v := 2
			if st.Len() > 0 && st.Top() > 0 {
				v += st.Pop()
			}
			st.Push(v)
		} else if st.Len() > 0 && st.Top() > 0 {
			v := st.Pop()
			if st.Len() > 0 && st.Top() == -1 {
				st.Pop()
				v += 2
				if st.Len() > 0 && st.Top() > 0 {
					v += st.Pop()
				}
				st.Push(v)
			} else {
				st.Push(v, 0)
			}

		} else {
			st.Push(0)
		}
	}

	max := 0
	for _, v := range st.values {
		if v > max {
			max = v
		}
	}
	return max
}
