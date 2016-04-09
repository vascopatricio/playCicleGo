package main

import (
	"bufio"
	"log"
    "errors"
	"os"
	"strings"
	"math/rand"
	"fmt"
	//"time"
)

func loadPlaylistFromFile(filename string) (Playlist, error){

	pl := Playlist{}
	pl.SetId(rand.Intn(1000))

	plZones := make([]PlaylistZone, 0)

	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    lineCounter := 0
    readingZoneMode := ""
	newZone := PlaylistZone{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	lineCounter += 1
    	curLine := scanner.Text()
    	lineLen := len(curLine)

    	if lineCounter == 1{
    		pl.SetTitle(curLine)

    	} else if readingZoneMode == "stats" {
            lineSplit := strings.Split(curLine," ")
    		fmt.Println("Zone settings: ",lineSplit)
    		newZone.SetSettings(lineSplit)
    		readingZoneMode = "songs"

    	} else if readingZoneMode == "songs" && lineLen > 0 {
            lineSplit := strings.Split(curLine," - ")
            fmt.Println("Adding song: ",lineSplit)

            if len(lineSplit) != 3{
                return pl, errors.New("Invalid song in line: "+curLine)
            }

            newSong := Song{}
            newSong.SetTitle(lineSplit[0])
            newSong.SetArtist(lineSplit[1])
            newSong.SetDuration(lineSplit[2])
            newZone.Songs = append(newZone.Songs, newSong)

    	} else if readingZoneMode == "songs" && lineLen == 0{
    		plZones = append(plZones, newZone)
    		fmt.Println("Finalizing zone")
    		readingZoneMode = ""

    	} else if readingZoneMode == "" {
    		newZone = PlaylistZone{}
    		newZone.SetId(rand.Intn(1000))
    		newZone.SetTitle(curLine)
            newZone.Songs = make([]Song, 0)

    		fmt.Println("Initializing zone with title: ", newZone.GetTitle())
    		readingZoneMode = "stats"
    	
    	} 
        //fmt.Println("Line with len ",lineLen,":",lineCounter,scanner.Text())
    }

    //	After reading all lines, if there's no empty line at the end, don't forget to close section
    if readingZoneMode == "songs" {
		plZones = append(plZones, newZone)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Zones length: ",len(plZones))

    pl.Zones = plZones

	return pl, nil
}

func getSamplePlaylist() (Playlist){

	plZones := make([]PlaylistZone, 0)
	
	letItGo := Song{1, "Let It Go", "James Bay", "", 0, 0}
	adv := Song{2, "Adventure of a Lifetime", "Coldplay", "", 0, 0}
	exs := Song{3, "Ex's & Oh's", "Elle King", "", 0, 0}

	plSongs := make([]Song, 0)
	plSongs = append(plSongs, letItGo)
	plSongs = append(plSongs, adv)
	plSongs = append(plSongs, exs)

	zone1 := PlaylistZone{1, "Warm-Up", "SONG_LIMIT", 3, "", "NORMAL", plSongs}

	plZones = append(plZones, zone1)

	pl := Playlist{1, "Test Playlist", plZones}

	return pl
}