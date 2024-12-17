package advanced

import (
	"database/sql"
	"fmt"
	"log"
)

func PreparedStatement(db *sql.DB) {
	stat, err := db.Prepare("insert INTO students(id,name,age) values($1,$2,$3)")
	if err != nil {
		log.Fatal("prepare error is", err)
	}
	exec, err := stat.Exec(9, "Hesen", 32)
	if err != nil {
		log.Fatal("exec error is", err)
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected error")
	}
	if affected == 1 {
		fmt.Println("Insert successful")
	}
}
