package visual

import (
	"fmt"
	"funge/internal/vm"
	"image/color"

	"github.com/tjweldon/p5"
)

type axis int

const (
	xAxis axis = iota
	yAxis
	_numAxes
)

type Stack struct {
	*p5.Proc
	w, h, maxItems int
	stackChan      <-chan vm.FungeStack
	offsets        [_numAxes]float64
}

func NewStack(maxItems int, stackChan <-chan vm.FungeStack) *Stack {
	s := &Stack{
		w:         cellWH.Int() * 2,
		h:         cellWH.Int() * maxItems,
		maxItems:  maxItems,
		stackChan: stackChan,
		offsets:   [_numAxes]float64{},
	}

	s.Proc = p5.NewProc(s.w, s.h)
	s.Setup = s.getSetup()
	s.Draw = s.getDraw()

	return s
}

func (s *Stack) UseOffset(offsetX, offsetY float64) {
	s.offsets[xAxis] = offsetX
	s.offsets[yAxis] = offsetY
}

func (s *Stack) DrawBackground() {
	s.Stroke(color.Black)
	s.StrokeWidth(2)
	for i := 0; i < s.maxItems; i++ {
		s.Line(0, float64(i*cellWH.Int()), float64(s.w), float64(i*cellWH.Int()))
	}
}

func (s *Stack) getSetup() func() {
	s.UseOffset(0.3*cellWH.Float(), -0.25*cellWH.Float())
	return func() {
		s.Background(color.White)
	}
}

func (s *Stack) DrawContent() {
	stack, ok := <-s.stackChan
	if !ok {
		return
	}
	display := make([]rune, s.maxItems)
	for i := 0; i < s.maxItems; i++ {
		display[i] = 0
	}

	if len(stack.Slice()) > s.maxItems {
		// show only the top of the stack
		copy(display, stack.Slice()[len(stack.Slice())-s.maxItems:])
	} else {
		// show the whole stack
		copy(display, stack.Slice())
	}

	s.TextSize(cellWH.Float() / 1.5)
	for i, item := range display {
		var text string
		if item <= 0 {
			text = ""
		} else if item < ' ' {
			text = fmt.Sprintf("0x%0x", item)
		} else {
			text = string(item)
		}
		s.Text(
			text,
			s.offsets[xAxis], // just the offset
			float64(s.h)-float64(i*cellWH.Int())+s.offsets[yAxis], // cell height + offset from the bottom
		)
	}
}

func (s *Stack) getDraw() func() {
	return func() {
		s.DrawBackground()
		s.DrawContent()
	}
}
