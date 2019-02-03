package main

import (
	"MapApi/src/web"
	"io"
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello!")
}

func main() {
	http.HandleFunc("/", web.MainHandler)
	http.HandleFunc("/path", web.PathHandler)
	http.HandleFunc("/test", Test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
