package main

import (
	httpProcess "SimilaritySearch/http"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/Reload", httpProcess.Reload)
	http.HandleFunc("/Unload", httpProcess.Unload)
	http.HandleFunc("/Addid", httpProcess.Addid)
	http.HandleFunc("/Search", httpProcess.Search)

	srv := &http.Server{
		Addr: ":60018",
	}

	fmt.Println(srv.ListenAndServe())

}
