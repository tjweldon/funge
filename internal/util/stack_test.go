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
	for idx, expected := range testCase.want {
		if actual := testCase.actual.Pop(); actual != expected {
			t.Errorf("actual[%d] = %d, want %d", idx, testCase.actual.Slice()[idx], expected)
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
		item := stack.Pop()
		if item != args[i] {
			t.Errorf("stack.Pop() = %s, want %s", item, args[i])
		}
	}
}

func TestStack_Pop_Empty(t *testing.T) {
	stack := NewStack[int]()
	zero := stack.Pop()
	if zero != 0 {
		t.Errorf("stack.Pop() = %d, want 0", zero)
	}
}

func TestStack_Clone(t *testing.T) {
	stack := NewStack[int]()
	args := [3]int{100, 200, 300}

	for _, arg := range args {
		stack.Push(arg)
	}

	clone := stack.Clone()
	for range args {
		cItem := clone.Pop()
		sItem := stack.Pop()
		if cItem != sItem {
			t.Errorf("clone.Pop() != stack.Pop(), %d != %d", cItem, sItem)
		}
	}
}
