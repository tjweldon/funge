package main

import (
	"funge/internal/interpreter"
	"funge/internal/visual"
	"io"
	"log"
	"os"
	"time"

	"github.com/alexflint/go-arg"
)

func LoadCode(path string) interpreter.FungeSpace {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(file)
	space := interpreter.MakeSpaceFromBytes(data)

	return space
}

type Cli struct {
	Code string `arg:"positional" help:"path to the code file" default:"./test.b98"`
}

var cli Cli

func init() {
	arg.MustParse(&cli)
}

func main() {
	interp := interpreter.NewInterpreter(LoadCode(cli.Code))

	visual.Visualise(interp)
	time.Sleep(10 * time.Second)
	// interp.Run()
}
