package util

import "testing"

func TestStack_Push(t *testing.T) {
	testCase := struct {
		calls  []int
		want   []int
		actual *Stack[int]
	}{
		calls:  []int{100, 200, 300},
		want:   []int{300, 200, 100},
		actual: NewStack[int](),
	}
	for _, arg := range testCase.calls {
		testCase.actual.Push(arg)
	}

	if testCase.actual == nil {
		t.Errorf("actual is nil")
	}
	for idx, actualElem := range testCase.want {
		if testCase.actual.Slice()[idx] != actualElem {
			t.Errorf("actual[%d] = %d, want %d", idx, testCase.actual.Slice()[idx], actualElem)
		}
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack[string]()
	args := [3]string{"foo", "bar", "baz"}

	for _, arg := range args {
		stack.Push(arg)
	}

	for i := len(args) - 1; i >= 0; i-- {
		if stack.Pop() != args[i] {
			t.Errorf("stack.Pop() = %s, want %s", stack.Pop(), args[i])
		}
	}
}
