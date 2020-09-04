package helper

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

func Catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	RespondwithJSON(w, code, map[string]string{"message": msg})
}

// respondwithJSON write json response format
func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload) //returns the json encoding of payload ([]byte, error)
	//ORDERED!!!
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Logger return log message
func Logger(router *chi.Mux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL) //example: 2020-07-06 09:32:44.634333 +0200 CEST m=+22.456178180 GET /posts
		router.ServeHTTP(w, r)                   // dispatch the request, we are dispatching the request to the router, which is who has the handlers to serve the request
	})
}
