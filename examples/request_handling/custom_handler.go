package main

import (
  "log"
  "net/http"
  "time"
  "fmt"
)

type timeHandler struct { //it could be anything else
  format string
}
//All that matter is that we add the method ServeHTTP to make a Handler of it
func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //now timeHandler is also a Handler type
  tm := time.Now().Format(th.format) //now time in the format specified 
  w.Write([]byte("The time is: " + tm)) 
  //func (r *Request) Write(w io.Writer) error
  /*
  type Writer interface {
      Write(p []byte) (n int, err error)
  }
  */

}

func main() {
  mux := http.NewServeMux()

  th := &timeHandler{format: time.RFC1123} //it's not necessary to use a pointer but we do it in case we modify the handler later
  //new handler with format : time(from package time).RFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"
  fmt.Printf("%T\n",th)//*main.timeHandler
  mux.Handle("/time", th) 
  //here we register the handler on our router, so it acts as a handler for incoming requests with the URLpath /time

  //other routes
  th1123 := &timeHandler{format: time.RFC1123}
  mux.Handle("/time/rfc1123", th1123)

  th3339 := &timeHandler{format: time.RFC3339}
  mux.Handle("/time/rfc3339", th3339)

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux)
}