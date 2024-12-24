package WorkingWithDatabaseGorm

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

func CrudNewBook(db *gorm.DB) {
	book := Book{
		Title:         "Go Programming",
		Author:        "John Doe",
		PublishedYear: 2023,
		CreatedAt:     time.Time{},
	}
	db.Create(&book)
	fmt.Println("Book record inserted successfully!")
	fmt.Println("Books in database:")

}
func CrudGetBook(db *gorm.DB) {
	book := Book{}
	db.First(&book, 5)
	fmt.Printf("ID: %d Title:%s ,Aythor:%s, Year:%d \n", book.ID, book.Title, book.Author, book.PublishedYear)
	//ID: 1, Title: Go Programming, Author: John Doe, Year: 2023
}
