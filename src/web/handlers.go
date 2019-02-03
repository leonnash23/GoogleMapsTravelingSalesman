package web

import (
	"MapApi/src/request"
	"html/template"
	"io"
	"net/http"
	"strings"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("main.html")
	t.Execute(w, nil)
}

func PathHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	mapRequest := request.MapRequest{}
	for _, s := range r.Form {
		for _, marker := range s {
			mapRequest.Markers.AddMarker(strings.TrimSpace(marker))
		}
	}
	resp, _ := http.Get(mapRequest.GetPathMapLink())
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}
