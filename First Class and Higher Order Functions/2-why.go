/*
Why First-class and Higher-Order Functions?

At first, it may seem like dynamically creating functions and passing them around as variables adds unnecessary complexity.
Most of the time you would be right. There are cases however when functions as values make a lot of sense. Some of these include:

    HTTP API handlers
    Pub/Sub handlers
    Onclick callbacks

Any time you need to run custom code at a time in the future, functions as values might make sense.
Definition: First-class Functions

A first-class function is a function that can be treated like any other value.
Go supports first-class functions. A function's type is dependent on the types of its parameters and return values.
For example, these are different function types:

func() int
func(string) int

Definition: Higher-Order Functions

A higher-order function is a function that takes a function as an argument or returns a function as a return value.
Go supports higher-order functions. For example, this function takes a function as an argument:

func aggregate(a, b, c int, arithmetic func(int, int) int) int
*/