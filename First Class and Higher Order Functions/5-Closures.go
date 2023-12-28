/*Function literals

A function literal represents an anonymous function. Function literals cannot declare type parameters.

FunctionLit = "func" Signature FunctionBody .

func(a, b int, z float64) bool { return a*b < int(z) }

A function literal can be assigned to a variable or invoked directly.

f := func(x, y int) int { return x + y }
func(ch chan int) { ch <- ACK }(replyChan)

Function literals are closures: they may refer to variables defined in a surrounding function.
Those variables are then shared between the surrounding function and the function literal, and they survive as long as they are accessible.
*/

/*
Go functions may be closures. A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the adder function returns a closure. Each closure is bound to its own sum variable.

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
*/

// another example
/*
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive
}
*/
