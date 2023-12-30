/*
Packages

Every Go program is made up of packages.

You have probably noticed the package main at the top of all the programs you have been writing.

A package named "main" has an entrypoint at the main() function.
A main package is compiled into an executable program.

A package by any other name is a "library package". Libraries have no entry point.
Libraries simply export functionality that can be used by other packages. For example:

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

This program is an executable. It is a "main" package and imports from the fmt and math/rand library packages.
*/

/*
Package Naming Convention

By convention, a package's name is the same as the last element of its import path. For instance, the math/rand package comprises files that begin with:
package rand

That said, package names aren't required to match their import path. For example, I could write a new package with the path github.com/mailio/rand and name the package random:
package random

While the above is possible, it is discouraged for the sake of consistency.
One Package / Directory

A directory of Go code can have at most one package. All .go files in a single directory must all belong to the same package. If they don't an error will be thrown by the compiler. This is true for main and library packages alike.
*/