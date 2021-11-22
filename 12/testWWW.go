package main

import (
	"fmt"
	"net/http"
)

func CheckStatusOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, `Fine!`)
}

func StatusNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	PORT := ":8001"

	r := http.NewServeMux()
	r.HandleFunc("/StatusNotFound", StatusNotFound)
	r.HandleFunc("/CheckStatusOK", CheckStatusOK)
	r.HandleFunc("/", Handler)

	err := http.ListenAndServe(PORT, r)

	if err != nil {
		fmt.Println(err)
		return
	}
}
