package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password0301@tcp(mysql)/sebastian")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("connect")
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE dictionary (id INT AUTO_INCREMENT, name TEXT, mean TEXT, PRIMARY KEY (id)) DEFAULT CHARSET=utf8;")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("create")
	}
}
