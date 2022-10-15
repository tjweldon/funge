package interpreter

import (
	"funge/internal/util"
	"log"
)

type FungeStack struct {
    *util.Stack[rune]
}

// Duplicate pushes another instance of the current topmost value onto the stack
func (fs *FungeStack) Duplicate() {
    val, ok := fs.Pop()
    if !ok {
        log.Fatalln("stack underflow")
    }
    fs.Push(val)
    fs.Push(val)
}

// Swap inverts the order of the top two stack elements
func (fs *FungeStack) Swap() {
    val1, ok1 := fs.Pop()
    val2, ok2 := fs.Pop()
    if !ok1 || !ok2 {
        log.Fatalln("stack underflowq")
    }

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
        _, _ = fs.Pop()
    }
}
