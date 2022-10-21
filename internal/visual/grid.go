package visual

import (
	"funge/internal/interpreter"
	"github.com/go-p5/p5"
	"image/color"
)

// Grid is the renderer for the FungeSpace and instruction pointer
type Grid struct {
	cellWidth, cellHeight float64
	offsetX, offsetY      float64
	content               interpreter.FungeSpace
}

// NewGrid creates a new Grid renderer
func NewGrid(cellWidth, cellHeight float64) *Grid {
	return &Grid{
		cellWidth:  cellWidth,
		cellHeight: cellHeight,
	}
}

// SetContent sets the FungeSpace to be rendered
func (g *Grid) SetContent(content interpreter.FungeSpace) {
	g.content = content
}

// UseOffset sets the offset for the position of each instruction in its grid cell
func (g *Grid) UseOffset(offsetX, offsetY float64) {
	g.offsetX = offsetX
	g.offsetY = offsetY
}

// DrawBackground draws the background of grid lines
func (g *Grid) DrawBackground() {
	p5.StrokeWidth(2)
	p5.Stroke(color.RGBA{R: 128, G: 128, B: 128, A: 255})

	// drawGrid horizontal lines
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

	// drawGrid vertical lines
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

// DrawContent draws the content of the FungeSpace
func (g *Grid) DrawContent() {
	// drawGrid the code
	for y, line := range g.content {
		yCanvas := g.cellHeight*float64(y) + g.offsetY
		for x, inst := range line {
			xCanvas := g.cellWidth*float64(x) + g.offsetX

			p5.Text(string(inst), xCanvas, yCanvas)
		}
	}
}

// FillCell fills the cell specified by the location with the given colour
func (g *Grid) FillCell(location interpreter.IPointerLocation, col color.Color) {
	xCanvas := g.cellWidth * float64(location[0])
	yCanvas := g.cellHeight * float64(location[1])

	p5.Fill(col)
	p5.Rect(xCanvas, yCanvas, g.cellWidth, g.cellHeight)
}

func setupGrid(interp *interpreter.Interpreter) func() {
	// set up the canvas
	space := interp.GetSpace()
	dimensions := space.Size()

	w, h = dimensions[0]*cellWH, dimensions[1]*cellWH

	interp.SetHandles(nil, &outBuf)

	setupFunc := func() {
		p5.Canvas(int(w), int(h))
		p5.Background(color.White)
	}

	return setupFunc
}

func drawGrid(interChan <-chan interpreter.Interpreter) func() {

	grid := NewGrid(float64(cellWH), float64(cellWH))
	grid.UseOffset(0.3*float64(cellWH), -0.25*float64(cellWH))

	renderFunc := func() {
		// do a tick
		inter := <-interChan
		grid.SetContent(inter.GetSpace())

		pointer := inter.Pointer()
		location := pointer.Location()

		// drawGrid where the instruction pointer is
		grid.FillCell(location, color.RGBA{R: 0, G: 255, B: 0, A: 255})
		p5.TextSize(cellWH / 1.5)
		grid.DrawBackground()
		grid.DrawContent()
	}

	return renderFunc
}
