package visual

import (
	"funge/internal/interpreter"
	"github.com/tjweldon/p5"
	"image/color"
)

type axis int

const (
	xAxis axis = iota
	yAxis
	_numAxes
)

type Stack struct {
	w, h, maxItems int
	stackChan      <-chan interpreter.FungeStack
	offsets        [_numAxes]float64
}

func NewStack(maxItems int, stackChan <-chan interpreter.FungeStack) *Stack {
	return &Stack{
		w:         cellWH.Int(),
		h:         cellWH.Int() * maxItems,
		maxItems:  maxItems,
		stackChan: stackChan,
	}
}

func (s *Stack) UseOffset(offsetX, offsetY float64) {
	s.offsets[xAxis] = offsetX
	s.offsets[yAxis] = offsetY
}

func (s *Stack) DrawBackground() {
	for i := 0; i < s.maxItems; i++ {
		p5.Line(0, float64(i*cellWH.Int()), float64(s.w), float64(i*cellWH.Int()))
	}
}

func (s *Stack) Setup() {
	p5.Canvas(s.w, s.h)
	p5.Background(color.White)
	s.DrawBackground()
}

func (s *Stack) DrawContent() {
	stack := <-s.stackChan
	display := make([]rune, s.maxItems)
	for i := 0; i < s.maxItems; i++ {
		display[i] = ' '
	}

	if len(stack.Slice()) > s.maxItems {
		// show only the top of the stack
		copy(display, stack.Slice()[len(stack.Slice())-s.maxItems:])
	} else {
		// show the whole stack
		copy(display, stack.Slice())
	}

	copy(display, stack.Slice())
	for i, item := range display {
		p5.Text(
			string(item),
			s.offsets[xAxis], // just the offset
			float64(i*cellWH.Int())+s.offsets[yAxis], // cell height + offset
		)
	}
}
