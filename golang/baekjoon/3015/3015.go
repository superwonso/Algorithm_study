package main

import (
	"bufio"
	"container/list" // stdlib
	"fmt"
	"os"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func (s *Stack) IsEmpty() bool {
	return s.v.Len() == 0
}

var r = bufio.NewReader(os.Stdin)

func main() {
	s_1 := NewStack()
	s_2 := NewStack()
	var n int
	fmt.Fscanf(r, "%d", &n)
	var height int
	var ans int
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &height)
		for !s_1.IsEmpty() && !s_2.IsEmpty() && height > s_1.v.Back().Value.(int) {
			ans += s_2.v.Back().Value.(int)
			s_1.Pop()
			s_2.Pop()
		}
		if s_1.IsEmpty() || s_2.IsEmpty() {
			s_1.Push(height)
			s_2.Push(1)
		} else {
			if s_1.v.Back().Value.(int) == height {
				var cur int = s_2.v.Back().Value.(int)
				s_1.Pop()
				s_2.Pop()
				if !s_1.IsEmpty() && !s_2.IsEmpty() {
					ans++
				}
				ans += cur
				s_1.Push(height)
				s_2.Push(cur + 1)
			} else {
				ans += 1
				s_1.Push(height)
				s_2.Push(1)
			}
		}
	}
	fmt.Println(ans)
}
