package main

import (
	//"net/http"
	//"fmt"
)

func main() {

	pl := getSamplePlaylist()
	plEdm := loadPlaylistFromFile("datasets/testEdmPlaylist.txt")
	plTest := loadPlaylistFromFile("datasets/testPlaylist.txt")

	printPlaylist(pl);
	printPlaylist(plEdm);
	printPlaylist(plTest);

	//pl.PrintSongsInOrder()
	//plEdm.PrintSongsInOrder()
	//plTest.PrintSongsInOrder()

  	//	Frontend - HTML
	//http.HandleFunc("/", handler)

	//http.ListenAndServe(":3000", nil)
}