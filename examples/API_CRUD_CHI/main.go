package main

// Orginial code: https://itnext.io/building-restful-web-api-service-using-golang-chi-mysql-d85f427dee54
import (
    "fmt"
    "database/sql"
    "net/http"
    "encoding/json"
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
    //func NewRouter() *Mux, its like if they did router=&newMux (if NewMux were a struct)
    router = chi.NewRouter() //NewRouter returns a new Mux object (*chi.Mux) that implements the Router interface, 
    fmt.Println("%T",router)                       //so we assign it to router (which a pointer to a chi.Mux)
    router.Use(middleware.Recoverer)  //Recoverer is a middleware that recovers from panics, logs the panic (and a
                                      // backtrace), and returns a HTTP 500 (Internal Server Error) status if
                                      // possible. Recoverer prints a request ID if one is provided.
    
    dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8",  dbPass, dbHost, dbPort, dbName) //%s for strings, %d for int
                            //username:password@protocol(address)/dbname?param=value
    var err error //error var
    db, err = sql.Open("mysql", dbSource) //Create the sql db with the name: "mysql" and the source
    
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


    query, err := db.Prepare("Insert posts SET title=?, content=?") //title=? means we have dynamic title data to
                                                                    // execute which we will retrieve from post variable
    catch(err)

    _,er := query.Exec(post.Title, post.Content)  //here we define the title and content as it is in post                                     
    catch(er)

    defer query.Close() //DONT FORGET TO CLOSE THE QUERY

    respondwithJSON(w, http.StatusCreated, map[string]string{"message": "succesfully created"})
}

func UpdatePost(w http.ResponseWriter, req *http.Request) {
    var post Post
    id := chi.URLParam(req, "id") //we take the id of the post 
    json.NewDecoder(req.Body).Decode(&post)

    query, err := db.Prepare("Update posts set title=?, content=? where id=?")
    catch(err)
    _, er := query.Exec(post.Title, post.Content, id)
    catch(er)

    defer query.Close()

    respondwithJSON(w, http.StatusOK, map[string]string{"message":"update succesfully"})
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    query, err := db.Prepare("delete from posts where id=?")
    catch(err)

    _,er := query.Exec(id)
    catch(er)
    query.Close()

    respondwithJSON(w, http.StatusOK, map[string]string{"message":"deleted succesfully"})

}


func main() {
    routers()
    http.ListenAndServe(":8000", Logger())
}