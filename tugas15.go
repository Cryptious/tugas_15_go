package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	for i := 1; i <= 100; i++ {
		fmt.Fprintln(w, i)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, " Hello, Selamat Datang")
	})
	http.HandleFunc("/index", index)
	fmt.Println("Memulai Web Server Pada Localhost:8080")
	http.ListenAndServe(":8080", nil)
}
