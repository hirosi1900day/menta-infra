package main

import (
	"fmt"
	"net/http"
)

func okHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	http.HandleFunc("/", okHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
