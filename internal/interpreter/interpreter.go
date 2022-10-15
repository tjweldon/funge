package interpreter

import "funge/internal/util"

// Interpreter is the funge VM
type Interpreter struct {
	stringMode         bool
	stopped            bool
	instructionPointer *InstructionPointer
	stack              *FungeStack
    space              FungeSpace
}

// NewInterpreter returns a new Interpreter
func NewInterpreter() *Interpreter {
	return &Interpreter{instructionPointer: newInstructionPointer()}
}

// Pointer returns the current instruction pointer value
func (i *Interpreter) Pointer() InstructionPointer {
	return *i.instructionPointer
}

// setDelta sets the delta of the instruction pointer
func (i *Interpreter) setDelta(delta IPointerDelta) {
	i.instructionPointer.delta = delta
}

// Stack returns a clone of the stack
func (i *Interpreter) Stack() *util.Stack[rune] {
	return i.stack.Clone()
}

// Tick executes the next instruction and updates the instruction pointer.
// Returns true if the interpreter is still running, false if it has stopped.
func (i *Interpreter) Tick() bool {
	if i.stopped {
		return false
	}

	// get the current currentInstruction
	currentInstruction := i.space.Get(i.Pointer())

	// execute the currentInstruction
	i.execute(instruction(currentInstruction))

	// update the currentInstruction pointer
	i.translate()

	return true
}

func (i *Interpreter) translate() {
	i.instructionPointer.Move()
}

// execute executes the given instruction
func (i *Interpreter) execute(instruction instruction) (stop bool) {
	instructionId := GetId(instruction)
	if i.stringMode && instructionId != StringMode {
		i.stack.Push(rune(instruction))
		return
	} else if i.stringMode && instructionId == StringMode {
		i.stringMode = false
		return
	}

	if instructionId >= NumericPushStart && instructionId <= NumericPushEnd {
		i.stack.Push(rune(instructionId - NumericPushStart))
		return
	}

	if instructionId >= IPMovementStart && instructionId <= IPMovementStart {
		i.setDelta(instructionId.NewDelta(i.stack))
		return
	}

	switch instructionId {
	case ReadAndPush:
		i.stack.Push(i.space.Get(i.Pointer()))
	case NoOp:
	case Stop:
		stop = true
	case Skip:
		i.translate()
	}
	return
}
