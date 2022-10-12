package interpreter

import "funge/internal/util"

type coord int

const (
	// represents an indexing of the axes on the instruction space
	x coord = iota
	y

	// Not an actual coordinate, but used to represent the number of coordinates
	_numCoords
)

// tuple is the discrete vector primitive used to represent the position or motion of
// the instruction pointer.
type tuple[T util.Numeric] [_numCoords]T

// IPointerLocation represents the location of the current instruction to be interpreted
type IPointerLocation tuple[int32]

// Translate returns a new IPointerLocation translated by the given tuple
func (i IPointerLocation) Translate(delta IPointerDelta) IPointerLocation {
	result := IPointerLocation{}
	for idx := range result {
		result[idx] = i[idx] + delta[idx]
	}

	return result
}

// IPointerDelta represents the direction of the current instruction to be interpreted
type IPointerDelta tuple[int32]

// constructors for cardinal directions

func North() IPointerDelta { return IPointerDelta{0, 1} }
func South() IPointerDelta { return IPointerDelta{0, -1} }
func East() IPointerDelta  { return IPointerDelta{1, 0} }
func West() IPointerDelta  { return IPointerDelta{-1, 0} }

type InstructionPointer struct {
	location IPointerLocation
	delta    IPointerDelta
}

// newInstructionPointer returns a new instruction pointer at (0, 0) with delta (0, 1)
func newInstructionPointer() *InstructionPointer {
	return &InstructionPointer{delta: East()}
}

// Delta returns the current delta of the instruction pointer
func (ip *InstructionPointer) Delta() IPointerDelta {
	return ip.delta
}

// SetDelta sets the delta of the instruction pointer
func (ip *InstructionPointer) SetDelta(delta IPointerDelta) {
	ip.delta = delta
}

// Location returns the current location of the instruction pointer
func (ip *InstructionPointer) Location() IPointerLocation {
	return ip.location
}

// Move moves the instruction pointer by its delta
func (ip *InstructionPointer) Move() {
	ip.location = ip.location.Translate(ip.delta)
}
