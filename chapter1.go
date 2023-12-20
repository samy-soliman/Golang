/*
Generally speaking, compiled languages run much faster than interpreted languages, and Go is no exception.
Go is one of the fastest programming languages, beating JavaScript, Python, and Ruby handily in most benchmarks.
However, Go code doesn't run quite as fast as its compiled Rust and C counterparts. That said, it compiles much
faster than they do, which makes the developer experience super productive
*/

/*
Compiled programs can be run without access to the original source code, and without access to a compiler.
For example, when your browser executes the code you write in this course, it doesn't use the original code, just the compiled result.
Note how this is different than interpreted languages like Python and JavaScript.
With Python and JavaScript the code is interpreted at runtime by a separate program known as the "interpreter".
Distributing code for users to run can be a pain because they need to have an interpreter installed,
and they need access to the original source code.
*/

/*
Go is generally faster and more lightweight than interpreted or VM-powered languages like:
    Python
    JavaScript
    PHP
    Ruby
    Java
However, in terms of execution speed, Go does lag behind some other compiled languages like:
    C
    C++
    Rust
Go is a bit slower mostly due to its automated memory management, also known as the "Go runtime".
Slightly slower speed is the price we pay for memory safety and simple syntax!
*/

/*
Computers need machine code, they don't understand English or even Go. We need to convert our high-level (Go) code
into machine language, which is really just a set of instructions that some specific hardware can understand.
In your case, your CPU.

The Go compiler's job is to take Go code and produce machine code. On Windows, that would be a .exe file. On Mac or Linux,
it would be any executable file. The code you write in your browser is being compiled for you on our servers,
then the machine code is executed in your browser as Web Assembly.
*/

/*
Computers need machine code
A computer's CPU only understands its own instruction set, which we call "machine code".
Instructions are basic math operations like addition, subtraction, multiplication, and the ability to save data temporarily.
For example, an ARM processor uses the ADD instruction when supplied with the number 0100 in binary.
*/

/*
the structure of a Go program
1- package main lets the Go compiler know that we want this code to compile and run as a standalone program,
as opposed to being a library that's imported by other programs.
2- func main() defines the main function. main is the name of the function that acts as the entry point for a Go program.
*/

/*

 */

package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
