package main

import (
	"github.com/MisakiFx/Golang/http/server"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.Ping)
	http.HandleFunc("/counter/", server.Counter)
	http.ListenAndServe(":8080", nil)
}
