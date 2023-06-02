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
```
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
