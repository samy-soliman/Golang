/*

Interfaces are collections of method signatures. A type "implements" an interface if it has all of the methods of the given interface defined on it.

In the following example, a "shape" must be able to return its area and perimeter. Both rect and circle fulfill the interface.

type shape interface {
  area() float64
  perimeter() float64
}

type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}

type circle struct {
    radius float64
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

When a type implements an interface, it can then be used as the interface type.

*/

/*
Interfaces are implemented implicitly.
A type never declares that it implements a given interface. If an interface exists and a type has the proper methods defined, then the type automatically fulfills that interface.
A type implements an interface by implementing its methods. Unlike in many other languages, there is no explicit declaration of intent, there is no "implements" keyword.
A type can implement any number of interfaces in Go
*/

/*
Name Your Interface Arguments

Consider the following interface:

type Copier interface {
  Copy(string, string) int
}

Based on the code alone, can you deduce what kinds of strings you should pass into the Copy function?

We know the function signature expects 2 string types, but what are they? Filenames? URLs? Raw string data? For that matter, what the heck is that int that's being returned?

Let's add some named arguments and return data to make it more clear.

type Copier interface {
  Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

Much better. We can see what the expectations are now. The first argument is the sourceFile, the second argument is the destinationFile, and bytesCopied, an integer, is returned.
*/

/*
Interfaces in Go are like contracts. They define a set of methods that a type must implement.
When you create a value of an interface type, it can hold any concrete type that satisfies the interface.
Interface assertion is a way to access the actual value stored inside an interface.
Imagine you have a box (the interface) that can hold different objects (concrete types). You want to peek inside the box and see what’s really there.
That’s what interface assertion does: it lets you open the box and get the actual object (concrete value) out.

// Define an interface
type Shape interface {
    Area() float64
}

// Create a circle
type Circle struct {
    Radius float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    // Create a Shape interface value
    var s Shape
    s = Circle{Radius: 5}

    // Interface assertion: get the actual Circle value
    circle, ok := s.(Circle)
    if ok {
        fmt.Printf("Circle area: %.2f\n", circle.Area())
    } else {
        fmt.Println("Not a Circle!")
    }
}

another real world example
Imagine you’re building a simple drawing application. Users can create different shapes (like circles, rectangles, and triangles) and perform actions on them. Here’s how you can use interface assertions:

Define the Shape Interface:

First, define an interface called Shape with a method Area() that calculates the area of the shape.
All shapes (circles, rectangles, etc.) will implement this interface.
Create Concrete Shape Types:

Implement concrete types for different shapes (e.g., Circle, Rectangle, Triangle).
Each type must have an Area() method that computes its area.
Store Shapes in a Slice:

In your drawing app, users can create shapes and add them to a list.
You can use a slice of the Shape interface to store different shapes.
Calculate Total Area:

To calculate the total area of all shapes, iterate through the slice and sum up the areas.
Use interface assertions to access the actual shape values and call their Area() methods.

package main

import (
    "fmt"
    "math"
)

// Shape defines the interface for all shapes
type Shape interface {
    Area() float64
}

// Circle implements the Shape interface
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Rectangle implements the Shape interface
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    // Create a slice of Shape (interface)
    shapes := []Shape{
        Circle{Radius: 5},
        Rectangle{Width: 3, Height: 4},
    }

    // Calculate total area
    totalArea := 0.0
    for _, shape := range shapes {
        // Interface assertion: get the actual shape value
        switch s := shape.(type) {
        case Circle:
            fmt.Printf("Circle area: %.2f\n", s.Area())
        case Rectangle:
            fmt.Printf("Rectangle area: %.2f\n", s.Area())
        }
        totalArea += shape.Area()
    }

    fmt.Printf("Total area of all shapes: %.2f\n", totalArea)
}

*/

/*
A type switch makes it easy to do several type assertions in a series.

A type switch is similar to a regular switch statement, but the cases specify types instead of values.

func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func main() {
	printNumericValue(1)
	// prints "int"

	printNumericValue("1")
	// prints "string"

	printNumericValue(struct{}{})
	// prints "struct {}"
}

fmt.Printf("%T\n", v) prints the type of a variable.
*/

/*
Writing clean interfaces is hard. Frankly, anytime you’re dealing with abstractions in code, the simple can become complex very quickly if you’re not careful. Let’s go over some rules of thumb for keeping interfaces clean.
1. Keep Interfaces Small

If there is only one piece of advice that you take away from this article, make it this: keep interfaces small! Interfaces are meant to define the minimal behavior necessary to accurately represent an idea or concept.

Here is an example from the standard HTTP package of a larger interface that’s a good example of defining minimal behavior:

type File interface {
    io.Closer
    io.Reader
    io.Seeker
    Readdir(count int) ([]os.FileInfo, error)
    Stat() (os.FileInfo, error)
}

Any type that satisfies the interface’s behaviors can be considered by the HTTP package as a File. This is convenient because the HTTP package doesn’t need to know if it’s dealing with a file on disk, a network buffer, or a simple []byte.
2. Interfaces Should Have No Knowledge of Satisfying Types

An interface should define what is necessary for other types to classify as a member of that interface. They shouldn’t be aware of any types that happen to satisfy the interface at design time.

For example, let’s assume we are building an interface to describe the components necessary to define a car.

type car interface {
	Color() string
	Speed() int
	IsFiretruck() bool
}

Color() and Speed() make perfect sense, they are methods confined to the scope of a car. IsFiretruck() is an anti-pattern. We are forcing all cars to declare whether or not they are firetrucks. For this pattern to make any amount of sense, we would need a whole list of possible subtypes. IsPickup(), IsSedan(), IsTank()… where does it end??

Instead, the developer should have relied on the native functionality of type assertion to derive the underlying type when given an instance of the car interface. Or, if a sub-interface is needed, it can be defined as:

type firetruck interface {
	car
	HoseLength() int
}

Which inherits the required methods from car and adds one additional required method to make the car a firetruck.
3. Interfaces Are Not Classes
    Interfaces are not classes, they are slimmer.
    Interfaces don’t have constructors or deconstructors that require that data is created or destroyed.
    Interfaces aren’t hierarchical by nature, though there is syntactic sugar to create interfaces that happen to be supersets of other interfaces.
    Interfaces define function signatures, but not underlying behavior. Making an interface often won’t DRY up your code in regards to struct methods. For example, if five types satisfy the fmt.Stringer interface, they all need their own version of the String() function.
*/

package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

// struct method
func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func test(m message) {
	sendMessage(m)
	fmt.Println("====================================\n")
}

func main() {
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "Bill Deer",
		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
	})
}
