package WorkingWithDatabaseGorm

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID            uint   `gorm:"primary_key;auto_increment"`
	Title         string `gorm:"size:255;not null"`
	Author        string `gorm:"size:100"`
	PublishedYear int    `gorm:"not null"`
	CreatedAt     time.Time
}

func AutoMigrater(db *gorm.DB) {
	err := db.AutoMigrate(&Book{})
	if err != nil {
		fmt.Errorf("Automigrate Error ", err)
	}
	fmt.Println("Database migration completed successfully!")
}
