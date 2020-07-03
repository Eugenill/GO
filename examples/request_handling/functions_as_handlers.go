package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
)
//this a function which we want to convert to a handler (func)
func timeHandler(w http.ResponseWriter, r *http.Request) {
  tm := time.Now().Format(time.RFC1123)
  w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux() //new router
	fmt.Printf("%T\n",mux)
	// Convert the timeHandler function to a HandlerFunc type
	th := http.HandlerFunc(timeHandler) 
	fmt.Printf("%T\n",th) //http.HandlerFunc (Handler)

	// And add it to the ServeMux
	mux.Handle("/time", th)

	//But instead of using this two lines above, Go provides a shortcut: mux.HandleFunc method
	
	//mux.HandleFunc("/time", timeHandler) //path + http.HandlerFunc


	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}