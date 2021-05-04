package stack

import (
	"strconv"
	"strings"
)

type Stack struct {
	data int32
	next *Stack
}

func (s Stack) Add(n int32) Stack {
	return Stack{n, &s}
}

func (s Stack) Pop() (Stack, int32) {
	return *s.next, s.data
}

func (s *Stack) String() string {
	if s == nil {
		return ""
	}
	sb := strings.Builder{}
	tmp := s
	for tmp.next != nil {
		v := strconv.FormatInt(int64(tmp.data), 10)
		sb.WriteString(v)
		sb.WriteString(" -> ")
		tmp = tmp.next
	}
	v := strconv.FormatInt(int64(tmp.data), 10)
	sb.WriteString(v)
	return sb.String()
}
