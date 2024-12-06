package main

import (
	"Sesssion-19/task/advanced"
	"Sesssion-19/task/db_sql_pkg"
	_ "github.com/lib/pq"
)

func main() {
	db := db_sql_pkg.Connectiondb()

	db_sql_pkg.QueryStudents(db)
	//advanced.PreparedStatement(db)
	advanced.Transaction(db)

	defer db.Close()

}
