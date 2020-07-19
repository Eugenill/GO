package main

import(
	"time"
	"fmt"
)

func main() {
//------------------------------------------------------------------------------------------------------

	//Goroutines
	fmt.Println("\nGoroutines------------------------------------------------------------------")
	go say("world") //go routine, the evaluation happens in the current routine
					//and the execution in the new go routine
	say("hello")
//------------------------------------------------------------------------------------------------------
	//Channels
	fmt.Println("\nChannels------------------------------------------------------------------")
	s := []int{7, 2, 8, -9, 4, 0}

	//Like maps and slices, channels must be created before use:
	c := make(chan int)
	go sum(s[:len(s)/2], c) // [:3] from 0 to 2
	go sum(s[len(s)/2:], c) // [3:] from 3 to 5
	//The data flows in the direction of the arrow:
	x, y := <-c, <-c // receive from 'c', to 'x' and 'y'

	fmt.Println(x, y, x+y) //17 -5 12
//------------------------------------------------------------------------------------------------------
	//Buffered Channels
	fmt.Println("\nBuffered Channels------------------------------------------------------------------")
	//Provide the buffer length as the second argument to make to initialize a buffered channel:
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//Sends to a buffered channel block only when the buffer is full. 
	//Receives block when the buffer is empty.

//------------------------------------------------------------------------------------------------------
	//Range and Close
	fmt.Println("\nRange and Close------------------------------------------------------------------")
	ch2 := make(chan int, 10) //cap(x) = capacity of x
	go fibonacci(cap(ch2), ch2)
	for i := range ch2 {
		fmt.Println(i)
	}
	//A sender can close a channel to indicate that no more values will be sent.
	_, ok := <-ch2 //to test if the channel has been closed we add the 'ok' parameter:
	fmt.Println(ok) //false = closed

//------------------------------------------------------------------------------------------------------
	//Select
	fmt.Println("\nSelect------------------------------------------------------------------")
	ch3 := make(chan int) //to send the fibonacci number
	quit := make(chan int) //to end the execution
	go func() { //subrutien to send and receive data from the channels
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch3) //print the number received from ch3
		}
		quit <- 0
	}()
	fibonacci2(ch3, quit) //run fibonacci2 to send and receive the data from the channels

//------------------------------------------------------------------------------------------------------
	//Default Selection
	fmt.Println("\nDefault Selection------------------------------------------------------------------")
	//The default case in a select is run if no other case is ready
	tick := time.Tick(100 * time.Millisecond) //trigger every 100 ms
	boom := time.After(500 * time.Millisecond) //pulse after 500 ms
	for {
		select { //if we dont receive any 'tick' or 'boom' we print default
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default: 
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}




//FUNCTIONS


//------------------------------------------------------------------------------------------------------
//Goroutines
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//------------------------------------------------------------------------------------------------------

//Channels
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {  //for index,value := range s ...
		sum += v
	}
	//The data flows in the direction of the arrow:
	c <- sum // send 'sum' to channel 'c'
}

//------------------------------------------------------------------------------------------------------

//Range and Close
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//------------------------------------------------------------------------------------------------------

//Select
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for { //infinit loop, untill return
		select {
		case c <- x: //we run this because '<-quit' is not ready
			x, y = y, x+y //update values
		case <-quit: //true when main loop has finished and '0' is received
			fmt.Println("quit")
			return
		}
	}
}

