
/*
Functions

Functions in Go can take zero or more arguments.

To make Go code easier to read, the variable type comes after the variable name.

For example, the following function:

	func sub(x int, y int) int {
	  return x-y
	}

Accepts two integer parameters and returns another integer.
Here, func sub(x int, y int) int is known as the "function signature".
*/

/*
When multiple arguments are of the same type, and are next to each other in the function signature, the type only needs to be declared after the last argument.

	func addToDatabase(hp, damage int, name string) {
	  // ?
	}
*/

/*
example of go functions systax
f func(func(int,int) int, int) int
A function named 'f' that takes a function and an int as arguments and returns an int
*/

/*
Until now Variables in Go are passed by value
*/

/*

A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore: _

For example:

func getPoint() (x int, y int) {
  return 3, 4
}

// ignore y value
x, _ := getPoint()

This is crucial to understand because the Go compiler will throw an error if you have unused variable declarations in your code,
so you need to ignore anything you don't intend to use.
*/

/*
Return values may be given names, and if they are, then they are treated the same as if they were new variables defined at the top of the function.

Named return values are best thought of as a way to document the purpose of the returned values.

According to the tour of go:

    A return statement without arguments returns the named return values. This is known as a "naked" return. Naked return statements should be used only in short functions. They can harm readability in longer functions.

func getCoords() (x, y int){
  // x and y are initialized with zero values

  return // automatically returns x and y
}

Is the same as:

func getCoords() (int, int){
  var x int
  var y int
  return x, y
}

*/

/*
Even though a function has named return values, we can still explicitly return values if we want to.

func getCoords() (x, y int){
  return x, y // this is explicit
}

Using this explicit pattern we can even overwrite the return values:

func getCoords() (x, y int){
  return 5, 6 // this is explicit, x and y are NOT returned
}

Otherwise, if we want to return the values defined in the function signature we can just use a naked return (blank return):

func getCoords() (x, y int){
  return // implicitly returns x and y

*/

/*
Early Returns

Go supports the ability to return early from a function. This is a powerful feature that can clean up code, especially when used as guard clauses.

Guard Clauses leverage the ability to return early from a function (or continue through a loop) to make nested conditionals one-dimensional. Instead of using if/else chains, we just return early from the function at the end of each conditional block.

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend/divisor, nil
}

*/