package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "view/index.html")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		panic(err)
	}
}
