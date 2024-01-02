/*
Sometimes you need the logic in your generic function to know something about the types it operates on.
The example we used in the first exercise didn't need to know anything about the types in the slice, so we used the built-in any constraint:

func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}

Constraints are just interfaces that allow us to write generics that only operate within the constraint of a given interface type.
In the example above, the any constraint is the same as the empty interface because it means the type in question can be anything.

Creating a custom constraint: Let's take a look at the example of a concat function.

It takes a slice of values and concatenates the values into a string.
This should work with any type that can represent itself as a string, even if it's not a string under the hood.
For example, a user struct can have a .String() that returns a string with the user's name and age.

type stringer interface {
    String() string
}

func concat[T stringer](vals []T) string {
    result := ""
    for _, val := range vals {
        // this is where the .String() method
        // is used. That's why we need a more specific
        // constraint instead of the any constraint
        result += val.String()
    }
    return result
}

*/

/*
package main

import (
	"errors"
	"fmt"
	"time"
)

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	newBalance := balance - newItem.GetCost()
	if newBalance < 0 {
		return nil, 0, errors.New("insufficient funds")
	}
	oldItems = append(oldItems, newItem)
	return oldItems, newBalance, nil
}

// don't edit below this line

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

func main() {
	test(subscription{
		userEmail: "john@example.com",
		startDate: time.Now().UTC(),
		interval:  "yearly",
	},
		[]subscription{},
		1000.00,
	)
	test(subscription{
		userEmail: "jane@example.com",
		startDate: time.Now().UTC(),
		interval:  "monthly",
	},
		[]subscription{
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7),
				interval:  "monthly",
			},
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7 * 52 * 2),
				interval:  "yearly",
			},
		},
		686.20,
	)
	test(oneTimeUsagePlan{
		userEmail:        "dillon@example.com",
		numEmailsAllowed: 5000,
	},
		[]oneTimeUsagePlan{},
		756.20,
	)
	test(oneTimeUsagePlan{
		userEmail:        "dalton@example.com",
		numEmailsAllowed: 100000,
	},
		[]oneTimeUsagePlan{
			{
				userEmail:        "dalton@example.com",
				numEmailsAllowed: 34200,
			},
		},
		32.20,
	)
}

func test[T lineItem](newItem T, oldItems []T, balance float64) {
	fmt.Println(" --- ")
	fmt.Printf("Charging customer for a '%s', current balance is %v...\n", newItem.GetName(), balance)
	newItems, newBalance, err := chargeForLineItem(newItem, oldItems, balance)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		return
	}
	fmt.Printf("New balance is: %v. Total number of line items is now %v\n", newBalance, len(newItems))
}
*/