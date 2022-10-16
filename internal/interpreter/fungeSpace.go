package interpreter

import (
	"bytes"
	"strings"
)

type FungeSpace [][]Instruction

func MakeSpaceFromBytes(raw []byte) FungeSpace {
	var w, h int
	lines := bytes.Split(raw, []byte{'\n'})
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	// height of the FungeSpace is the number of lines
	h = len(lines)

	// width of the FungeSpace is the length of the longest line
	for _, line := range lines {
		if w < len(line) {
			w = len(line)
		}
	}

	space := NewFungeSpace(w, h)
	for y_, line := range lines {
		for x_, inst := range bytes.Runes(line) {
			space[y_][x_] = Instruction(inst)
		}
	}

	return space
}

// NewFungeSpace returns a new FungeSpace with the given dimensions
func NewFungeSpace(width, height int) FungeSpace {
	space := make([][]Instruction, height)
	for i := range space {
		space[i] = make([]Instruction, width)
		for j := range space[i] {
			space[i][j] = ' '
		}
	}

	return space
}

// Get returns the Instruction at the given location
func (fs FungeSpace) Get(pointer InstructionPointer) rune {
	return rune(fs[pointer.location[y]][pointer.location[x]])
}

// GetXY returns the rune content of the cell in FungeSpace
func (fs FungeSpace) GetXY(x, y int) rune {
	return rune(fs[y][x])
}

// Set sets the Instruction at the given location
func (fs FungeSpace) Set(pointer InstructionPointer, instruction Instruction) {
	fs[pointer.location[y]][pointer.location[x]] = instruction
}

func (fs FungeSpace) String() string {
	lines := make([]string, len(fs))
	for y_, line := range fs {
		lines[y_] = "I" + string(line) + "I"
	}

	return strings.Join(lines, "\n")
}

func (fs FungeSpace) Size() (size tuple[int32]) {
	size = tuple[int32]{
		x: int32(len(fs[0])),
		y: int32(len(fs)),
	}

	return size
}
