/*

In Golang, or Go, channels are a means through which different goroutines communicate.

Think of them as pipes through which you can connect with different concurrent goroutines. The communication is bidirectional by default, meaning that you can send and receive values from the same channel.

Moreover, by default, channels send and receive until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

Syntax....
Make a channel: make(chan [value-type]), where [value-type] is the data type of the values to send and receive, e.g., int.

Send and receive values: channel <- and <- channel,
where <- is the channel operator.

Close a channel: close(channel).
After closing, no value will be sent to the channel.

Both sending and receiving are blocking operations by default.

channels are bidirectional communication
 conceptually the same as a two ended pipe:
  write at one end and read out to the other end
  this is also called sending and recieving

utilizing channels enables goroutine to communicate:
 can send/ receive messages or computational results
channel ends can be duplicated across goroutines

messages are serialized sent one after the other and read one after the other
we can have multiple reading ends and single sending or writing end
*/

//creation and usage
/*
channel := make(chan int)   ..... creating channel that can communicate using only integer

//send to channel
go func(){channel <- 1}()      ..... we are sending 1 to the channel
go func(){channel <- 2}()	   ..... we are sending 2 to the channel
go func(){channel <- 3}()

each one of these sends to the channel occuring with the go routine that's because this (channel := make(chan int)) channel is unbuffered
that means every message sent in needs to be received out
whenever we send a message to unbuffered channel the channel operation is going to wait , this is blocking , this blocks until something reads from the channel
thats why we have spawn it using goroutines , goroutines can all be blocking , and our main thread can be executing


above we have 3 goroutines sending in to the channel
below in our main thread we are receiving from the channel
first := <-channel
second := <-channel
third := <-channel

// above we are assigning the read data

channels can be buffered or unbuffered
 unbuffered channels are will block when sending until a reader is available
 buffered channels have specified capacity
  can send messages upto the capacity , even without a reader
messages on a channel are FIFO ordering
*/

//buffered channel
/*
channel := make(chan int, 2)

channel <- 1
channel <- 2   ..... this has space for only 2 but if we want for to add more then we need to use goroutines

go func(){channel <- 3 }()    above two channel will be send even without a reader but this one will block as there are no readers yet

first := <- channel
second := <- channel
third := <- channel

fmt.Println(first , second , third) first and second will always be 1 and 2 and third one will be 3 as there is only one goroutine

*/

//goroutines unidirectional
/*
suppose we have a main thread for sneding to channel and we have 1, 2, 3 goroutines then main thread will only send the data to these goroutines and these goroutines cannot send data back to main thread
*/

//goroutines : control channel
/*
so in this case main thread can send the data to goroutines and goroutines can have seperate channels to send data back to main thread
*/

/*
channel selection
the select keyword lets you work with multiple , potentially blocking , channels
send / receive attempts are made , regardless of blocking status

example
one := make(chan int) ...unbuffered channels
two := make(chan int)

for{
	select {
	case o := <- one
		fmt.Println("one",o)  ....here we are trying to read from this channel since nothing is written to the channel yet so this would normally block and will move to next
	case t := <- two
		fmt.Println("two",t)
	default:
		fmt.Println("no data to receive")
		time.Sleep(50*time.Millisecond)
	}
}
*/

//timeouts
/*
the time package can be combined with select to create timouts
one := make(chan int) ...unbuffered channels
two := make(chan int)

for{
	select {
	case o := <- one
		fmt.Println("one",o)
	case t := <- two
		fmt.Println("two",t)
	case <-time.After(300*time.Millisecond):   with this configuration if data doesn't comes from channels after 300 ms then it will time out and return this kind of configuration are important when data goes stale after certain amount of time
		fmt.Println("timed out")
		return
	}
}
*/

/*
channels are bidirectional communication pipes
 they have a send/write end and a receive / read end
the ends of a channel can be duplicated across goroutines
select can be used to send or receive on multiple different channels
buffered unblocking , unbuffered blocking
*/

package main

import (
	"fmt"
	"time"
)

type ControllerMsg int

const (
	DoExit = iota
	ExitOk
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

// we are only going to be reading the data from jobs channel which is why we kept <- before chan (receive only)
// and we will be writing to result chan which is why we kept <- after chan keyword (send only)
// control is going to bidirectional we will be reading and responding to the same channel
func Doubler(jobs <-chan Job, results chan<- Result, control chan ControllerMsg) {
	fmt.Println("Inside doubler")
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("exit goroutine")
				control <- ExitOk
				return
			default:
				panic("unhandled ")
			}
		case job := <-jobs:
			results <- Result{result: job.data * 2, job: job}
		}
	}
}

func greet(c <-chan string) {
	greetingMsg := <-c
	fmt.Println(greetingMsg)
}

func main() {
	jobs := make(chan Job, 50)
	result := make(chan Result, 50)
	control := make(chan ControllerMsg)

	fmt.Println("calling goroutine doubler")
	go Doubler(jobs, result, control)

	c := make(chan string)
	go greet(c)

	// c <- "Hello channels"

	fmt.Println("assigning jobs")
	for i := 0; i <= 30; i++ {
		jobs <- Job{i}
	}

	c <- "Hello channels"
	for {
		fmt.Println("here")
		select {
		case result := <-result:
			fmt.Println(result)
		case <-time.After(10000 * time.Millisecond):
			fmt.Println("timed out")
			control <- DoExit // we are gonna send a message to control channel indicating to the goroutine that it is a quit using doexit message
			<-control         // wait for the goroutine to respond
			fmt.Println("exit program")
			return
		}
	}
}
