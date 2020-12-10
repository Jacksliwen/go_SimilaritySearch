package main

import (
	"SimilaritySearch/searchengine"
	"fmt"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	searchengine.Init()
	w.Write([]byte("Reload OK!"))
}
func main() {
	http.HandleFunc("/Reload", DefaultHandler)
	srv := &http.Server{
		Addr: ":60018",
	}

	fmt.Println(srv.ListenAndServe())

}
