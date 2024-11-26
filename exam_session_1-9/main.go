package main

import (
	"fmt"
)

type Book struct {
	BookID       int
	Title        string
	Author       string
	Availability string
}
type User struct {
	UserID        int
	Name          string
	BorrowedBooks []int
}

var users = []*User{
	{UserID: 1, Name: "Ali"},
	{UserID: 2, Name: "Babek"},
	{UserID: 3, Name: "Cabir"},
	{UserID: 4, Name: "Diana"},
	{UserID: 5, Name: "Elvira"},
}

var books = []*Book{
	{BookID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Availability: "available"},
	{BookID: 2, Title: "Learn Go in One Day", Author: "Alex Brown", Availability: "available"},
	{BookID: 3, Title: "Go Web Programming", Author: "Sau Sheong Chang", Availability: "available"},
	{BookID: 4, Title: "Go in Action", Author: "William Kennedy", Availability: "available"},
	{BookID: 5, Title: "Go Programming Blueprints", Author: "Mat Ryer", Availability: "available"},
	{BookID: 6, Title: "Go Concurrency Patterns", Author: "Katherine Cox-Buday", Availability: "available"},
	{BookID: 7, Title: "Go Design Patterns", Author: "Mitchell Hashimoto", Availability: "available"},
	{BookID: 8, Title: "Go Recipes", Author: "Shannon Smith", Availability: "available"},
	{BookID: 9, Title: "Mastering Go", Author: "Mihalis Tsoukalos", Availability: "available"},
	{BookID: 10, Title: "Go Systems Programming", Author: "Vladimir Vivien", Availability: "available"},
}

func main() {
	err := borrowBook(1, 2)
	if err != nil {
		fmt.Println(err)
	}
	err1 := borrowBook(1, 4)
	if err1 != nil {
		fmt.Println(err1)
	}
	err3 := borrowBook(1, 3)
	if err3 != nil {
		fmt.Println(err3)
	}

	errReurun := returnBook(1, 2)
	if errReurun != nil {
		fmt.Println(errReurun)
	}

	aviableBooks := listAvailableBooks()
	fmt.Println("All Aviable books")
	for _, b := range aviableBooks {
		fmt.Println(*b)
	}
	fmt.Println("Returns a list of books borrowed by a specific user")
	us := viewBorrowedBooks(2)
	for _, v := range us {
		fmt.Println(*v)
	}

}

func borrowBook(userID, bookID int) error {
	var user *User
	var book *Book
	for _, u := range users {
		if u.UserID == userID {
			user = u
		}
	}
	if len(user.BorrowedBooks) > 3 {
		return fmt.Errorf("your have borroed limit 3")
	}

	for _, b := range books {
		if b.BookID == bookID {
			book = b
		}
	}
	if book.Availability != "available" {
		return fmt.Errorf("your have book availability %s", book.Availability)
	} else if book.Availability == "available" {
		user.BorrowedBooks = append(user.BorrowedBooks, bookID)
		book.Availability = "borrowed"
	}

	return nil
}

func returnBook(userID int, bookID int) error {
	var user *User
	var book *Book

	for _, u := range users {
		if u.UserID == userID {
			user = u
		}

	}
	for _, b := range books {
		if b.BookID == bookID {
			book = b

		}
	}
	book.Availability = "available"
	for k, b := range user.BorrowedBooks {
		if b == bookID {
			user.BorrowedBooks = append(user.BorrowedBooks[:k], user.BorrowedBooks[k+1:]...)
		}
	}

	return nil
}

func listAvailableBooks() []*Book {
	var AviableBooks []*Book
	for _, v := range books {
		if v.Availability == "available" {
			AviableBooks = append(AviableBooks, v)
		}
	}
	return AviableBooks
}
func viewBorrowedBooks(userID int) []*Book {
	//var book *Book
	var borrowedBooks []*Book
	var user *User

	for _, u := range users {
		if u.UserID == userID {
			user = u
		}
	}
	if len(user.BorrowedBooks) == 0 {
		fmt.Println("no Borroed books")
	}

	for _, bookID := range user.BorrowedBooks {
		for _, b := range books {
			if b.BookID == bookID {
				borrowedBooks = append(borrowedBooks, b)
			}
		}
	}

	return borrowedBooks
}
