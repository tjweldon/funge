package visual

import (
	"bytes"
	"funge/internal/interpreter"
	"funge/internal/util"
	"image/color"
	"time"

	"github.com/go-p5/p5"
)

var outBuf bytes.Buffer

func Visualise(interp *interpreter.Interpreter) {
	p5.Run(setup(interp), draw(interp))
}

const cellWH = 32

var w, h int32

func setup(interp *interpreter.Interpreter) func() {
	// set up the canvas
	space := interp.GetSpace()
	dimensions := space.Size()

	w, h = dimensions[0]*cellWH, dimensions[1]*cellWH
	w *= 2
    h = util.Max(w/2, h)

	interp.SetHandles(nil, &outBuf)

	setupFunc := func() {
		p5.Canvas(int(w), int(h))
		p5.Background(color.White)
	}

	return setupFunc
}

func draw(interp *interpreter.Interpreter) func() {

	// render a frame
	space := interp.GetSpace()
	offsets := struct{ X, Y float64 }{
		X: float64(cellWH),
		Y: float64(cellWH),
	}

	renderFunc := func() {
		// do a tick
		interp.RunFor(1)
        time.Sleep(time.Second/2)

		pointer := interp.Pointer()
		location := pointer.Location()

		// draw where the instruction pointer is
		p5.StrokeWidth(0)
		p5.Fill(color.RGBA{G: 128, A: 255})
		p5.Rect(
			float64(cellWH)*float64(location[0]),
			float64(cellWH)*float64(location[1]) + offsets.Y/4,
			float64(cellWH),
			float64(cellWH),
		)

		p5.TextSize(24)

		// draw the stdout
		p5.Text(
            util.WrapString("Stdout: " + outBuf.String(), 30),
			float64(w/2)+2*offsets.X,
			float64(cellWH),
		)

		// draw the code
		for y, line := range space {
			yCanvas := float64(cellWH)*float64(y) + offsets.Y
			for x, inst := range line {
				xCanvas := float64(cellWH)*float64(x) + offsets.X

				p5.Text(string(inst), xCanvas, yCanvas)
			}
		}
	}

	return renderFunc
}
