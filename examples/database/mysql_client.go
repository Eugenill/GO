package main
//package database

import (
	"database/sql"
	"fmt"
)



type MySqlClient struct {

}

func NewSqlClient(db_params string) *sql.DB {
	db, err := sql.Open("mysql", db_params) //crear el db

	if err != nil {
		_ = fmt.Errorf("cannot create db tentat: %s", err.Error()) //print error if error is arraised
		panic("Db creation terminated")//to stop the execution
	}

	return db
}

func main() {
	NewSqlClient("user:password@tcp(127.0.0.1:3306)/hello")
}