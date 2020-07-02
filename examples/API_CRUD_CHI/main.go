package main

// Orginial code: https://itnext.io/building-restful-web-api-service-using-golang-chi-mysql-d85f427dee54
import (
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware" //f any request fails, your app won’t die, you can request again without restarting your app.
    _ "github.com/go-sql-driver/mysql" //driver (_ ) for database/sql, if we import the driver we can use the whole API
)

var router *chi.Mux // router is a Mux object from Chi 
var db *sql.DB //db is a DB variable from sql

//constants for the db
const (
    dbName = "go-mysql-crud"
    dbPass = "******"
    dbHost = "localhost"
    dbPort = "33066"
)

//function with the routers, returning the router with new available functionalities
func routers() *chi.Mux {
    router.Get("/posts", AllPosts)
    router.Get("/posts/{id}", DetailPost)
    router.Post("/posts", CreatePost)
    router.Put("/posts/{id}", UpdatePost)
    router.Delete("/posts/{id}", DeletePost)
    
    return router
}

/*

    CREATE DATABASE  IF NOT EXISTS `go-mysql-crud` !40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci ;
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
    router = chi.NewRouter() //NewRouter returns a new Mux object (*chi.Mux) that implements the Router interface, so we assign it to router (which a pointer to a chi.Mux)
    router.Use(middleware.Recoverer)  //Recoverer is a middleware that recovers from panics, logs the panic (and a
                                      // backtrace), and returns a HTTP 500 (Internal Server Error) status if
                                      // possible. Recoverer prints a request ID if one is provided.
    
    dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8",  dbPass, dbHost, dbPort, dbName) //%s for strings, %d for int
                            //username:password@protocol(address)/dbname?param=value
    var err error //error var
    db, err = sql.Open("mysql", dbSource) //Create the sql db with the name: "mysql" and the source: 
    
    catch(err)
}

type Post struct {
    ID      int    `json: "id"`
    Title   string `json: "title"`
    Content string `json: "content"`
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
    var post Post
    json.NewDecoder(r.Body).Decode(&post) //we catch the data in the Body of the POST ( r )
                                        //and decode it to our struct model and save it in the post var
                                        //so it is a pointer because we are modifying post itself


    query, err := db.Prepare("Insert posts SET title=?, content=?")
    catch(err)

    _,er := query.Exec(post.Title, post.Content)                                       
    catch(err)

    defer query.Close()

    respondwithJSON(w, http.StatusCreated, map[string]string{"message": "succesfully created"})
}