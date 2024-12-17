package WorkingWithDatabaseGorm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConDb() *gorm.DB {
	conStr := "host=localhost user=postgres port=5432 password=postgres dbname=testdb sslmode=disable TimeZone=Asia/Baku"

	db, err := gorm.Open(postgres.Open(conStr), &gorm.Config{})
	if err != nil {
		fmt.Errorf("connection error", err)
	}

	fmt.Println("Connected to PostgreSQL database successfully!")
	return db
}
