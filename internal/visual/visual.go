package visual

import (
	"bytes"
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

const cellWH flint = 32

var w, h int32

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
	interChan := make(chan interpreter.Interpreter)
	grid := NewGrid(cellWH.Float(), cellWH.Float(), inter)
	grid.SetIncoming(interChan)

	go func(ic chan<- interpreter.Interpreter) {
		for {
			ic <- *inter
			inter.RunFor(1)
			time.Sleep(100 * time.Millisecond)
		}
	}(interChan)

	go grid.Run()
	time.Sleep(10 * time.Second)
}
