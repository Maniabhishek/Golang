# Concurrency in Go, Goroutines, and channels 

> Go is a powerful language 
>> Because it can handle several things effectively and efficiently, and something which is more facinating and powerful is how it handles concurrency

## What is concurrency 
> - A program is called as concurrent if it can handle multiple tasks at the same time. 
> - It can run several operations concurrently, which doesn't mean all the operations are running explicitly at the same time, each task can start at a different point in time 
> - In concurrency , all the operation may not start at the same time , but they are nunning concurrently to each other. which means that a task need not to wait until another task finish before running.
> - there is something called parallelism , if your program runs multiple operations at the same time, meaning they start exactly at the same point in time.
> - using concurrency your program can achieve parallelism, but it depends on your use cases

## Concurrency in Go
> - Golang uses goroutines for concurrency.
>> A goroutine is similar to thread but with several advantages and goroutine is managed by go runtime, which allows us to run multiple operations concurrently.
> - what happens in a multithreaded environment?
>> if you want to run various operations concurrently a new thread has to be created by the OS, which involves a considerable amount of resources, memory and time thus using thread to these stuffs are more expensive by the OS 
> - What about goroutines 
>> goroutines are lightweight, efficient and it does not cost too many resources to be created, you can spin up many goroutines which will not be a problem in GO

## Race Condition
> - when you are running multiple goroutines, you will find that goroutines need to access and modify the shared resources , for example if multiple goroutines try to access and modify the same data at the same time this will lead to several problems, unexpected results... this is called as race conditions(in this problem multiple operations running concurrently attempt to read/write the same data at the same time)

## How to avoid this situation
> TO avoid this , go uses locks so that only one goroutine can modify a certain piece of data at a time 

## Let's see how can we create a go routine

```go
package main
import "fmt"

func main(){
    fmt.Println("start")
    go printNum(10)

    fmt.Println("end")
}

func printNum(num int){
    num2 := 5
    for(i:= 1; i< num ; i++){
        num2 = num2 * i
    }
    fmt.Printf(num2)
}
```
> we have a main function above which is printing start at line 5 and end at 8, when program runs it create the main goroutine, this goroutine is created automatically , and this is where all your code is executed, we can run another go routine to execute another operation in order to run it concurrently , at line 6 we are telling to create a new goroutine for the printNum function. then in main goroutine there is one print statement 

> - So, why don’t we see the print statement that is inside the printNum function?. Well, it is because the main goroutine does not wait for other goroutines to finish their work, the main goroutine will continue the execution of the main program and it will terminate without waiting to see if other goroutines have finished.
> - To see the print in printNum we can put time.Sleep(1*time.Second) at line 7
> - this solution is not ideal in any concurrent program, later we will see how we can solve using channels 

## What is a channel 
> - A channel is a way to cimmunicate between different goroutines. 
> - It is a safe way of sending value of one data type to another goroutine
> - A channel is passed by reference meaning that when creating a channel and then passing it to other functions, these functions will have the same reference pointing to the same channel. 
>> We can compare channels only if they have the same type and as I previously mentioned since they are passed by reference, a comparison between two channels will evaluate true if both are pointing to the same reference in memory. We can also compare a channel with nil.
> - The purpose of a channel is to allow goroutines to send and receive information, but frequently they are also used to inform other goroutines that a process has finished and not necessarily sending any information through the channel.
> - A channel can also be closed, meaning it will no longer accept any more messages to be sent or received and if a goroutine tries to send or receive a message from a closed channel the program will panic, unless we use a special syntax to read from the channel or we use the a range loop. We’ll see in a moment how this works.

```go
package main 
import (
    "fmt"
    "time"
)

func sender(c chan string){
    for i:=0; ; i++{
        c <- "ping"
    }
}

func printer(c chan string){
    for{
        msg := <- c
        fmt.println(msg)
        time.sleep(time.Second * 1)
    }
}

func main(){
    var c chan string = make(chan string)
    go sender(c)
    go printer(c)

    var input string 
    fmt.Scanln(&input)
}
```

> - this above program will print ping forever (hit enter to stop it). 
> - A channel type is represented with the keyword chan followed by the type of the things that are passed on the channel (in this case we are passing strings)
> - the left arrow operator (<-) is used to send and receive messages on the channel.
> - c <- "ping" means send "ping"
> - msg := <- c means receive a message and store it in msg. The fmt line could also have been written like fmt.Println(<-c), in which case we could remove the previous line.
> - Using a channel like this synchronizes the two goroutines. When sender attempts to send a message on the channel, it will wait until printer is ready to receive the message (this is known as blocking). 
> - Let’s add another sender to the program and see what happens. Add this function:

```go
func sender2(c chan string) {
    for i := 0; ; i++ {
        c <- "pong"
    }
}
// and modify main
func main() {
    var c chan string = make(chan string)

    go sender(c)
    go sender2(c)
    go printer(c)

    var input string
    fmt.Scanln(&input)
}

// the program will now start printing ping and pong
```

## channel direction
> - we can specify a direction on a channel type, thus restricting it to either sending or receiving. For example, sender's function signature can be changed to this:
```go
func sender(c chan<- string)
```
* Now the above sender is only allowed to send to c. Attempting to receive from c will result in a compile time error. Similarly, we can change printer to this:
```go
func pointer(c <-chan string)
```
* A channel that doesn't have these restriction is know as ***bidirectional***
* A bidirectional channel can be passed to a function that takes send-only or receive-only channels, but the reverse is not true.

## Select
> - Go has a special statement called select that works like a switch but for channels:

```go
func main(){
    c1 := make(chan string)
    c2 := make(chan string)

    go func(){
        for{
            c1 <- "func 1"
            time.Sleep(time.Second * 1)
        }
    }()

    go func(){
        for{
            c2 <- "func 2"
            time.Sleep(time.Second * 1)
        }
    }()
    go func(){
        for {
            select {
                case msg1 := <-c1:
                    fmt.Println(msg1)
                case msg2 := <-c2:
                    fmt.Println(msg2)
            }
        }
    }()

    var input string
    fmt.Println(&input)
}
```
> - This program prints “from 1” every 2 seconds and “from 2” every 3 seconds. select picks the first channel that is ready and receives from it (or sends to it).
> - If more than one of the channels are ready, then it randomly picks which one to receive from. If none of the channels are ready, the statement blocks until one becomes available.
> - The ***select*** statement is often used to implement a timeout
```go
select {
case msg1 := <- c1:
    fmt.Println("Message 1", msg1)
case msg2 := <- c2:
    fmt.Println("Message 2", msg2)
case <- time.After(time.Second):
    fmt.Println("timeout")
}

```
> time.After creates a channel, and after the given duration, will send the current time on it (we weren’t interested in the time, so we didn’t store it in a variable). We can also specify a default case:

```go
select {
case msg1 := <- c1:
    fmt.Println("Message 1", msg1)
case msg2 := <- c2:
    fmt.Println("Message 2", msg2)
case <- time.After(time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("nothing ready")
}
```
The default case happens immediately if none of the channels are ready.

## Type of channels 
### Buffered Channels 
* It’s also possible to pass a second parameter to the make function when creating a channel:
```go
c := make(chan int, 1)
```
> This creates a buffered channel with a capacity of 1. Normally, channels are synchronous both sides of the channel will wait until the other side is ready. A buffered channel is asynchronous; sending or receiving a message will not wait unless the channel is already full. If the channel is full, then sending will wait until there is room for at least one more int.

## Example 
> - Program that uses goroutines and channels. It fetches several web pages simulatneously using the net/http package, and prints the URL of the biggest home page(defined as the most bytes in the response)

```go
package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
)

type PageSize struct {
    URL string 
    Size int
}

func main(){
    urls := []string{
        "http://www.apple.com",
        "http://www.amazon.com",
        "http://www.google.com",
        "http://www.microsoft.com",
    }

    results := make(chan PageSize)
    for _, url := range urls {
        go func(){
            res, err := http.Get(url)
            if err != nil {
                panic(err)
            }
            defer res.Body.close()
            bs, err := ioutil.ReadAll(res.body)
            if err != nil {
                panic(err)
            }

            results <- PageSize{
                URL: url,
                Size: len(bs),
            }
        }(url)
    }

    var biggest PageSize
    for range urls {
        result := <-results
        if biggest.Size < result.Size {
            biggest = result 
        }
    }
    fmt.Println("the biggest home page", biggest.URL)
}
```

> first we define the type that will store PageSize  , then we have list of urls to fetch the data from and then we create the channels and start a new goroutine for each URL (so we will be fetching urls simultaneously). for each url we make an HTTP get request and we store the size of the response body
> Notice that this is an unnamed function that is immediately invoked. This is a common pattern with goroutines, but also notice that we defined the function as taking a single parameter (the url). The reason this function doesn’t reference the url directly—which it is allowed to do—is that the rules of closure are such that if we did that, all four goroutines would probably end up seeing the same value for url. This is because url is changed by the for loop.
