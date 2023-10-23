package main

import (
	"fmt"
	"net/http"
)

func main() {
	handledRequests()
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "God i thank YOu")

}
func handledRequests() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":8082", nil)

}
