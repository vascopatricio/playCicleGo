package main

import (
	"fmt"
	"strconv"
	"strings"
	//"time"
	//"time"
)

type Playlist struct {
	Id int
	Title string

	Zones []PlaylistZone
}
func (f Playlist) GetTitle() string {
    return f.Title
}
func (f Playlist) GetId() int {
    return f.Id
}
func (f *Playlist) SetId(id int) {
    f.Id = id
}
func (f *Playlist) SetTitle(title string) {
    f.Title = title
}
func (f *Playlist) PrintSongsInOrder(){

	fmt.Println("> Playing PL ",f.GetTitle())

	for _, zone := range(f.Zones){
		fmt.Println(">> Playing ZONE: ",zone.GetTitle())

		currentSongTotal := 0
		currentMinsTotal := 0
		currentSecsTotal := 0

		zoneTotalMins, zoneTotalSecs := zone.ParseTimeLimit()

		songs := zone.Songs
		if zone.PlayOrder == "SHUFFLE" {
			songs = shuffleSongArray(songs)
		}

		for _, song := range(songs){

			if zone.LimitMode == "SONG_LIMIT" {
				// fmt.Println("Current mode is SONG LIMIT, checking if ",currentSongTotal,
				// 	" is over ",zone.SongLimit)
				if currentSongTotal > zone.SongLimit {
					//fmt.Println("It is - breaking")
					break
				}
			} else if zone.LimitMode == "TIME_LIMIT" && zoneTotalMins > 0 {
				// fmt.Println("Current mode is TIME LIMIT, checking if ",
				// 	currentMinsTotal,currentSecsTotal," is over ",zoneTotalMins,zoneTotalSecs)
				if currentMinsTotal > zoneTotalMins || 
					(currentMinsTotal == zoneTotalMins && currentSecsTotal >= zoneTotalSecs) {
					// fmt.Println("It is - breaking")
					break;
				}
			}

			currentSongTotal += 1
			currentMinsTotal += song.DurMins
			currentSecsTotal += song.DurSecs

			for {
				if currentSecsTotal < 60 {
					break
				}
				currentMinsTotal += int(currentSecsTotal / 60)
				currentSecsTotal = currentSecsTotal % 60
			}

			_ = strconv.Itoa(currentMinsTotal) + ":" + strconv.Itoa(currentSecsTotal)

			fmt.Println(">>>> Playing SONG ",song.GetTitle())
			// if zone.LimitMode == "SONG_LIMIT" {
			// 	fmt.Println("Zone total songs so far: ",currentSongTotal)
			// } else if zone.LimitMode == "TIME_LIMIT" {
			// 	fmt.Println("Zone total time so far: ",curTotalTimeStr)
			// }
		}
	}
}

type PlaylistZone struct {
	Id int
	Title string

	//	Can be one of 3 for now: NO_LIMIT, SONG_LIMIT, or TIME_LIMIT
	//	SONG_LIMIT plays XYZ number of songs
	//	TIME_LIMIT plays until XYZ time is reached
	LimitMode string
	//	Likewise, SongLimit will be the # of songs, 
	//	andTimeLimit the MM:SS time limit
	SongLimit int
	TimeLimit string

	//	Play order: Can be "Normal" or "Shuffle"
	PlayOrder string

	Songs []Song
}
func (f PlaylistZone) GetTitle() string {
    return f.Title
}
func (f PlaylistZone) GetId() int {
    return f.Id
}
func (z PlaylistZone) ParseTimeLimit() (int, int) {
    
	//fmt.Println("Trying to parse time limit from: ",z.TimeLimit)

	if z.LimitMode == "SONG_LIMIT"{
		return -1, -1
	}

	split := strings.Split(z.TimeLimit, ":")
	if len(split) != 2{
		return -1, -1
	}

	mins, _ := strconv.Atoi(split[0])
	secs, _ := strconv.Atoi(split[1])

	//fmt.Println("Returning: ",mins, secs)

	return mins, secs
}
func (f *PlaylistZone) SetId(id int) {
    f.Id = id
}

func (f *PlaylistZone) SetSettings(settings []string) {
    f.PlayOrder = settings[0]
    f.LimitMode = settings[1]

    if f.LimitMode == "SONG_LIMIT" {
    	f.SongLimit, _ = strconv.Atoi(settings[2])

    } else if f.LimitMode == "TIME_LIMIT"{
    	f.TimeLimit = settings[2]
    }
}
func (f *PlaylistZone) SetTitle(title string) {
    f.Title = title
}

type Song struct {
	Id int
	Title string
	Artist string

	Duration string
	DurMins int
	DurSecs int
}
func (f Song) GetTitle() string {
    return f.Title
}
func (f Song) GetId() int {
    return f.Id
}
func (f Song) GetArtist() string {
    return f.Artist
}
func (f Song) GetDuration() string {
    return f.Duration
}
func (f *Song) SetId(id int) {
    f.Id = id
}
func (f *Song) SetTitle(title string) {
    f.Title = title
}
func (f *Song) SetDuration(dur string) {
    f.Duration = dur
    duration := strings.Split(f.Duration,":")

    if len(duration) != 2 {
    	return;
    }

    f.DurMins, _ = strconv.Atoi(duration[0])
    f.DurSecs, _ = strconv.Atoi(duration[1])
}
func (f *Song) SetArtist(art string) {
    f.Artist = art
}