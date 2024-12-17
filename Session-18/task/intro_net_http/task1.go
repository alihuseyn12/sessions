package intro_net_http

import (
	"fmt"
	"net/http"
)

func GetTask1(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Hello, World!")
}
