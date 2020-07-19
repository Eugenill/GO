package main

// Orginial code: https://itnext.io/building-restful-web-api-service-using-golang-chi-mysql-d85f427dee54
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware" //f any request fails, your app won’t die, you can request again without restarting your app.
	_ "github.com/go-sql-driver/mysql" //driver (_ ) for database/sql, if we import the driver we can use the whole API
	"net/http"
	"strconv"
)

var router *chi.Mux // router is a Mux object from Chi
var db *sql.DB      //db is a DB variable from sql

//constants for the db
const (
	dbName = "go-mysql-crud"
	dbPass = "*******"
	dbHost = "localhost"
	dbPort = "3306" //TCP - MySQL clients to the MySQL server (MySQL Protocol)
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
	//so we assign it to router (which a pointer to a chi.Mux)
	router.Use(middleware.Recoverer) //Recoverer is a middleware that recovers from panics, logs the panic (and a
	// backtrace), and returns a HTTP 500 (Internal Server Error) status if
	// possible. Recoverer prints a request ID if one is provided.

	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName) //%s for strings, %d for int
	//username:password@protocol(address)/dbname?param=value
	var err error                         //error var
	db, err = sql.Open("mysql", dbSource) //Create the sql db with the name: "mysql" and the source

	catch(err)
}

type Post struct {
	ID      int    `json: "id"`
	Title   string `json: "title"`
	Content string `json: "content"`
}

func AllPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from posts") //*sql.rows
	catch(err)
	defer rows.Close()
	columns, err := rows.Columns() //[id title content]
	catch(err)

	fmt.Println(columns, rows)

	count := len(columns)
	tableData := make([]map[string]interface{}, 0) //list of maps[string]interface{}, 0 maps initially
	values := make([]interface{}, count)           //list of interfaces = [<nil> <nil> <nil>]
	valuePtrs := make([]interface{}, count)        //pointer of values

	for rows.Next() { //for every row.....
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i] //we assign every pointer to every value
		}
		rows.Scan(valuePtrs...) //Scan(dest ...interface{}) valuePtrs is variadic: https://golang.org/ref/spec#Passing_arguments_to_..._parameters
		//Scan copies the columns in the current row into the values pointed at by dest.
		//The number of values in dest must be the same as the number of columns in Rows.
		entry := make(map[string]interface{}) //map[]
		fmt.Println(values)
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			fmt.Println(b)
			if col == "id" {
				if ok {
					v, _ = strconv.Atoi(string(b))
				} else {
					v = val
				}
			} else {
				if ok {
					v = string(b)
				} else {
					v = val
				}
			}

			entry[col] = v //we define {"column": value, ...}
		}
		tableData = append(tableData, entry) //append map to the list of maps, for every row
	}
	respondwithJSON(w, http.StatusOK, tableData)
}

func DetailPost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	rows, err := db.Query("select * from posts") //*sql.rows
	catch(err)
	row := db.QueryRow("select * from posts where id=?", id)
	defer rows.Close()
	columns, err := rows.Columns() //[id title content]
	catch(err)

	count := len(columns)
	tableData := make([]map[string]interface{}, 0) //list of maps[string]interface{}, 0 maps initially
	values := make([]interface{}, count)           //list of interfaces = [<nil> <nil> <nil>]
	valuePtrs := make([]interface{}, count)        //pointer of values

	for i := 0; i < count; i++ {
		valuePtrs[i] = &values[i] //we assign every pointer to every value
	}
	row.Scan(valuePtrs...) //Scan(dest ...interface{}) valuePtrs is variadic: https://golang.org/ref/spec#Passing_arguments_to_..._parameters
	//Scan copies the columns in the current row into the values pointed at by dest.
	//The number of values in dest must be the same as the number of columns in Rows.
	entry := make(map[string]interface{}) //map[]
	fmt.Println(values)
	for i, col := range columns {
		var v interface{}
		val := values[i]
		b, ok := val.([]byte)
		fmt.Println(b)
		if col == "id" {
			if ok {
				v, _ = strconv.Atoi(string(b))
			} else {
				v = val
			}
		} else {
			if ok {
				v = string(b)
			} else {
				v = val
			}
		}

		entry[col] = v //we define {"column": value, ...}
	}
	tableData = append(tableData, entry) //append map to the list of maps, for every row

	respondwithJSON(w, http.StatusOK, tableData)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post) //we catch the data in the Body of the POST ( r )
	//and decode it to our struct model and save it in the post var
	//so it is a pointer because we are modifying post itself

	query, err := db.Prepare("Insert posts SET id=?, title=?, content=?") //title=? means we have dynamic title data to
	// execute which we will retrieve from post variable
	catch(err)

	_, er := query.Exec(post.ID, post.Title, post.Content) //here we define the title and content as it is in post
	//Exec executes a query without returning any rows
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

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "update succesfully"})
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := db.Prepare("delete from posts where id=?")
	catch(err)

	_, er := query.Exec(id)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "deleted succesfully"})

}

func main() {
	routers()
	http.ListenAndServe(":8000", Logger())
}
