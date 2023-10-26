package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func checkError(e error) { //and error logic: this will log an error message if the connection to the mysql server didnt go through
	if e != nil {
		log.Fatalln(e)
	}
}

type Data struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open(mysql, connectionString)
	connectionString := fmt.Sprintf("%v:%v@tcp(localhost:3306)/%v", DbUser, DbPassword, DbName)
	db, err := sql.Open(mysql, connectionString)
	defer db.Close()
	result, err := db.Exec("insert into data values(4, 'xyz')")
	checkError(err)
	lastInsertedId, err := result.LastInsertId()
	fmt.Println("last insertedId: ", lastInsertedId)
	checkError(err)
	rowsAffected, err := result.RowsAffected()
	fmt.Println(rowsAffected)
	checkError(err)
	rows, err := sql.Query("SELECT * from data")
	checkError(err)

	//creates a variable called data off the struct called Data and saves it,scans through it and returns and error
	//
	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		checkError(err)
		fmt.Println(data)
	}

}
