package http

import (
	"SimilaritySearch/searchengine"
	"net/http"
)

//Reload Reload
func Reload(w http.ResponseWriter, req *http.Request) {
	searchengine.InitEngine("myok")
	w.Write([]byte("Reload OK!"))
}

//Addid Addid
func Addid(w http.ResponseWriter, r *http.Request) {
	var ff *float32
	searchengine.LoadData("myok", ff, 1)
	w.Write([]byte("Addid OK!"))
}

//Search Search
func Search(w http.ResponseWriter, r *http.Request) {
	searchengine.Search()
	w.Write([]byte("Search OK!"))
}

//Unload Unload
func Unload(w http.ResponseWriter, req *http.Request) {
	searchengine.InitEngine("myok")
	w.Write([]byte("Unload OK!"))
}
