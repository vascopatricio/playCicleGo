package main

import (
	"net/http"
	//"fmt"
)

func main() {

  	//	Frontend - HTML
	http.HandleFunc("/", mainPage)

	http.HandleFunc("/tests/", playlistTestsView)
	
	http.ListenAndServe(":3000", nil)
}