package visual

import (
	"bytes"
	"fmt"
	"funge/internal/interpreter"
	"log"
	"time"

	"gioui.org/text"
	"github.com/tjweldon/p5"
)

var outBuf bytes.Buffer

type flint int

func (f flint) Float() float64 { return float64(f) }
func (f flint) Int() int       { return int(f) }
func (f flint) Int32() int32   { return int32(f) }

const cellWH flint = 64

var font text.Font

func init() {
	p5.TextFont(
		text.Font{
			Typeface: "",
			Variant:  "Mono",
			Style:    text.Regular,
			Weight:   text.Normal,
		},
	)
}

func Visualise(inter *interpreter.Interpreter) {
	log.Println("Visualise: inter", inter)

	// runInterpreter is the befunge code execution goroutine
	runInterpreter := func(
		ic chan<- interpreter.Interpreter,
		sc chan<- interpreter.FungeStack,
		kill chan<- struct{},
	) {
		defer func() {
			close(ic)
			close(sc)
			close(kill)
		}()
		for {
			ic <- *inter
			sc <- inter.Stack()
			inter.RunFor(1)
			if inter.IsStopped() {
				kill <- struct{}{}
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}

	// make interChan and supply to grid on instantiation
	interChan := make(chan interpreter.Interpreter)
	grid := NewGrid(cellWH.Float(), cellWH.Float(), inter)
	grid.SetIncoming(interChan)

	// make stackChan and supply to stack on instantiation
	stackChan := make(chan interpreter.FungeStack)
	stack := NewStack(5, stackChan)

	// make completion signal channel
	done := make(chan struct{})

	// start the interpreter goroutine
	go runInterpreter(interChan, stackChan, done)

	// start the visualisation goroutines
	go grid.Run()
	go stack.Run()

	// let it run until execution completes
	select {
	case <-done:
		fmt.Println("Program complete!")
		return
	}
}
