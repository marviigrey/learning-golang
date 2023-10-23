package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

type Data struct {
	id   int
	name string
}

func main() {
	connectionString := fmt.Sprintf("%v:%v@tcp(54.221.99.8:3306)/%v", DbUser, DbPassword, DBname)
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	rows, err := db.Query("SELECT * from data")
	checkError(err)

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		checkError(err)
		fmt.Println(data)
	}
}
