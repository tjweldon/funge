package interpreter

import (
	"funge/internal/util"
)

type FungeStack struct {
	*util.Stack[rune]
}

func NewFungeStack() *FungeStack {
	return &FungeStack{util.NewStack[rune]()}
}

// Duplicate pushes another instance of the current topmost value onto the stack
func (fs *FungeStack) Duplicate() {
	val := fs.Pop()
	fs.Push(val)
	fs.Push(val)
}

// Swap inverts the order of the top two stack elements
func (fs *FungeStack) Swap() {
	val1 := fs.Pop()
	val2 := fs.Pop()

	fs.Push(val1)
	fs.Push(val2)
}

// ExecuteInstruction carries out stack Manipulation instructions. If any other instructionId is passed, it's a no-op.
func (fs *FungeStack) ExecuteInstruction(instructionId InstructionId) {
	switch instructionId {
	case Duplicate:
		fs.Duplicate()
	case Swap:
		fs.Swap()
	case Pop:
		_ = fs.Pop()
	}
}

func (fs *FungeStack) evalBinaryOp(operation BinOp) {
	val1 := fs.Pop()
	val2 := fs.Pop()

	fs.Push(operation(val1, val2))
}

func (fs *FungeStack) evalUnaryOp(operation UnOp) {
	val := fs.Pop()

	fs.Push(operation(val))
}
