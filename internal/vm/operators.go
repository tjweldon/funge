package vm

// Binary Operators
// ================

// BinOp is a function that takes two arguments and returns a single value
type BinOp func(x, y rune) rune

func add(x, y rune) rune {
	return x + y
}

func subtract(x, y rune) rune {
	return x - y
}

func multiply(x, y rune) rune {
	return x * y
}

func divide(x, y rune) rune {
	return x / y
}

func modulo(x, y rune) rune {
	return x % y
}

func greaterThan(x, y rune) rune {
	if x > y {
		return 1
	}

	return 0
}

// BinaryOperators is a map of instructions to their corresponding arithmetic operations
// as functions with the signature func(x, y rune) rune
func BinaryOperators(id InstructionId) BinOp {
	switch id {
	case Add:
		return add
	case Sub:
		return subtract
	case Mul:
		return multiply
	case Div:
		return divide
	case Mod:
		return modulo
	case GreaterThan:
		return greaterThan
	}

	return func(x, y rune) rune { return 0 }
}

// Unary Operators
// ===============

// UnOp is a function with the signature func(x rune) rune
type UnOp func(x rune) rune

func not(x rune) rune {
	if x == 0 {
		return 1
	}

	return 0
}

// UnaryOperators is a map of instructions to their corresponding unary operations
func UnaryOperators(id InstructionId) UnOp {
	switch id {
	case Not:
		return not
	}

	return func(x rune) rune { return 0 }
}
