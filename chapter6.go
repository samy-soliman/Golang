// The Error Interface

/*
Go programs express errors with error values. An Error is any type that implements the simple built-in error interface:

type error interface {
    Error() string
}

When something can go wrong in a function, that function should return an error as its last return value. Any code that calls a function that can return an error should handle errors by testing whether the error is nil.

// Atoi converts a stringified number to an integer
i, err := strconv.Atoi("42b")
if err != nil {
    fmt.Println("couldn't convert:", err)
    // because "42b" isn't a valid integer, we print:
    // couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
    // Note:
    // 'parsing "42b": invalid syntax' is returned by the .Error() method
    return
}
// if we get here, then
// i was converted successfully

A nil error denotes success; a non-nil error denotes failure.
*/

// code example 1
/*
package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	cost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}

	costSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}
	return costSpouse + cost, nil
}

// don't edit below this line

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

func test(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("========")
	fmt.Println("Message for customer:", msgToCustomer)
	fmt.Println("Message for spouse:", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

func main() {
	test(
		"Thanks for coming in to our flower shop today!", "We hope you enjoyed your gift.")
	test(
		"Thanks for joining us!",
		"Have a good day.",
	)
	test(
		"Thank you.",
		"Enjoy!",
	)
	test(
		"We loved having you in!",
		"We hope the rest of your evening is absolutely fantastic.",
	)
}
*/

// code example 2
/*
package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", de.dividend)
}

// don't edit below this line

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		// We convert the `divideError` struct to an `error` type by returning it
		// as an error. As an error type, when it's printed its default value
		// will be the result of the Error() method
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func test(dividend, divisor float64) {
	defer fmt.Println("====================================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", dividend, divisor)
	quotient, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}

func main() {
	test(10, 0)
	test(10, 2)
	test(15, 30)
	test(6, 3)
}
*/

// code example 3
/*
package main

import (
	"fmt"
)

type car struct {
	Make   string
	Model  string
	Height int
	Width  int
}

func (sCar car) Error() string {
	return fmt.Sprintf("Error")
}

func getCar(newCar string) (string, error) {
	if newCar == "toyota" {
		return "empty", car{Model: newCar, Height: 10}
	}
	return "Congratulations", nil
}

func main() {
	getCarResult, err := getCar("toyota")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("getCarResult: %v \n", getCarResult)

	getCarResult2, err2 := getCar("nessan")
	if err2 == nil {
		fmt.Println(err2)
	}
	fmt.Printf("getCarResult2: %v \n", getCarResult2)
}
*/
