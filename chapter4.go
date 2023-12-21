/*
Structs in Go

We use structs in Go to represent structured data. It's often convenient to group different types of variables together. For example, if we want to represent a car we could do the following:

type car struct {
  Make string
  Model string
  Height int
  Width int
}

This creates a new struct type called car. All cars have a Make, Model, Height and Width.

In Go, you will often use a struct to represent information that you would have used a dictionary for in Python

The fields of a struct can be accessed using the dot . operator.

myCar := car{}
myCar.FrontWheel.Radius = 5

when you declare a struct, the fields within it are initialized to their zero values. These zero values depend on the data type of each field
*/

/*
An anonymous struct is just like a normal struct, but it is defined without a name and therefore cannot be referenced elsewhere in the code.

To create an anonymous struct, just instantiate the instance immediately using a second pair of brackets after declaring the type:

	myCar := struct {
	  Make string
	  Model string
	} {

	  Make: "tesla",
	  Model: "model 3",
	}

You can even nest anonymous structs as fields within other structs:

	type car struct {
	  Make string
	  Model string
	  Height int
	  Width int
	  // Wheel is a field containing an anonymous struct
	  Wheel struct {
	    Radius int
	    Material string
	  }
	}

When should you use an anonymous struct?

In general, prefer named structs. Named structs make it easier to read and understand your code, and they have the nice side-effect of being reusable.

If a struct is only meant to be used once, then it makes sense to declare it in such a way that developers down the road won’t be tempted to accidentally use it again.
*/

/*
While Go is not object-oriented, it does support methods that can be defined on structs. Methods are just functions that have a receiver. A receiver is a special parameter that syntactically goes before the name of the function.

type rect struct {
  width int
  height int
}

// area has a receiver of (r rect)
func (r rect) area() int {
  return r.width * r.height
}

r := rect{
  width: 5,
  height: 10,
}

fmt.Println(r.area())
// prints 50

A receiver is just a special kind of function parameter.
*/

package main

import (
	"fmt"
)

/*
	 code1

		type messageToSend struct {
			message   string
			sender    user
			recipient user
		}

		type user struct {
			name   string
			number int
		}

		func canSendMessage(mToSend messageToSend) bool {
			if len(mToSend.sender.name) > 0 && len(mToSend.recipient.name) > 0 && mToSend.sender.number != 0 && mToSend.recipient.number != 0 {
				return true
			}
			return false
		}

// don't touch below this line

	func test(mToSend messageToSend) {
		fmt.Printf(`sending "%s" from %s (%v) to %s (%v)...`,
			mToSend.message,
			mToSend.sender.name,
			mToSend.sender.number,
			mToSend.recipient.name,
			mToSend.recipient.number,
		)
		fmt.Println()
		if canSendMessage(mToSend) {
			fmt.Println("...sent!")
		} else {
			fmt.Println("...can't send message")
		}
		fmt.Println("====================================")
	}
*/
func main() {
	/* code1

	test(messageToSend{
		message: "you have an appointment tommorow",
		sender: user{
			name:   "Brenda Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have an event tommorow",
		sender: user{
			number: 16545550987,
		},
		recipient: user{
			name:   "Suzie Sall",
			number: 0,
		},
	})
	test(messageToSend{
		message: "you have an party tommorow",
		sender: user{
			name:   "Njorn Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have a birthday tommorow",
		sender: user{
			name:   "Eli Halafax",
			number: 0,
		},
		recipient: user{
			name:   "Whitaker Sue",
			number: 19035558973,
		},
	})
	*/

	type car struct {
		make  string
		model string
	}

	type truck struct {
		// "car" is embedded, so the definition of a
		// "truck" now also additionally contains all
		// of the fields of the car struct
		car
		bedSize int
	}
	// created an instance of truck called lanesTruck
	lanesTruck := truck{
		bedSize: 10,
		car: car{
			make:  "toyota",
			model: "camry",
		},
	}

	fmt.Println(lanesTruck.bedSize)

	// embedded fields promoted to the top-level
	// instead of lanesTruck.car.make
	fmt.Println(lanesTruck.make)
	fmt.Println(lanesTruck.model)

}
