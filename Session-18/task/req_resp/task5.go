package req_resp

import (
	"fmt"
	"net/http"
)

func HandlGrateName(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Println("Hello, Stranger!")
	} else {
		fmt.Println("Hello ," + name)
	}

}
