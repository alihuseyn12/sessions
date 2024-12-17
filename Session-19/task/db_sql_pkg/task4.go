package db_sql_pkg

import (
	"database/sql"
	"fmt"
	"log"
)

func QueryStudents(db *sql.DB) {
	rows, err := db.Query("select * from students")
	if err != nil {
		log.Fatal("error is query", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("error is query", err)
		}
		fmt.Printf("id:%d , name:%s ,age:%d\n", id, name, age)

	}
}
