// loops

/*
The basic loop in Go is written:

for INITIAL; CONDITION; AFTER{
  // do something
}

INITIAL is run once at the beginning of the loop and can create
variables within the scope of the loop.

CONDITION is checked before each iteration. If the condition doesn't pass
then the loop breaks.

AFTER is run after each iteration.

For example:

for i := 0; i < 10; i++ {
  fmt.Println(i)
}
// Prints 0 through 9
*/

/*
Loops in Go can omit sections of a for loop. For example, the CONDITION (middle part) can be omitted which causes the loop to run forever.

for INITIAL; ; AFTER {
  // do something forever
}
*/

/*
Most programming languages have a concept of a while loop. Because Go allows for the omission of sections of a for loop, a while loop is just a for loop that only has a CONDITION.

for CONDITION {
  // do some stuff while CONDITION is true
}

For example:

plantHeight := 1
for plantHeight < 5 {
  fmt.Println("still growing! current height:", plantHeight)
  plantHeight++
}
fmt.Println("plant has grown to ", plantHeight, "inches")

*/

/*
Go supports the standard modulo operator:

7 % 3 // 1

Logical AND operator:

true && false // false
true && true // true

Logical OR operator:

true || false // true
false || false // false

*/

/*
The continue keyword stops the current iteration of a loop and continues to the next iteration. continue is a powerful way to use the "guard clause" pattern within loops.

for i := 0; i < 10; i++ {
  if i % 2 == 0 {
    continue
  }
  fmt.Println(i)
}
// 1
// 3
// 5
// 7
// 9

The break keyword stops the current iteration of a loop and exits the loop.

for i := 0; i < 10; i++ {
  if i == 5 {
    break
  }
  fmt.Println(i)
}
// 0
// 1
// 2
// 3
// 4

*/