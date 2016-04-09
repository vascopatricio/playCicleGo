package main

import (
	"fmt"
	//"time"
)

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