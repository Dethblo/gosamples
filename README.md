# godethblo
A collection of Go Language code used to experiment and try different concepts.

# Setup
The code is set up to follow the GO language conventions for packaging.  The /cmd tree represents multiple GO programs (each with their own 'main').  The /internal tree is packaged code (not intended to be shared outside the applications).

Reference: `https://go.dev/doc/code`

## How the structure was created
1. Create the godethblo project directory
2. Change to the ./godethblo directory
3. Using my Github.com repository Execute: `go mod init github.com/Dethblo/godethblo`
4. Create the `./cmd/hello` directory and `hello.go` source code.
5. Execute the code: `go run cmd/hello/hello.go`
6. Create other directories and code as needed.


# Application Examples
## Bank
A program that illustrates if/else, for, read/write files, and basic error handling.

        go run cmd/bank/bank.go

## Collections
A program that illustrates usage of collections in Go.

        go run cmd/collections/collections.go

## Functions
A program that illustrates usage of functions in Go.

        go run cmd/functions/functions.go

## Hello
A simple `hello world` program.

        go run cmd/hello/hello.go

## Investment Calculator
A program that illustrates basic investment calculations using GO functions and project structure.

        go run cmd/investcalc/investment_calculator.go

## Note
A program that asks the user for input, keeps the data in structs, and saves it in JSON format to a file with the help of struct tags and the encoding/json standard library.

        go run cmd/note/note.go

## Price Calculator
A program that illustrates a price calculator.  The program illustrates several useful features including Error Handling, Interfaces, JSON Marshalling, and reading/writing with files.

        go run cmd/pricecalc/pricecalc.go

## Structs
A program which illustrates basic structure usage.

        go run cmd/structs/structs.go

## Thread Greet
A program which illustrates basic usage of goroutines for a multi-threaded greet example.

        go run cmd/threadgreet/threadgreet.go

