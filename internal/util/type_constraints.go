package util

// Numeric is a type constraint for discrete numeric types
type Numeric interface {
	int | int8 | int16 | int32 | uint8 | uint16 | uint32
}
