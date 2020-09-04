package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//usage: go run handling_server.go log.go
func main() {
	http.HandleFunc("/", Decorate(handler))
	panic(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx2 := &gin.Context{
		Request: r,
	}
	ctx := r.Context()
	//ctx = context&channels.WithValue(ctx, key(50), int64(100)) //-> here we specify that we want every requestId in the context&channels to be 100
	ctx = context.WithValue(ctx, 50, int64(100)) //here we are defining the value 100 in the key 50, but the Decorator only modifies in the key key(50)
	//log inici i final en el server
	Print(ctx2.Request.Host, "handler started")
	defer Println(ctx, "handler ended")
	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "Hello")
	case <-ctx.Done():
		err := ctx.Err()
		Println(ctx, err.Error())                                  //will print de log "context&channels canceled"
		http.Error(w, err.Error(), http.StatusInternalServerError) //if canceled it will return an interal server error
	}
	time.Sleep(5 * time.Second)
}
