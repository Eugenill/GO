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

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

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
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

//------------------------------------------------------------------------------------------------------
