/*
1- A Go program is a combination of packages
2- A package is composed of one or more source files. Into those source files, the Go programmer declares :
   constant, variables, functions, types and methods

3- The package main is often composed of an individual file.
4- The function main is the program entry point.
5- In a Go program, you will also find a file named go.mod.
*/

/*
The following snippet is an example of a source file that belongs to a package occupancy :

// package-imports/occupancy/occupancy.go
package occupancy

const highLimit = 70.0
const mediumLimit = 20.0

func level(occupancyRate float32) string {
    if occupancyRate > highLimit {
        return "High"
    } else if occupancyRate > mediumLimit {
        return "Medium"
    } else {
        return "Low"
    }
}

func rate(roomsOccupied int, totalRooms int) float32 {
    return (float32(roomsOccupied) / float32(totalRooms)) * 100
}

1- At the top of the source file, we find the package clause in our example, it’s : package occupancy
The package clause is the first line of each source file. It defines the name of the current package.

2- Then the set of imports declarations.In this section of the source file, we define all the other packages that we want to use in this package. The package occupancy does not import other packages. Let’s take another example: here is a source file from the package room :
// package-imports/import-declaration/room/room.go
package room

import "fmt"

func printDetails(roomNumber, size, nights int) {
    fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
}

Here we import one package : "fmt" wich is part of the standard library

3- After the imports declarations, we find the most important part, the source code of the package. This is where you can declare variables, constants, functions, types, and methods.
*/

/*
We must group the source files of a package into a single directory. The directory must have the same name as the package. For instance, the source files of the baz package must be stored into the baz folder
*/

/*
A Go program starts by initializing the main package and then run the function main from that package.
The main package is where your program starts doing what it was built to do.

Here is an example of a main package :

// package-imports/main-package/main.go
package main

import "fmt"

func init() {
    fmt.Println("launch initialization")
}

func main() {
    fmt.Println("launch the program !")
}

This program has an init function. This function can hold all the initialization tasks necessary for your program to run correctly.

The program also defines a main function. Both functions do not have a return type.

The main function will be executed after all the initialization tasks have been performed. In this function, you usually call other packages and implement your program logic.
*/

/*
1- One main package per project?
This is not always the case, but a project can have several main packages and thus several main functions.
Usually, different main packages are present in large projects. Here are some common examples :
    A main package for launching the web server of the application
    Another main package to run scheduled database maintenance
    Another one that has been developed for a specific punctual intervention...
For instance, the Kubernetes (one of the most starred Go projects) holds at the time of speaking 20 main packages.

2- Should I name the file that holds my main package main.go?
No, it’s not an obligation. You can do that if you have just one main function in your project. But often, it’s good practice to give it an informational name.

3- Our main packages stored in a specific directory
Again nothing is written in the Go specification about the folder that should contain the main package.
There is still a strong usage main packages should live inside a cmd/ folder at the root of the repository.
*/

/*
1- Here is an example of a go.mod file :

module maximilien-andile.com/myProg
go 1.13

2- The first line is composed of the word module followed by the module path.

a module is “a collection of Go packages stored in a file tree with a go.mod file at its root”
When we add a package main to a module, we make it a program that can be compiled and executed.

We need to define a module path for our module (our program). The path is the unique location of the module.
Here is an example for famous Go program >> Hashicorp Vault,2 : module github.com/hashicorp/vault

"github.com/hashicorp/vault" is an URL to a Github repository. A Github repository is a folder that contains the source code of a project. This folder can be publicly accessible (i.e., by anyone who has the link) or private.
In the latter case, the code is only accessible to authorized developers. If you have access to a repository that has a go.mod inside, you can import it in your project.

Note that your program will not be automatically shared or exposed by Go if you choose a remote path! For testing purpose, you can choose a path that is not an URL :
module thisIsATest

go 1.13
*/

/*
Expected Go version
The next line of the go.mod file declares Go’s expected version. In this program, version 1.13 of go is expected
*/

/*
To define a module, we will use Gauthier’s standard definition.
A module is a chunk of code that :
    Performs a specific task
    With inputs and outputs well defined
    That can be tested independently

A large project that requires to perform several tasks can be split into different tasks, each one living inside a module with a well-defined API and that can be tested independently from other systems.
*/

/*
How to decompose a program into modules

Decomposition of a system into modules can be difficult. There is no unique answer to this issue. Here is some advice that is based on experience and my readings :

    Do not create a module for every single task of a system; you will end up with too many modules with a size that is too small.

    Instead, create modules that group functionalities that are close, for instance, a module for handling the database queries, a module to handle logging.

    Inside your modules, the code is usually tightly coupled, meaning that the different parts of your module are strongly linked.

    Between two modules, you should enforce loose coupling. Each module is treated as a component, and each component do not need another component to work. Modules should be independent. It allows you to replace one module with another one without touching the other modules.
*/

/*
1- Go packages are a tool that you can use to write modular programs.
In a Go package We can group source files with functions (or methods) related to a particular functionality.
2- Go modules are a way of grouping packages together to form an application or a library.
*/

/*
1- How to download code from distant packages
To download a package, you will use the go get command inside your terminal.

2- An import declaration declares that we will use some external code in the current package.
An import declaration can have different forms :

2.1- Standard import >>> import "gitlab.com/loir402/foo"
Then in the file, you can access the exported element of the imported package with the syntax : foo.MyFunction()
2.2- Multiple package import
You rarely have just one import. Often you have several packages to import :
import (
   "fmt"
   "maximilien-andile.com/application2/foo"
   "maximilien-andile.com/application2/bar"
)
Here we import three packages. Here are their import paths :
    fmt
    maximilien-andile.com/application2/foo
    maximilien-andile.com/application2/bar
2.3- Import with an explicit package name (alias)
You can also give your package an alias (a sort of surname for the package) to call its exported elements with this alias. For instance :
import bar "gitlab.com/loir402/foo"
bar.MyFunction()
2.4- Import with a dot
import ."gitlab.com/loir402/foo"
With this syntax, all the functions, types, variables, and constants will be declared in the current package. As a consequence, all the exported identifiers of the package can be used without any qualifier. I do not recommend using this because it can become confusing.
2.5- The blank import
With a blank import, just the package’s init function will be called. A blank import can be specified with the following syntax :
import _ "gitlab.com/loir402/foo"
*/

/*
The internal directory

When you create a package and when your program is available on a code-sharing site (like Github, Gitlab ...), other developers will be able to use your package.
This is a great responsibility for you. Other programs will rely on your code. It means that when you change a function or the name of an Exported identifier, other code might break.

To forbid the import of packages, you can put them into a directory called “internal”.

In the 4-image , you can see an example directory structure with an internal directory :

    cmd/main.go contain the main package and the main function

    internal/booking is the directory of the package booking.

    internal/booking/booking.go is a source file of the package booking.

    Any exported identifiers inside the booking package are accessible into the current program (ie. we can call a function from booking package into package main)

    BUT, other developers will not be able to use it in their program.
*/

/*
Key takeaways
    A package is a group of source files that live in the same directory
    A name identifies a package
    Identifiers with a first letter uppercase are exported.
    An exported identifier can be used in any other packages.
    Packages that live inside an internal directory can be used inside packages of your module. However, they cannot be used by other modules.
    A module is “a collection of Go packages stored in a file tree with a go.mod file at its root”
*/
