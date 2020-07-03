package main

import (
  "log"
  "net/http"
  "time"
)
//The difference is that we are using the timeHandler function to return a handler
func timeHandler(format string) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(format) //we can acces the params in the first function
    w.Write([]byte("The time is: " + tm))
  }
  return http.HandlerFunc(fn) //fn must have the nature func(http.ResponseWriter, *http.Request)
}

/*
SAME AS ABOVE, but with the return in top:

  func timeHandler(format string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      tm := time.Now().Format(format)
      w.Write([]byte("The time is: " + tm))
    })
  }

SAME AS ABOVE, but with implicit conversion to the HandlerFunc type on return:

func timeHandler(format string) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(format)
    w.Write([]byte("The time is: " + tm))
  }
}
*/

func main() {
  mux := http.NewServeMux()

  th := timeHandler(time.RFC1123) //we call the function and return a Handler(Func)
  mux.Handle("/time", th)//We add this handler to the mux

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux)
}