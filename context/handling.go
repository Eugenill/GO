package main

import (
	"fmt"
)
func main() {
	for days := 360; days>=30; {
		fmt.Println(days)
		days--
	}
}