# Funge

This is a golang implementation of [Befunge](https://en.wikipedia.org/wiki/Befunge), an esoteric programming language.

The intent of this project is largely just for fun(ge).

## Usage

Clone the repository, and create a test.b98 file in your project directory. See [wikipedia](https://en.wikipedia.org/wiki/Befunge#Sample_Befunge-93_code) for example code.

Be careful with the random number generator example, it just spams random digits out of stdout as fast as it can until you tell it to stop.

Once you have your sample that you want to run, use

```bash
go run main.go
```

in the project directory and it will execute the befunge code.

You can pass a path to any text file as a positional command line argument to the interpreter, the default behavior wiht no arguments is equivalent to

```bash
go run main.go test.b98
```


## Goals

 - Implement the interpreter instructions per the befunge 93 spec. - Done!
 - Implement a visual mode that animates the execution of a befunge program.
 - Possibly extend to funge 98 and beyond...
 
