package main

// Orginial code: https://itnext.io/building-restful-web-api-service-using-golang-chi-mysql-d85f427dee54
import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware" //f any request fails, your app won’t die, you can request again without restarting your app.
	_ "github.com/go-sql-driver/mysql" //driver (_ ) for database/sql, if we import the driver we can use the whole API
	"go_examples/api_chi_sql/handlers"
	"go_examples/api_chi_sql/helper"
	"go_examples/api_chi_sql/router"
	"net/http"
)

var route *chi.Mux // server is a Mux object from Chi

//constants for the db
const (
	dbName = "go-mysql-crud"
	dbPass = "*******"
	dbHost = "localhost"
	dbPort = "3306" //TCP - MySQL clients to the MySQL server (MySQL Protocol)
)

/*

   CREATE DATABASE IF NOT EXISTS `go-mysql-crud` !40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci ;
   USE `go-mysql-crud`;
    Create table for `posts`

   DROP TABLE IF EXISTS `posts`;
   CREATE TABLE `posts` (
     `id` int(11) NOT NULL AUTO_INCREMENT,
     `title` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
     `content` longtext COLLATE utf8_unicode_ci NOT NULL,
     PRIMARY KEY (`id`)
   );
*/

func init() {
	//func NewRouter() *Mux, its like if they did router=&newMux (if NewMux were a struct)
	route = chi.NewRouter() //NewRouter returns a new Mux object (*chi.Mux) that implements the Router interface,
	//so we assign it to router (which a pointer to a chi.Mux)
	route.Use(middleware.Recoverer) //Recoverer is a middleware that recovers from panics, logs the panic (and a
	// backtrace), and returns a HTTP 500 (Internal Server Error) status if
	// possible. Recoverer prints a request ID if one is provided.
	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName) //%s for strings, %d for int
	//username:password@protocol(address)/dbname?param=value
	var err error                                  //error var
	handlers.DB, err = sql.Open("mysql", dbSource) //Create the sql db with the name: "mysql" and the source
	helper.Catch(err)
}

func main() {
	router.SetEndpoints(route)
	http.ListenAndServe(":8000", helper.Logger(route))
}
