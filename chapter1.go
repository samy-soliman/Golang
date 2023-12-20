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
Go enforces strong and static typing, meaning variables can only have a single type.
A string variable like "hello world" can not be changed to an int

Contrast this with most interpreted languages, where the variable types are dynamic. Dynamic typing can lead to subtle bugs that are hard to detect.
With interpreted languages, the code must be run to catch syntax and type errors. (sometimes in production if you are unlucky 😨
*/

/*
Go programs are fairly lightweight. Each program includes a small amount of "extra" code that's included in the executable binary.
This extra code is called the Go Runtime. One of the purposes of the Go runtime is to cleanup unused memory at runtime.

In other words, the Go compiler includes a small amount of extra logic in every Go program to make it easier for developers to write code that's memory efficient.
*/

/*
Go's basic variable types are:

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

*/

/*
Variables are declared using the var keyword
ex: var number int
*/

/*
Inside a function (even the main function), the := short assignment statement can be used in place of a var declaration.
The := operator infers the type of the new variable based on the value.

ex: var empty string  is same as empty := ""

Outside of a function (in the global/package scope),
every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
*/

/*
Type Inference

To declare a variable without specifying an explicit type (either by using the := syntax or variable = expression syntax),
the variable's type is inferred from the value on the right hand side.
When the right hand side of the declaration is typed, the new variable is of that same type:

ex: var i int
j := i // j is also an int

However, when the right hand side is a literal value (an untyped numeric constant like 42 or 3.14),
the new variable will be an int, float64, or complex128 depending on its precision:

i := 42           // int
f := 3.14         // float64
g := 0.867 + 0.5i // complex128

*/

/*
We can declare multiple variables on the same line:

mileage, company := 80276, "Tesla"

// is the same as
mileage := 80276
company := "Tesla"

*/

/*
Type Sizes

Ints, uints, floats, and complex numbers all have type sizes.

int  int8  int16  int32  int64 // whole numbers

uint uint8 uint16 uint32 uint64 uintptr // positive whole numbers

float32 float64 // decimal numbers

complex64 complex128 // imaginary numbers (rare)

The size (8, 16, 32, 64, 128, etc) indicates how many bits in memory will be used to store the variable.
The default int and uint types are just aliases that refer to their respective 32 or 64 bit sizes depending on the environment of the user.
*/

/*
Some types can be converted the following way:

temperatureFloat := 88.26
temperatureInt := int64(temperatureFloat)

*/

/*
Constants are declared like variables but use the const keyword. Constants can't use the := short declaration syntax.
Constants can be character, string, boolean, or numeric values. They can not be more complex types like slices, maps and structs
*/

/*
Computed Constants

Constants must be known at compile time. More often than not they will be declared with a static value:

const myInt = 15

However, constants can be computed so long as the computation can happen at compile time.

For example, this is valid:

const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName

That said, you cannot declare a constant that can only be computed at run-time.
*/
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
