/*
Sometimes we have a single goroutine listening to multiple channels and want to process data in the order it comes through each channel.

A select statement is used to listen to multiple channels at the same time.
It is similar to a switch statement but for channels.

select {
  case i, ok := <- chInts:
    fmt.Println(i)
  case s, ok := <- chStrings:
    fmt.Println(s)
}

The first channel with a value ready to be received will fire and its body will execute.
If multiple channels are ready at the same time one is chosen randomly.
The ok variable in the example above refers to whether or not the channel has been closed by the sender yet.
*/

/*
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case email, ok := <-chEmails:
			if !ok {
				return
			}
			logEmail(email)
		case sms, ok := <-chSms:
			if !ok {
				return
			}
			logSms(sms)
		}
	}
}

// TEST SUITE - Don't touch below this line

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

func test(sms []string, emails []string) {
	fmt.Println("Starting...")

	chSms, chEmails := sendToLogger(sms, emails)

	logMessages(chEmails, chSms)
	fmt.Println("===============================")
}

func main() {
	rand.Seed(0)
	test(
		[]string{
			"hi friend",
			"What's going on?",
			"Welcome to the business",
			"I'll pay you to be my friend",
		},
		[]string{
			"Will you make your appointment?",
			"Let's be friends",
			"What are you doing?",
			"I can't believe you've done this.",
		},
	)
	test(
		[]string{
			"this song slaps hard",
			"yooo hoooo",
			"i'm a big fan",
		},
		[]string{
			"What do you think of this song?",
			"I hate this band",
			"Can you believe this song?",
		},
	)
}

func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
	chSms = make(chan string)
	chEmails = make(chan string)
	go func() {
		for i := 0; i < len(sms) && i < len(emails); i++ {
			done := make(chan struct{})
			s := sms[i]
			e := emails[i]
			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
			go func() {
				time.Sleep(t1)
				chSms <- s
				done <- struct{}{}
			}()
			go func() {
				time.Sleep(t2)
				chEmails <- e
				done <- struct{}{}
			}()
			<-done
			<-done
			time.Sleep(10 * time.Millisecond)
		}
		close(chSms)
		close(chEmails)
	}()
	return chSms, chEmails
}
*/

/*
The default case in a select statement executes immediately if no other channel has a value ready.
A default case stops the select statement from blocking.

select {
  case v := <-ch:
    // use v
  default:
    // receiving from ch would block
    // so do something else
}

Tickers

    time.Tick() is a standard library function that returns a channel that sends a value on a given interval.
    time.After() sends a value once after the duration has passed.
    time.Sleep() blocks the current goroutine for the specified amount of time.

Read-only Channels
A channel can be marked as read-only by casting it from a chan to a <-chan type. For example:

func main(){
  ch := make(chan int)
  readCh(ch)
}

func readCh(ch <-chan int) {
  // ch can only be read from
  // in this function
}

Write-only Channels
The same goes for write-only channels, but the arrow's position moves.

func writeCh(ch chan<- int) {
  // ch can only be written to
  // in this function
}
*/

/*
package main

import (
	"fmt"
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot()
		case <-saveAfter:
			saveSnapshot()
			return
		default:
			waitForData()
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// TEST SUITE - Don't touch below this line

func takeSnapshot() {
	fmt.Println("Taking a backup snapshot...")
}

func saveSnapshot() {
	fmt.Println("All backups saved!")
}

func waitForData() {
	fmt.Println("Nothing to do, waiting...")
}

func test() {
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("===========================")
}

func main() {
	test()
}
*/