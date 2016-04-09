package main

import (
	"time"
	"math/rand"
	"fmt"
	//"time"
)

func shuffleSongArray(songs []Song) ([]Song){
	rand.Seed(time.Now().UnixNano())

	//fmt.Println("Array before shuffle: ",songs)

	for i := range songs {
	    j := rand.Intn(i + 1)
	    songs[i], songs[j] = songs[j], songs[i]
	}

	//fmt.Println("Array after shuffle: ",songs)

	return songs
}

func printPlaylist(p Playlist){

	fmt.Println("Playlist: ", p.GetId(), p.GetTitle())

	for _, zone := range p.Zones {
		fmt.Println("Zone: ",zone.GetId(), zone.GetTitle(), 
			zone.LimitMode, zone.TimeLimit, zone.SongLimit, 
			zone.PlayOrder, len(zone.Songs), "Songs")

		for _, element := range zone.Songs {
			fmt.Println("Song: ",element.GetTitle())
		}
	}

}