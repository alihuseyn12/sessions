package advanced

import (
	"database/sql"
	"log"
)

func Transaction(db *sql.DB) {
	begin, err := db.Begin()
	if err != nil {

		log.Fatal("begin db err", err)
	}
	_, err = begin.Exec("update students SET age =39 where name='Ali'")
	if err != nil {
		begin.Rollback()
		log.Fatal("exec err", err)
	}
	err = begin.Commit()
	if err != nil {
		log.Fatal("commit err", err)
	}

}
