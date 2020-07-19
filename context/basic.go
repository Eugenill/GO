package main

import (
	"context"
	"time"
	"log"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 7* time.Second) //will be canceled after x second

	defer cancel()
	
	mySleepAndTalk(ctx, 5*time.Second, "hello")		
	//depending on when ctx is canceled we wil print "hello" or not
}

func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}