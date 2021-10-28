package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("go sql")

	db, err := sql.Open("mysql", "root:megh1612@tcp(127.0.0.1:3306)/godat")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Print("hehe heehe heehe")
}
