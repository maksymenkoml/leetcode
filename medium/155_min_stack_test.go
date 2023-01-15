package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test155(t *testing.T) {
	const (
		doConstruct = "minStack"
		doPush      = "push"
		doPop       = "pop"
		doTop       = "top"
		doGetMin    = "getMin"
	)

	type example struct {
		do  string
		val int
	}
	type examples []example
	s := []examples{
		[]example{
			{do: doConstruct, val: 0},
			{do: doPush, val: 2147483646},
			{do: doPush, val: 2147483646},
			{do: doPush, val: 2147483647},
			{do: doTop, val: 2147483647},
			{do: doPop},
			{do: doGetMin, val: 2147483646},
			{do: doPop},
			{do: doGetMin, val: 2147483646},
			{do: doPop},
			{do: doPush, val: 2147483647},
			{do: doTop, val: 2147483647},
			{do: doGetMin, val: 2147483647},
			{do: doPush, val: -2147483648},
			{do: doTop, val: -2147483648},
			{do: doGetMin, val: -2147483648},
			{do: doPop},
			{do: doGetMin, val: 2147483647},
		},
		[]example{
			{do: doConstruct, val: 0},
			{do: doPush, val: 5},
			{do: doPush, val: 3},
			{do: doPush, val: 4},
			{do: doGetMin, val: 3},
			{do: doPop},
			{do: doPop},
			{do: doTop, val: 5},
			{do: doGetMin, val: 5},
		},
		[]example{
			{do: doConstruct, val: 0},
			{do: doPush, val: -2},
			{do: doPush, val: 0},
			{do: doPush, val: -3},
			{do: doGetMin, val: -3},
			{do: doPop},
			{do: doTop, val: 0},
			{do: doGetMin, val: -2},
		},
	}

	var obj MinStack
	for _, tst := range s {
		for _, act := range tst {
			switch act.do {
			case doConstruct:
				obj = Constructor()
			case doPush:
				obj.Push(act.val)
			case doPop:
				obj.Pop()
			case doTop:
				require.Equal(t, act.val, obj.Top(), "Wrong TOP")
			case doGetMin:
				require.Equal(t, act.val, obj.GetMin(), "Wrong MIN")
			}
		}
	}
}

// ===

type MinStack struct {
	store [][]int
}

func Constructor() MinStack {
	res := MinStack{}
	return res
}

func (s *MinStack) Push(val int) {
	min := val
	if len(s.store) > 0 {
		lastMin := s.store[len(s.store)-1][1]
		if lastMin < min {
			min = lastMin
		}
	}
	s.store = append(s.store, []int{val, min})
}

func (s *MinStack) Pop() {
	s.store = s.store[:len(s.store)-1]
}

func (s *MinStack) Top() int {
	return s.store[len(s.store)-1][0]
}

func (s *MinStack) GetMin() int {
	return s.store[len(s.store)-1][1]
}

// ===

type MinStack2 struct {
	store []int
	mins  []int
}

func Constructor2() MinStack2 {
	res := MinStack2{}
	return res
}

func (s *MinStack2) Push(val int) {
	s.store = append(s.store, val)

	if s.mins == nil {
		s.mins = append(s.mins, val)
		return
	}

	var resMins []int
	for i := 0; i < len(s.mins); i++ {
		if s.mins[i] > val {
			if i != 0 {
				resMins = append(resMins, s.mins[:i]...)
			}
			resMins = append(resMins, val)
			resMins = append(resMins, s.mins[i:len(s.mins)]...)
			s.mins = resMins
			return
		}
	}
	s.mins = append(s.mins, val)
}

func (s *MinStack2) Pop() {
	cur := s.store[len(s.store)-1]

	s.store = s.store[:len(s.store)-1]

	var resMins []int
	for i := 0; i < len(s.mins); i++ {
		if s.mins[i] == cur {
			if i == 0 {
				s.mins = s.mins[1:len(s.mins)]
				return
			}

			resMins = append(resMins, s.mins[:i]...)

			if i != len(s.mins)-1 {
				resMins = append(resMins, s.mins[i+1:]...)
			}

			s.mins = resMins
			return
		}
	}
}

func (s *MinStack2) Top() int {
	return s.store[len(s.store)-1]
}

func (s *MinStack2) GetMin() int {
	return s.mins[0]
}
