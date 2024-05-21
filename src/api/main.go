package main

import (
	"fmt"
	"net/http"

	createUser "github.com/jorgeloch/expenses-tracker/src/api/handlers/user/create_user"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, world!"))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &handler{})
	mux.Handle("/users", &createUser.Handler{})

	http.ListenAndServe(":8080", mux)
	fmt.Println("Server running on port 8080")
}
