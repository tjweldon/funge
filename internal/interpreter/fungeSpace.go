package interpreter

type FungeSpace [][]instruction

// NewFungeSpace returns a new FungeSpace with the given dimensions
func NewFungeSpace(width, height int) FungeSpace {
	space := make([][]instruction, height)
	for i := range space {
		space[i] = make([]instruction, width)
	}

	return space
}

// Get returns the instruction at the given location
func (fs FungeSpace) Get(pointer InstructionPointer) rune {
	return rune(fs[pointer.location[y]][pointer.location[x]])
}

// Set sets the instruction at the given location
func (fs FungeSpace) Set(pointer InstructionPointer, instruction instruction) {
	fs[pointer.location[y]][pointer.location[x]] = instruction
}
