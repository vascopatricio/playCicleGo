package main

import (
	"fmt"
	"strconv"
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

	fmt.Println("Playing playlist: ",f.GetTitle())

	for _, zone := range(f.Zones){
		fmt.Println("Playing: Current Zone: ",zone.GetTitle())
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
func (f *Song) SetId(id int) {
    f.Id = id
}
func (f *Song) SetTitle(title string) {
    f.Title = title
}
func (f *Song) SetArtist(title string) {
    f.Title = title
}