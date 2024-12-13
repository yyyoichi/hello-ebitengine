package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	api := http.NewServeMux()
	api.HandleFunc("/api/echo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Hello, World!! %s", time.Now().Format(time.RFC3339))))
		w.WriteHeader(http.StatusOK)
	})
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.Handle("/api/", api)
	http.ListenAndServe(":3000", mux)
}
