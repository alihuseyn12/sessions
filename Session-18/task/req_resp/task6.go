package req_resp

import (
	"fmt"
	"net/http"
	"strconv"
)

func HandlDivide(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))
	fmt.Println(a, b)
	if b == 0 {
		http.Error(w, "Division by zero is not allowed", http.StatusBadRequest)
		fmt.Println("400", http.StatusText(http.StatusBadRequest), "Division by zero is not allowed.")
	} else {

		fmt.Println("Result:", a/b)
	}
}
