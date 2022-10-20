package visual

import (
	"bytes"
	"funge/internal/interpreter"
	"gioui.org/text"
	"image/color"

	"github.com/go-p5/p5"
)

var outBuf bytes.Buffer

func Visualise(interp *interpreter.Interpreter) {
	p5.Run(setup(interp), draw(interp))
}

const cellWH = 64

var w, h int32

var font text.Font

func init() {

	p5.TextFont(font)
}

func setup(interp *interpreter.Interpreter) func() {
	// set up the canvas
	space := interp.GetSpace()
	dimensions := space.Size()

	w, h = dimensions[0]*cellWH, dimensions[1]*cellWH
	w *= 2

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

	grid := NewGrid(float64(cellWH), float64(cellWH), space)
	grid.UseOffset(float64(cellWH)/2.0, float64(cellWH)/2.0)

	renderFunc := func() {
		// do a tick
		interp.RunFor(1)

		pointer := interp.Pointer()
		location := pointer.Location()

		// draw where the instruction pointer is
		grid.FillCell(location, color.RGBA{R: 0, G: 255, B: 0, A: 255})

		p5.TextSize(cellWH / 2)
		grid.DrawBackground()
		grid.DrawContent()
	}

	return renderFunc
}

type Grid struct {
	cellWidth, cellHeight float64
	offsetX, offsetY      float64
	content               interpreter.FungeSpace
}

func NewGrid(cellWidth, cellHeight float64, content interpreter.FungeSpace) *Grid {
	return &Grid{
		cellWidth:  cellWidth,
		cellHeight: cellHeight,
		content:    content,
	}
}

func (g *Grid) SetContent(content interpreter.FungeSpace) {
	g.content = content
}

func (g *Grid) UseOffset(offsetX, offsetY float64) {
	g.offsetX = offsetX
	g.offsetY = offsetY
}

func (g *Grid) DrawBackground() {
	p5.StrokeWidth(2)
	p5.Stroke(color.RGBA{R: 128, G: 128, B: 128, A: 255})

	// draw horizontal lines
	for y := 0; y < len(g.content); y++ {
		start := [2]float64{
			0,                         // x
			g.cellHeight * float64(y), // y
		}

		end := [2]float64{
			g.cellWidth * float64(len(g.content[y])), // x
			g.cellHeight * float64(y),                // y
		}

		p5.Line(start[0], start[1], end[0], end[1])
	}

	// draw vertical lines
	for x := 0; x < len(g.content[0]); x++ {
		start := [2]float64{
			g.cellWidth * float64(x), // x
			0,                        // y
		}

		end := [2]float64{
			g.cellWidth * float64(x),               // x
			g.cellHeight * float64(len(g.content)), // y
		}

		p5.Line(start[0], start[1], end[0], end[1])
	}
}

func (g *Grid) DrawContent() {
	// draw the code
	for y, line := range g.content {
		yCanvas := g.cellHeight*float64(y) + g.offsetY
		for x, inst := range line {
			xCanvas := g.cellWidth*float64(x) + g.offsetX

			p5.Text(string(inst), xCanvas, yCanvas)
		}
	}
}

func (g *Grid) FillCell(location interpreter.IPointerLocation, col color.Color) {
	xCanvas := g.cellWidth * float64(location[0])
	yCanvas := g.cellHeight * float64(location[1])

	p5.Fill(col)
	p5.Rect(xCanvas, yCanvas, g.cellWidth, g.cellHeight)
}
