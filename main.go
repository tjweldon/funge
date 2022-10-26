package main

import (
	"funge/internal/visual"
	"funge/internal/vm"
	"github.com/alexflint/go-arg"
	"io"
	"log"
	"os"
	"time"
)

func LoadCode(path string) vm.FungeSpace {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	space := vm.MakeSpaceFromBytes(data)

	return space
}

type Cli struct {
	Code   string `arg:"positional" help:"path to the code file" default:"./test.b98"`
	Visual bool   `arg:"-v, --visual" help:"If this arg is supplied, the interpreter is started in visual mode"`
	Tick   int    `arg:"-t, --tick" help:"Dictates the min length of time a vm 'cpu-cycle' can take in ms." default:"100"`
}

func (c *Cli) CycleTime() time.Duration {
	return time.Duration(c.Tick) * time.Millisecond
}

var cli Cli

func init() {
	arg.MustParse(&cli)
}

func main() {
	interpreter := vm.NewInterpreter(LoadCode(cli.Code))
	if cli.Visual {
		visual.Visualise(interpreter, cli.CycleTime())
	} else {
		interpreter.Run()
	}
}
