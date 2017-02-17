package main

import (
	"io"
	"net/http"
	"log"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

type demoHandle func(w http.ResponseWriter, req *http.Request)

func main() {
	//http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/hello",demoHandle(func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello, world!\n")
	}))
	log.Fatal(http.ListenAndServe(":12345", nil))
}