/*
code only executes line by line, one line at a time
concurrency allows multiple lines to be executed
two types of concurrent code:
 threaded : code runs in parallel based on number of CPU cores
 asynchronous : code can pause and resume execution
  while paused other code can resume
go will automatically choose the appropriate concurrency method
*/

/*
threaded execution
there will be one main thread and it will branch of other executions based on number of cores
and it will wait on the main thread until all the process is executed
*/
/*
main thread which suppose execute 4 jobs a b c d  in this case it utilizes cpu
it will execute a if it takes time then it go on and execute b if this takes time then it will
jump to c and so on, until all the task is completed in this way cpu is utilized
*/

/*
Details
single-threaded code runs deterministically
 each run will produce the same result
concurrent code runs non deterministically
 code no longer executes line by line in predefined order
  cannot rely on results being the same each program run
extra care should be taken to ensure results are in order / sorted properly
 accomplished using synchronization or by checking the final results in a single thread

*/

//go routines
/*
Goroutines allow functions to run concurrently
 can also run function literals / closures
Go will automatically select parallel or asynchronous execution
New goroutines can be created with the go keyword
*/

//example basic
/*
func count(amount int){
	for i:=1 ; i<amount ; i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(i)
	}
}

func main(){
	go count(5)
	fmt.Println("wait for go routine")
	time.Sleep(1000*time.Millisecond)
	fmt.Println(end Program)
}
*/

/*
example closures
counter := 0
wait := func(ms time.Duration){
	time.Sleep(ms * time.Millisecond)
	count +=1
}
fmt.Println("launching goroutines")
go wait(100)
go wait(900)
go wait(1000)
time.Sleep(1100*time.Milleseconds)
fmt.Println(counter)
*/

package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune
	capIt := func(char rune) {
		capitalized = append(capitalized, unicode.ToUpper(char))
		fmt.Printf("%c done \n", char)
	}

	for _, r := range data {
		go capIt(r)
	}
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("%c", capitalized)
}
