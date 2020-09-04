package main

// Orginial code: https://itnext.io/building-restful-web-api-service-using-golang-chi-mysql-d85f427dee54
import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware" //f any request fails, your app won’t die, you can request again without restarting your app.
	_ "github.com/lib/pq"
	"go_examples/api_chi_sql/handlers"
	"go_examples/api_chi_sql/helper"
	"go_examples/api_chi_sql/router"
	"log"
	"net/http"
)

var route *chi.Mux // server is a Mux object from Chi

//constants for the db
const (
	dbName = "go-api-sql"
	dbUser = "postgres"
	dbPass = "postgres"
	dbHost = "localhost"
	dbPort = "5432"
)

/*
   1. Create SQL Database

   CREATE DATABASE IF NOT EXISTS go-api-sql ;
   \connect go-api-sql;

   2. Create table for `posts`

   DROP TABLE IF EXISTS posts;
   CREATE TABLE posts (
     id INT NOT NULL PRIMARY KEY,
     title VARCHAR(100) NOT NULL,
     content TEXT NOT NULL
   );
*/

func init() {
	//func NewRouter() *Mux, its like if they did router=&newMux (if NewMux were a struct)
	route = chi.NewRouter() //NewRouter returns a new Mux object (*chi.Mux) that implements the Router interface,
	//so we assign it to router (which a pointer to a chi.Mux)
	route.Use(middleware.Recoverer) //Recoverer is a middleware that recovers from panics, logs the panic (and a
	// backtrace), and returns a HTTP 500 (Internal Server Error) status if
	// possible. Recoverer prints a request ID if one is provided.
	dbSource := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPass)
	//%s for strings, %d for int
	var err error                                     //error var
	handlers.DB, err = sql.Open("postgres", dbSource) //Create the sql db with the name: "mysql" and the source
	helper.Catch(err)
	log.Println("Connected to the local DB!")
}

func main() {
	router.SetEndpoints(route)
	err := http.ListenAndServe(":8000", helper.Logger(route))

	helper.Catch(err)
}
