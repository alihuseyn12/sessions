package web_serv

import (
	"fmt"
	"net/http"
)

func HandlLog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method, ",", "URL:", r.URL)
	w.Write([]byte("Request logged."))
}
