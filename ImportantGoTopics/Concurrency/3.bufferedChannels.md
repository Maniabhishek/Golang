## Buffered Channels

> * This type of channel allows you to store more than one piece of data specified by the capacity, when that capacity is reached, subsequent messages that are sent to the channel will block until at least one message is read so that the channel has capacity again.
> * To create a buffered channel we only need to pass an additional parameter to the make function: 
```go bufchan := make(chan int, 5)```
> * this channel will accept 5 integers , if 6th integer is sent then receive operation has to be made until then send will be blocked similary when receive operation is made when channel is completely emtpy then receive operation will be blocked until sent operation is performed
> * The data structure used to keep track of the capacity of the channel is a queue, which means that the first element that gets into the queue will be the first getting out of the queue.

### lets look at an example for a buffered channels
```go
package dummy

import (
	"fmt"
	"math/rand"
	"time"
)

func BufferedFunc() {
	bufferedChan := make(chan int, 5)
	go generateRandomNumber(bufferedChan)

	time.Sleep(time.Second * 5)
	for buf := range bufferedChan {
		fmt.Println(buf)
	}
}

func generateRandomNumber(bufferedChan chan<- int) {
	defer close(bufferedChan)
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		randomNum := rand.Intn(10)
		bufferedChan <- randomNum
	}
}
```

### let's look at an another example
```go 
package dummy

import (
	"fmt"
	"time"
)

func BufferedChan() {
	superHeroesChan := make(chan string, 3)
	go getSuperHeroes(superHeroesChan)
	time.Sleep(time.Second * 5)
	for hero := range superHeroesChan {
		fmt.Println("mainfunc ", hero)
	}
}

func getSuperHeroes(superHeroesChan chan<- string) {
	defer close(superHeroesChan)
	heroes := []string{"batman", "superman", "ironman", "hulk", "captain america"}
	for _, name := range heroes {
		superHeroesChan <- name
		fmt.Println(name)
	}
}

```

> * in the above program we created a channel called superHeroesChan which is buffered channel with capacity 3 
> * then we are calling a go routine and passing channel to it and after that we have 5 seconds simulation to stop the current go routine 
> * the second go routine send 3 name continously and then gets blcoked after after 3 seconds as it has no more space 
> * once main go routine is resumes after 5 seconds it will start reading from the channel the mnoment it start reading chanel will again have space then it will send another one again and so on  
