/*
Buffered channels in Go do not necessarily need to be run in goroutines.
However, they are often used in conjunction with goroutines to handle concurrent tasks.

A buffered channel allows the goroutine that is adding data to the channel to keep running and doing things, even if the goroutines reading from the channel are starting to fall behind a little bit.
This is because a buffered channel can hold a certain number of values in its buffer without needing a corresponding receiver for those values.
*/

/*
if you create a buffered channel with a capacity of 3, you can perform three send operations on this channel without blocking,
even if there’s no other goroutine ready to receive these values.
This is because the values are stored in the channel’s buffer until another goroutine receives them.
*/

/*
Buffered Channels

Channels can optionally be buffered.

Creating a channel with a buffer
- You can provide a buffer length as the second argument to make() to create a buffered channel:
- ch := make(chan int, 100)

Sending on a buffered channel only blocks when the buffer is full.

Receiving blocks only when the buffer is empty.
*/

/*
package main

import "fmt"

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

// TEST SUITE - Don't Touch Below This Line

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email:", email)
	}
}

func test(emails ...string) {
	fmt.Printf("Adding %v emails to queue...\n", len(emails))
	ch := addEmailsToQueue(emails)
	fmt.Println("Sending emails...")
	sendEmails(len(emails), ch)
	fmt.Println("==========================================")
}

func main() {
	test("Hello John, tell Kathy I said hi", "Whazzup bruther")
	test("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
	test("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")
}
*/