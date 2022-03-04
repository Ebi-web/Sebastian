package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConf := "root:password0301@tcp(mysql)/sebastian"
	db, err := sql.Open("mysql", dbConf)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("connect")
	}
	defer db.Close()
}
