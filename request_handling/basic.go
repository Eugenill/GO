package main

import (
  "log"
  "net/http"
)

func main() {
  mux := http.NewServeMux() //New router

  rh := http.RedirectHandler("http://example.org", 307) 
  //returns a new handler, which redirects requests to http://examples.org
  mux.Handle("/foo", rh)
  //here we register the handler on our router, so it acts as a handler for incoming requests with the URLpath /foo

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux) //Initialize router
  /*
  	The eagle-eyed of you might have noticed something interesting: 
  	The signature for the ListenAndServe function is 
  	ListenAndServe(addr string, handler Handler), 
  	but we passed a ServeMux as the second parameter.

	We were able to do this because the ServeMux type also has a ServeHTTP method, 
	meaning that it too satisfies the Handler interface.

	Which instead of providing a response itself passes the request on to a second handler.
  */
}

