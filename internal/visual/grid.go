package visual

import (
	"funge/internal/vm"
	"github.com/tjweldon/p5"
	"image/color"
)

// Grid is the renderer for the FungeSpace and instruction pointer
type Grid struct {
	*p5.Proc
	w, h                  int
	cellWidth, cellHeight float64
	offsetX, offsetY      float64
	content               vm.FungeSpace
	interChan             <-chan vm.Interpreter
}

// NewGrid creates a new Grid renderer
func NewGrid(cellWidth, cellHeight float64, initial *vm.Interpreter) *Grid {
	g := &Grid{
		cellWidth:  cellWidth,
		cellHeight: cellHeight,
	}
	space := initial.GetSpace()
	dimensions := space.Size()
	w, h := dimensions[0]*cellWH.Int(), dimensions[1]*cellWH.Int()

	g.Proc = p5.NewProc(int(w), int(h))
	g.Setup = g.setupGrid(initial)
	g.Draw = g.drawGrid()

	return g
}

func (g *Grid) SetIncoming(incoming <-chan vm.Interpreter) {
	g.interChan = incoming
}

// SetContent sets the FungeSpace to be rendered
func (g *Grid) SetContent(content vm.FungeSpace) {
	g.content = content
}

// UseOffset sets the offset for the position of each instruction in its grid cell
func (g *Grid) UseOffset(offsetX, offsetY float64) {
	g.offsetX = offsetX
	g.offsetY = offsetY
}

// DrawBackground draws the background of grid lines
func (g *Grid) DrawBackground() {
	g.StrokeWidth(2)
	g.Stroke(color.RGBA{R: 128, G: 128, B: 128, A: 255})

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

		g.Line(start[0], start[1], end[0], end[1])
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

		g.Line(start[0], start[1], end[0], end[1])
	}
}

// DrawContent draws the content of the FungeSpace
func (g *Grid) DrawContent() {
	// drawGrid the code
	for y, line := range g.content {
		yCanvas := g.cellHeight*float64(y) + g.offsetY
		for x, inst := range line {
			xCanvas := g.cellWidth*float64(x) + g.offsetX

			g.Text(string(inst), xCanvas, yCanvas)
		}
	}
}

// FillCell fills the cell specified by the location with the given colour
func (g *Grid) FillCell(location vm.IPointerLocation, col color.Color) {
	xCanvas := g.cellWidth * float64(location[0])
	yCanvas := g.cellHeight * float64(location[1])

	g.Fill(col)
	g.Rect(xCanvas, yCanvas, g.cellWidth, g.cellHeight)
}

func (g *Grid) setupGrid(interp *vm.Interpreter) func() {
	// set up the canvas
	space := interp.GetSpace()
	dimensions := space.Size()

	w, h := dimensions[0]*cellWH.Int(), dimensions[1]*cellWH.Int()

	setupFunc := func() {
		g.Canvas(int(w), int(h))
		g.Background(color.White)
	}

	return setupFunc
}

func (g *Grid) drawGrid() func() {

	g.UseOffset(0.3*cellWH.Float(), 0.75*cellWH.Float())

	renderFunc := func() {
		// do a tick
		inter, ok := <-g.interChan
		if !ok {
			return
		}
		g.SetContent(inter.GetSpace())

		pointer := inter.Pointer()
		location := pointer.Location()

		// drawGrid where the instruction pointer is
		g.FillCell(location, color.RGBA{R: 0, G: 255, B: 0, A: 255})
		g.TextSize(cellWH.Float() / 1.5)
		g.DrawBackground()
		g.DrawContent()
	}

	return renderFunc
}
