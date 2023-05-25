package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed static/index.html
var index string

//go:embed static/taskmaster.png
var png string

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/taskmaster.png", servePng)
	mux.HandleFunc("/", serveIndex)
	mux.HandleFunc("/*", serveIndex)
	fmt.Println("listening on port 5555")
	http.ListenAndServe(":5555", mux)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html; charset=UTF-8")
	w.Write([]byte(index))
}

func servePng(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "image/png")
	w.Write([]byte(png))
}
