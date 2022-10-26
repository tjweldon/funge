package visual

import (
	"funge/internal/vm"
	"time"

	"gioui.org/text"
	"github.com/tjweldon/p5"
)

// flint is just a bit of syntactic sugar for casting
type flint int

func (f flint) Float() float64 { return float64(f) }
func (f flint) Int() int       { return int(f) }

const cellWH flint = 64

func initFont(p *p5.Proc) {
	p.TextFont(
		text.Font{
			Typeface: "",
			Variant:  "Mono",
			Style:    text.Regular,
			Weight:   text.Normal,
		},
	)
}

func Visualise(inter *vm.Interpreter, cycleDuration time.Duration) {
	// runInterpreter is the befunge code vm 'runtime'
	runInterpreter := func(
		clock <-chan time.Time, // vm clock, one op per tick
		ic chan<- vm.Interpreter, // interpreter state output
		sc chan<- vm.FungeStack, // stack state output
	) {
		defer func() {
			close(ic)
			close(sc)
		}()
		for range clock {
			inter.RunFor(1)
			ic <- *inter
			sc <- inter.Stack()
			if inter.IsStopped() {
				return
			}
		}
	}

	// make interChan and supply to grid using setter injection
	interChan := make(chan vm.Interpreter)
	grid := NewGrid(cellWH.Float(), cellWH.Float(), inter)
	grid.SetIncoming(interChan)
	initFont(grid.Proc)

	// make stackChan and supply to stack on instantiation
	stackChan := make(chan vm.FungeStack)
	stack := NewStack(5, stackChan)
	initFont(stack.Proc)

	// make the clock tick according to the passed cycle time
	clock := time.NewTicker(cycleDuration)
	defer clock.Stop()

	// start the visualisation goroutines
	go grid.Run()
	go stack.Run()

	// let the funge vm run until execution completes (or forever)
	runInterpreter(clock.C, interChan, stackChan)
}
