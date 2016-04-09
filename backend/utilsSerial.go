package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type PlaylistsResponse struct {
	Playlists []Playlist
    Successful string
}

func servePlaylistsJsonResponse(w http.ResponseWriter, playlists []Playlist) (error) {

	response := PlaylistsResponse{playlists, "true"}

	res, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error mashaling JSON topics",err)
		return err
	}

	return serveGenericJsonResponse(w, res)
}

func serveGenericJsonResponse(w http.ResponseWriter, res []byte) error{

	w.Header().Set("Content-Type", "application/javascript")
	fmt.Fprintf(w, string(res))

	return nil
}