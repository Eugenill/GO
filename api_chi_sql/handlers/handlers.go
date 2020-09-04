package handlers

import (
	"database/sql"
	_ "database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"go_examples/api_chi_sql/helper"
	"net/http"
	"strconv"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var DB *sql.DB //DB is a DB variable from sql

func AllPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("select * from posts") //*sql.rows
	helper.Catch(err)
	defer rows.Close()
	columns, err := rows.Columns() //[id title content]
	helper.Catch(err)

	fmt.Println(columns, rows)

	count := len(columns)
	tableData := make([]map[string]interface{}, 0) //list of maps[string]interface{}, 0 maps initially
	values := make([]interface{}, count)           //list of interfaces = [<nil> <nil> <nil>]
	valuePtrs := make([]interface{}, count)        //pointer of values

	for rows.Next() { //for every row.....
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i] //we assign every pointer to every value
		}
		err = rows.Scan(valuePtrs...) //Scan(dest ...interface{}) valuePtrs is variadic: https://golang.org/ref/spec#Passing_arguments_to_..._parameters
		//Scan copies the columns in the current row into the values pointed at by dest.
		//The number of values in dest must be the same as the number of columns in Rows.
		helper.Catch(err)
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
	helper.RespondwithJSON(w, http.StatusOK, tableData)
}

func DetailPost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	rows, err := DB.Query("select * from posts") //*sql.rows
	helper.Catch(err)
	row := DB.QueryRow("select * from posts where id=?", id)
	defer rows.Close()
	columns, err := rows.Columns() //[id title content]
	helper.Catch(err)

	count := len(columns)
	tableData := make([]map[string]interface{}, 0) //list of maps[string]interface{}, 0 maps initially
	values := make([]interface{}, count)           //list of interfaces = [<nil> <nil> <nil>]
	valuePtrs := make([]interface{}, count)        //pointer of values

	for i := 0; i < count; i++ {
		valuePtrs[i] = &values[i] //we assign every pointer to every value
	}
	err = row.Scan(valuePtrs...) //Scan(dest ...interface{}) valuePtrs is variadic: https://golang.org/ref/spec#Passing_arguments_to_..._parameters
	//Scan copies the columns in the current row into the values pointed at by dest.
	//The number of values in dest must be the same as the number of columns in Rows.
	helper.Catch(err)
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

	helper.RespondwithJSON(w, http.StatusOK, tableData)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post) //we helper.Catch the data in the Body of the POST ( r )
	//and decode it to our struct model and save it in the post var
	//so it is a pointer because we are modifying post itself

	query, err := DB.Prepare("Insert posts SET id=?, title=?, content=?") //title=? means we have dynamic title data to
	// execute which we will retrieve from post variable
	helper.Catch(err)

	_, er := query.Exec(post.ID, post.Title, post.Content) //here we define the title and content as it is in post
	//Exec executes a query without returning any rows
	helper.Catch(er)

	defer query.Close() //DONT FORGET TO CLOSE THE QUERY

	helper.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func UpdatePost(w http.ResponseWriter, req *http.Request) {
	var post Post
	id := chi.URLParam(req, "id") //we take the id of the post
	json.NewDecoder(req.Body).Decode(&post)

	query, err := DB.Prepare("Update posts set title=?, content=? where id=?")
	helper.Catch(err)
	_, er := query.Exec(post.Title, post.Content, id)
	helper.Catch(er)

	defer query.Close()

	helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := DB.Prepare("delete from posts where id=?")
	helper.Catch(err)

	_, er := query.Exec(id)
	helper.Catch(er)
	defer query.Close()

	helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "deleted successfully"})

}
