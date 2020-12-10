package http

import (
	"SimilaritySearch/searchengine"
	"net/http"
)

//Reload Reload
func Reload(w http.ResponseWriter, req *http.Request) {
	searchengine.Reload("myok")
	w.Write([]byte("Reload OK!"))
}

//Addid Addid
func Addid(w http.ResponseWriter, r *http.Request) {
	searchengine.Addid("myok", "myok")
	w.Write([]byte("Addid OK!"))
}

//Search Search
func Search(w http.ResponseWriter, r *http.Request) {
	searchengine.Search()
	w.Write([]byte("Search OK!"))
}

func Unload(w http.ResponseWriter, req *http.Request) {
	searchengine.Reload("myok")
	w.Write([]byte("Reload OK!"))
}
