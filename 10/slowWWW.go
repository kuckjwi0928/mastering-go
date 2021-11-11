package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func generateRandomValue(min, max int) int {
	return rand.Intn(max-min) + min
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	delay := generateRandomValue(0, 15)
	time.Sleep(time.Duration(delay) * time.Second)

	_, _ = fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	_, _ = fmt.Fprintf(w, "Delay: %d\n", delay)
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	http.HandleFunc("/", myHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(10)
	}
}
