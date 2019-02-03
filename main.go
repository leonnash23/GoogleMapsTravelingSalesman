package main

import (
	"MapApi/src/web"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", web.MainHandler)
	http.HandleFunc("/path", web.PathHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
