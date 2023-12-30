/*
The go mod init command creates and initializes a new module.
On execution of the command, a go.mod file is created to manage all project dependencies and versions:
declare your module's name:

go mod init {REMOTE}/{USERNAME}/hellogo

Where {REMOTE} is your preferred remote source provider (i.e. github.com) and {USERNAME} is your Git username.
*/

/*
go run main.go

The go run command is used to quickly compile and run a Go package.
The compiled binary is not saved in your working directory.
*/

/*
go build compiles go code into an executable program
Build an executable

Ensure you are in your hellogo repo, then run:

go build

Run the new program:

./hellogo
*/

/*
go get

The go get command is used to add, update, or remove packages as well as their dependencies.
It also updates the go.mod file to reflect the new dependencies.

The syntax for the go get command is the words go and get, followed by the package path:

> go get [package_path]

On execution of the command, the go.mod and go.sum files automatically update to reflect the specified version:
*/

/*
The go install command is used to build and install executables

go install

Go has installed the hellogo program globally. Run it with:

hellogo
*/

/*
go mod tidy

On execution of the go mod tidy command, the import statements of all project files are analyzed and the go.mod file is updated to include only the packages used in your code:

> go mod tidy
*/

/*
go list

The go list command lists packages and modules contained in your project.
The go list command also accepts flags and arguments such as -m and -m all, which streamlines the command to packages in the current module:

> go list
*/

/*
go mod vendor

The go mod vendor command creates a self-contained project with all its dependencies stored locally.
On execution of this command, a vendor folder containing copies or clones of all the packages used in the project is created:

> go mod vendor
*/

/*
go mod verify

The go mod verify command verifies the integrity of the packages and dependencies by comparing the checksums in the go.sum file:

> go mode verify
*/

/*
As you begin to integrate new dependencies into your project, the go.mod dynamically updates to keep track of these dependencies and their corresponding versions.
So introduce a new dependency in your project to observe changes in the go.mod file
*/

/*
The go.sum file is equally generated on the addition of your first dependency. This file serves as a comprehensive record, capturing not only the dependencies themselves but also crucial information like version identifiers and cryptographic checksums:
*/

/*
It’s also possible to precisely define or modify the preferred version of a dependency for integration into your project.
This can be achieved by indicating the intended version within the scope of the go get command.
The execution of this command triggers an update of the require directive in the go.mod file, solidifying the dependency’s version.

Alternatively, you can manually edit the go.mod file; however, it’s recommended to use the go get command approach.
*/