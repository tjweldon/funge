package visual

import (
	"bytes"
	"funge/internal/interpreter"
	"gioui.org/text"
	"github.com/go-p5/p5"
	"time"
)

var outBuf bytes.Buffer

const cellWH = 32

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
	interChan := make(chan interpreter.Interpreter)
	go p5.Run(setupGrid(inter), drawGrid(interChan))

	for {
		interChan <- *inter
		inter.RunFor(1)
		time.Sleep(100 * time.Millisecond)
	}
}
