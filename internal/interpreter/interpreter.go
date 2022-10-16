package interpreter

import (
	"fmt"
	"funge/internal/util"
	"io"
	"log"
	"os"
)

type Descriptor uint8

const (
	Stdin Descriptor = iota
	Stdout
	_descriptorCount
)

type handles [_descriptorCount]io.ReadWriter

const NoLimit int = iota

// Interpreter is the funge VM
type Interpreter struct {
	stringMode         bool
	stopped            bool
	instructionPointer *InstructionPointer
	stack              *FungeStack
	space              FungeSpace
	ioHandles          handles
	ticks              int
}

// NewInterpreter returns a new Interpreter
func NewInterpreter(code FungeSpace) *Interpreter {
	return &Interpreter{
		stopped:            false, // stops if true
		stringMode:         false, // treats instructions as characters if true
		instructionPointer: newInstructionPointer(code.Size()),
		ioHandles:          handles{os.Stdin, os.Stdout},
		stack:              NewFungeStack(),

		// this is the befunge-98 space i.e. the code.
		space: code,
	}
}

// SetHandles allows replacing the stdin and stdout of the interpreter with in-memory io.ReadWriters,
// or any other io.ReadWriters.
func (i *Interpreter) SetHandles(in, out io.ReadWriter) {
	if in != nil {
		i.ioHandles[Stdin] = in
	}
	if out != nil {
		i.ioHandles[Stdout] = out
	}
}

// GetSpace returns the code as a FungeSpace
func (i *Interpreter) GetSpace() FungeSpace {
	return i.space
}

// Pointer returns the current Instruction pointer value
func (i *Interpreter) Pointer() InstructionPointer {
	return *i.instructionPointer
}

// Stack returns a clone of the stack
func (i *Interpreter) Stack() *util.Stack[rune] {
	return i.stack.Clone()
}

// Tick executes the next Instruction and updates the Instruction pointer.
// Returns true if the interpreter is still running, false if it has stopped.
func (i *Interpreter) Tick() (bool, Instruction) {
	// get the current currentInstruction
	currentInstruction := i.space.Get(i.Pointer())

	// execute the currentInstruction
	i.stopped = i.execute(Instruction(currentInstruction))

	// update the currentInstruction pointer
	i.translate()

	// increment the ticks
	i.ticks++

	return i.stopped, Instruction(currentInstruction)
}

// RunFor executes the next Instruction until the interpreter stops or the tickLimit is reached.
// A ticks less than or equal to zero is interpreted as no limit on the number of ticks.
func (i *Interpreter) RunFor(ticks int) {
	debugFuncs := struct {
		silent, verbose func(i *Interpreter, ctx ...any)
	}{
		silent: func(*Interpreter, ...any) {},
		verbose: func(i *Interpreter, ctx ...any) {
			fmt.Println("instruction", ctx)
			fmt.Println("AFTER TICK")
			fmt.Println("pointer", i.instructionPointer.location)
			fmt.Println("delta", i.instructionPointer.delta)
			fmt.Println("stack", i.stack.Slice())
			fmt.Println("stringMode", i.stringMode)
			fmt.Println()
		},
	}

	i.run(ticks, debugFuncs.silent)
}

func (i *Interpreter) Run() {
	i.RunFor(NoLimit)
}

func (i *Interpreter) run(ticks int, debugOut func(*Interpreter, ...any)) {
	var (
		stopped bool
		inst    Instruction
	)
	stopped, inst = i.Tick()

	for !stopped {
		if ticks > 0 && i.ticks >= ticks {
			i.ticks = 0
			break
		}

		debugOut(i, inst)
		stopped, inst = i.Tick()
	}
	fmt.Println("Done!")
}

func (i *Interpreter) translate() {
	i.instructionPointer.Move()
}

// execute executes the given Instruction
func (i *Interpreter) execute(instruction Instruction) (stop bool) {
	instructionId := GetId(instruction)

	if instructionId == NoOp {
		return
	}

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

	if instructionId >= IPMovementStart && instructionId <= IPMovementEnd {
		i.instructionPointer.SetDelta(instructionId.NewDelta(i.stack))
		return
	}

	if instructionId >= StackManipulationStart && instructionId <= StackManipulationEnd {
		i.stack.ExecuteInstruction(instructionId)
		return
	}

	if instructionId >= ArithmeticStart && instructionId <= ArithmeticEnd {
		i.DoStackOp(instruction)
		return
	}

	if instructionId >= IOStart && instructionId <= IOEnd {
		i.DoIo(instruction)
		return
	}

	switch instructionId {
	case ReadAndPush:
		i.stack.Push(i.space.Get(i.Pointer()))
	case StringMode:
		i.stringMode = true
	case Stop:
		return true
	case Skip:
		i.translate()
	case Not:
		i.stack.evalUnaryOp(UnaryOperators(instructionId))
	case GreaterThan:
		i.DoStackOp(instruction)
	case Put:
		i.put()
	case Get:
		i.get()
	}

	return
}

func (i *Interpreter) get() {
	yCoord := i.stack.Pop()
	xCoord := i.stack.Pop()
	i.stack.Push(
		i.space.Get(
			InstructionPointer{location: IPointerLocation{xCoord, yCoord}},
		),
	)
}

func (i *Interpreter) put() {
	yCoord := i.stack.Pop()
	xCoord := i.stack.Pop()
	v := i.stack.Pop()
	i.space.Set(
		InstructionPointer{location: IPointerLocation{xCoord, yCoord}},
		Instruction(v),
	)
}

func (i *Interpreter) DoIo(instruction Instruction) {
	switch GetId(instruction) {
	case PrintChr:
		item := i.stack.Pop()
		if _, err := fmt.Fprintf(i.ioHandles[Stdout], "%s", string(item)); err != nil {
			log.Fatal(err)
		}
	case PrintInt:
		item := i.stack.Pop()
		if _, err := fmt.Fprintf(i.ioHandles[Stdout], "%d", item); err != nil {
			log.Fatal(err)
		}
	case ReadInt:
		var input int
		if _, err := fmt.Fscanf(i.ioHandles[Stdin], "%d", &input); err != nil {
			log.Fatal(err)
		} else {
			i.stack.Push(rune(input))
		}
	case ReadChr:
		var input rune
		if _, err := fmt.Fscanf(i.ioHandles[Stdin], "%s", &input); err != nil {
			log.Fatal(err)
		} else {
			i.stack.Push(input)
		}
	}
}

func (i *Interpreter) DoStackOp(instruction Instruction) {
	i.stack.evalBinaryOp(BinaryOperators(GetId(instruction)))
}
