package main

import (
  "log"
  "net/http"
  "time"
)
/*
The DefaultServeMux is just a plain ol' ServeMux like we've already been using, 
which gets instantiated by default when the HTTP package is used. 
Here's the relevant line from the Go source:

  var DefaultServeMux = NewServeMux()

Generally you shouldn't use the DefaultServeMux because it poses a security risk.
Because the DefaultServeMux is stored in a global variable, 
any package is able to access it and register a route 
*/
func timeHandler(format string) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(format)
    w.Write([]byte("The time is: " + tm))
  }
  return http.HandlerFunc(fn)
}

func main() {
  // Note that we skip creating the ServeMux...

  var format string = time.RFC1123
  th := timeHandler(format)

  // We use http.Handle instead of mux.Handle, we use the DefaultServeMux
  http.Handle("/time", th)
  //SAME WITH http.HandleFunc()

  log.Println("Listening...")
  // And pass nil as the handler to ListenAndServe (DefaultServeMux)
  http.ListenAndServe(":3000", nil) 
}

