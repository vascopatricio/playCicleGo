package main

import (
	"fmt"
	"io/ioutil"
	"html/template"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {

	contents, err := ioutil.ReadFile("../frontend/assets/index.html")
	if err != nil {
		fmt.Println("Error:",err)
	}

	tmpl := template.New("page")
	tmpl, _ = tmpl.Parse(string(contents))
	tmpl.Execute(w, nil)

	//fmt.Fprintf(w, "Hello world!")
}

func playlistTestsView(w http.ResponseWriter, r *http.Request) {

	DATASETS_ROOT := "../datasets/"

	//_ = getSamplePlaylist()
	plEdm, _ := loadPlaylistFromFile(DATASETS_ROOT+"/testEdmPlaylist.txt")
	plTest, _ := loadPlaylistFromFile(DATASETS_ROOT+"/testPlaylist.txt")

	//printPlaylist(pl);
	//printPlaylist(plEdm);
	//printPlaylist(plTest);

	//pl.PrintSongsInOrder()
	plEdm.PrintSongsInOrder()
	plTest.PrintSongsInOrder()

	playlists := []Playlist{plEdm, plTest}

	err := servePlaylistsJsonResponse(w, playlists)
	if err != nil {
		fmt.Printf("Error obtaining playlists: ",err)
		return
	}
}