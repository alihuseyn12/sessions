package web_serv

import (
	"fmt"
	"net/http"
)

func HandWelcame(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to the homepage!")
}
func HandAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the about page")
}
func HandContact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Contact us at contact@example.com.")
}
