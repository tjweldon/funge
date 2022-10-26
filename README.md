# Funge

This is a golang implementation of [Befunge](https://en.wikipedia.org/wiki/Befunge), an esoteric programming language.

The intent of this project is largely just for fun(ge).

## Goals

- Implement the interpreter instructions per the befunge 93 spec. - Done!
- Implement a visual mode that animates the execution of a befunge program. - Done!
- Possibly extend to funge 98 and beyond...

## Usage

Clone the repository, and create a test.b98 file in your project directory. See [wikipedia](https://en.wikipedia.org/wiki/Befunge#Sample_Befunge-93_code) for example code.

Be careful with the random number generator example, it just spams random digits out of stdout as fast as it can until you tell it to stop.

Once you have your sample that you want to run, use

```bash
go run main.go
```

in the project directory and it will execute the befunge code.

You can pass a path to any text file as a positional command line argument to the interpreter, the default behavior with no arguments is equivalent to supplying the argument as follows:

```bash
go run main.go test.b98
```

There are some example programs in the `./examples` directory.

The interpreter has a visual mode to allow you to watch the code execution in real time using (a fork of) go-p5.
The tick rate of the befunge vm can be controlled from the CLI. Note: this parameter does nothing outside visual mode.

```
Usage: main [--visual] [--tick TICK] [CODE]

Positional arguments:
  CODE                   path to the code file

Options:
  --visual, -v           If this arg is supplied, the interpreter is started in visual mode
  --tick TICK, -t TICK   Dictates the min length of time a vm 'cpu-cycle' can take in ms. [default: 100]
  --help, -h             display this help and exit

```

