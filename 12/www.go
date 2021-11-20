package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	_, _ = fmt.Fprintf(w, "<h1 align=center>%s</h1>", Body)
	_, _ = fmt.Fprintf(w, "<h2 align=center>%s</h1>", t)
	_, _ = fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	PORT := ":8001"

	r := http.NewServeMux()
	r.HandleFunc("/time", timeHandler)
	r.HandleFunc("/", myHandler)
	r.HandleFunc("/debug/pprof", pprof.Index)
	r.HandleFunc("/debug/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/profile", pprof.Profile)
	r.HandleFunc("/debug/trace", pprof.Trace)

	err := http.ListenAndServe(PORT, r)

	if err != nil {
		fmt.Println(err)
		return
	}
}
