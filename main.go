package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

func insertCrap(db *sql.DB) {
	for {
		stmt, err := db.Prepare("INSERT INTO answers VALUES(0)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		stmt.Exec()
		stmt.Close()
	}
}

func readCrap(db *sql.DB) {
	for {
		stmt, err := db.Prepare("select count(*) from answers")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		rows, err := stmt.Query()
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			var answerid int
			err := rows.Scan(&answerid)
			if err != nil {
				panic(err)
			}

			fmt.Println(answerid)
		}
	}
}

func main() {
	ping()

	db, err := sql.Open("mysql", "root@/answer_predict")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go insertCrap(db)
	readCrap(db)

	wg.Wait()
}

func ping() {
	db, err := sql.Open("mysql", "root@/answer_predict")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
