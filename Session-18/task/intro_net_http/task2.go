package intro_net_http

import (
	"fmt"
	"net/http"
)

func GetTask2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Custom-Header", "Learning Go")
	fmt.Fprintln(w, "Hello, World!")
}
