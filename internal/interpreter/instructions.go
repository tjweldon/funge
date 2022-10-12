package interpreter

import (
	"funge/internal/util"
	"log"
)

type InstructionId int

const (
	// default instruction
	ReadAndPush InstructionId = iota

	// special instructions
	NoOp       // no operation
	Stop       // stop execution
	Skip       // skip the next instruction
	StringMode // skip the next instruction

	// PushZero is equivalent to ReadAndPush 0
	PushZero  // push 0
	PushOne   // push 1
	PushTwo   // push 2
	PushThree // push 3
	PushFour  // push 4
	PushFive  // push 5
	PushSix   // push 6
	PushSeven // push 7
	PushEight // push 8
	PushNine  // push 9

	// Instruction Pointer Movement instructions
	MoveNorth        // move north
	MoveEast         // move east
	MoveSouth        // move south
	MoveWest         // move west
	MoveEastOrWest   // pop a value off the stack and move east if the value is 0, west otherwise
	MoveNorthOrSouth // pop a value off the stack and move north if the value is 0, south otherwise
	MoveRandom       // move in a random direction

	// Stack Manipulation instructions
	Duplicate // duplicate the top value on the stack
	Swap      // swap the top two values on the stack
	Pop       // pop the top value off the stack

	// Arithmetic instructions
	Add // pop a and b, push a + b
	Sub // pop a and b, push a - b
	Mul // pop a and b, push a * b
	Div // pop a and b, push a / b
	Mod // pop a and b, push a % b

	// Logical instructions
	GreaterThan // pop a and b, push 1 if a > b, 0 otherwise
	Not         // pop a push 1 if a == 0, 0 otherwise

	// I/O instructions
	PrintInt // pop a, print a as an integer
	PrintChr // pop a, print a as a character
	ReadInt  // read an integer from stdin, push it
	ReadChr  // read a character from stdin, push its ascii value

	// no actual instructions after this point
	instructionCount

	// Instruction Categories
	SpecialStart = NoOp
	SpecialEnd   = Skip

	NumericPushStart = PushZero
	NumericPushEnd   = PushNine

	IPMovementStart = MoveEast
	IPMovementEnd   = MoveRandom

	StackManipulationStart = Duplicate
	StackManipulationEnd   = Pop

	ArithmeticStart = Add
	ArithmeticEnd   = Mod

	LogicalStart = GreaterThan
	LogicalEnd   = Not

	IOStart = PrintInt
	IOEnd   = ReadChr
)

// instruction is a single cell in the instruction space
type instruction rune

var instructionMap = [instructionCount]instruction{
	// default instruction
	ReadAndPush:      '\x00',
	NoOp:             ' ',
	Stop:             '@',
	Skip:             '#',
	StringMode:       '"',
	PushZero:         '0',
	PushOne:          '1',
	PushTwo:          '2',
	PushThree:        '3',
	PushFour:         '4',
	PushFive:         '5',
	PushSix:          '6',
	PushSeven:        '7',
	PushEight:        '8',
	PushNine:         '9',
	MoveNorth:        '^',
	MoveEast:         '>',
	MoveSouth:        'v',
	MoveWest:         '<',
	MoveEastOrWest:   '_',
	MoveNorthOrSouth: '|',
	MoveRandom:       '?',
	Duplicate:        ':',
	Swap:             '\\',
	Pop:              '$',
	Add:              '+',
	Sub:              '-',
	Mul:              '*',
	Div:              '/',
	Mod:              '%',
	GreaterThan:      '`',
	Not:              '!',
	PrintInt:         '.',
	PrintChr:         ',',
	ReadInt:          '&',
	ReadChr:          '~',
}

func (id InstructionId) NewDelta(stack *util.Stack[rune]) IPointerDelta {
	switch id {
	case MoveNorth:
		return North()
	case MoveEast:
		return East()
	case MoveSouth:
		return South()
	case MoveWest:
		return West()
	case MoveEastOrWest:
		if item, ok := stack.Pop(); !ok {
			log.Fatal("stack underflow")
		} else if item != rune(0) {
			return East()
		} else {
			return West()
		}
	case MoveNorthOrSouth:
		if item, ok := stack.Pop(); !ok {
			log.Fatal("stack underflow")
		} else if item != rune(0) {
			return North()
		} else {
			return South()
		}
	}

	return IPointerDelta{}
}

var inverseLookup = make(map[rune]InstructionId)

func init() {
	for id, inst := range instructionMap {
		inverseLookup[rune(inst)] = InstructionId(id)
	}
}

func GetId(inst instruction) InstructionId {
	id, ok := inverseLookup[rune(inst)]
	if !ok {
		return ReadAndPush
	}
	return id
}
