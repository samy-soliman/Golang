
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

/*
fmt.Printf - Prints a formatted string to standard out.
fmt.Sprintf() - Returns the formatted string
*/

/*
Interpolate the default representation

// The %v variant prints the Go syntax representation of a value.
s := fmt.Sprintf("I am %v years old", 10)
// I am 10 years old

// Interpolate a string
s := fmt.Sprintf("I am %s years old", "way too many")
// I am way too many years old

// Interpolate an integer in decimal form
s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old

// Interpolate a decimal
s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old
// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old


*/

/*
Conditionals

if statements in Go don't use parentheses around the condition:
else if and else are supported as you would expect:

if height > 6 {
    fmt.Println("You are super tall!")
} else if height > 4 {
    fmt.Println("You are tall enough!")
} else {
    fmt.Println("You are not tall enough!")
}

*/

/*
if conditional can have an "initial" statement. The variable(s) created in the initial statement are only defined within the scope of the if body.

if INITIAL_STATEMENT; CONDITION {
}

Why would I use this?

This is just some syntactic sugar that Go offers to shorten up code in some cases. For example, instead of writing:

length := getLength(email)

	if length < 1 {
	    fmt.Println("Email is invalid")
	}

We can do:

	if length := getLength(email); length < 1 {
	    fmt.Println("Email is invalid")
	}
*/
