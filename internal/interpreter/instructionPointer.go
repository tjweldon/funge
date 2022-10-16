package interpreter

import (
	"funge/internal/util"
	"math/rand"
	"time"
)

type coord int

const (
	// represents an indexing of the axes on the Instruction space
	x coord = iota
	y

	// Not an actual coordinate, but used to represent the number of coordinates
	_numCoords
)

// tuple is the discrete vector primitive used to represent the position or motion of
// the Instruction pointer.
type tuple[T util.Numeric] [_numCoords]T

// IPointerLocation represents the location of the current Instruction to be interpreted
type IPointerLocation tuple[int32]

// Translate returns a new IPointerLocation translated by the given tuple
func (i IPointerLocation) Translate(delta IPointerDelta, size tuple[int32]) IPointerLocation {
	result := IPointerLocation{}
	for idx := range result {
		result[idx] = (i[idx] + delta[idx]) % size[idx]
	}

	return result
}

// IPointerDelta represents the direction of the current Instruction to be interpreted
type IPointerDelta tuple[int32]

// constructors for cardinal directions

func North() IPointerDelta { return IPointerDelta{0, -1} }
func South() IPointerDelta { return IPointerDelta{0, 1} }
func East() IPointerDelta  { return IPointerDelta{1, 0} }
func West() IPointerDelta  { return IPointerDelta{-1, 0} }

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func Random() IPointerDelta {
	choices := []IPointerDelta{
		North(),
		South(),
		East(),
		West(),
	}

	idx := rand.Intn(4)
	return choices[idx]
}

type InstructionPointer struct {
	location IPointerLocation
	delta    IPointerDelta
	size     tuple[int32]
}

// newInstructionPointer returns a new Instruction pointer at (0, 0) with delta (0, 1)
func newInstructionPointer(size tuple[int32]) *InstructionPointer {
	return &InstructionPointer{
		location: IPointerLocation{0, 0},
		delta:    East(),
		size:     size,
	}
}

// Delta returns the current delta of the Instruction pointer
func (ip *InstructionPointer) Delta() IPointerDelta {
	return ip.delta
}

// SetDelta sets the delta of the Instruction pointer
func (ip *InstructionPointer) SetDelta(delta IPointerDelta) {
	ip.delta = delta
}

// Location returns the current location of the Instruction pointer
func (ip *InstructionPointer) Location() IPointerLocation {
	return ip.location
}

// Move moves the Instruction pointer by its delta
func (ip *InstructionPointer) Move() {
	ip.location = ip.location.Translate(ip.delta, ip.size)
}
