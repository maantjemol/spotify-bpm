package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type Playlist struct {
	Collaborative bool         `json:"collaborative"`
	Description   string       `json:"description"`
	ExternalUrls  ExternalUrls `json:"external_urls"`
	Followers     Followers    `json:"followers"`
	Href          string       `json:"href"`
	ID            string       `json:"id"`
	Images        []Image      `json:"images"`
	Name          string       `json:"name"`
	Owner         Owner        `json:"owner"`
	PrimaryColor  interface{}  `json:"primary_color"`
	Public        bool         `json:"public"`
	SnapshotID    string       `json:"snapshot_id"`
	Tracks        Tracks       `json:"tracks"`
	Type          string       `json:"type"`
	URI           string       `json:"uri"`
}

type ExternalUrls struct {
	Spotify string `json:"spotify"`
}

type Followers struct {
	Href  interface{} `json:"href"`
	Total int64       `json:"total"`
}

type Image struct {
	Height *int64 `json:"height"`
	URL    string `json:"url"`
	Width  *int64 `json:"width"`
}

type Owner struct {
	DisplayName  *string      `json:"display_name,omitempty"`
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Type         Type         `json:"type"`
	URI          string       `json:"uri"`
	Name         *string      `json:"name,omitempty"`
}

type Tracks struct {
	Href     string      `json:"href"`
	Items    []Item      `json:"items"`
	Limit    int64       `json:"limit"`
	Next     interface{} `json:"next"`
	Offset   int64       `json:"offset"`
	Previous interface{} `json:"previous"`
	Total    int64       `json:"total"`
}

type Item struct {
	AddedAt        string         `json:"added_at"`
	AddedBy        Owner          `json:"added_by"`
	IsLocal        bool           `json:"is_local"`
	PrimaryColor   interface{}    `json:"primary_color"`
	Track          Track          `json:"track"`
	VideoThumbnail VideoThumbnail `json:"video_thumbnail"`
}

type Track struct {
	Album            Album        `json:"album"`
	Artists          []Owner      `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int64        `json:"disc_number"`
	DurationMS       int64        `json:"duration_ms"`
	Episode          bool         `json:"episode"`
	Explicit         bool         `json:"explicit"`
	ExternalIDS      ExternalIDS  `json:"external_ids"`
	ExternalUrls     ExternalUrls `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	IsLocal          bool         `json:"is_local"`
	Name             string       `json:"name"`
	Popularity       int64        `json:"popularity"`
	PreviewURL       *string      `json:"preview_url"`
	Track            bool         `json:"track"`
	TrackNumber      int64        `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
}

type Album struct {
	AlbumType            string       `json:"album_type"`
	Artists              []Owner      `json:"artists"`
	AvailableMarkets     []string     `json:"available_markets"`
	ExternalUrls         ExternalUrls `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Image      `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	TotalTracks          int64        `json:"total_tracks"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
}

type ExternalIDS struct {
	Isrc string `json:"isrc"`
}

type VideoThumbnail struct {
	URL interface{} `json:"url"`
}

type Type string

const (
	Artist Type = "artist"
	User   Type = "user"
)

type RespTrack struct {
	ID string `json:"id"`
	BPM int `json:"BPM"`
	Key int `json:"key"`
	Url string `json:"url"`
	TimeSignature int `json:"timeSignature"`
	PreviewURL string `json:"previewURL"`
	Image string `json:"image"`
	Name string `json:"name"`
	ArtistName string `json:"artistName"`
}

func GetPlaylist(playlistID string, token string) ([]RespTrack, error) {

	url := "https://api.spotify.com/v1/playlists/" + playlistID

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Bearer " + token)

	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	var playlistData Playlist
	
	err := json.Unmarshal(body, &playlistData)

	if (err != nil){
		fmt.Println("can't decode spotify response")
	}

	amountSongs := len(playlistData.Tracks.Items)

	songIds := make([]RespTrack, amountSongs)
	
	for i :=  0; i < amountSongs; i++ {
		track := playlistData.Tracks.Items[i].Track
		previewURL := ""
		if track.PreviewURL != nil {
			previewURL = *track.PreviewURL
		}
		artistName := ""
		if len(track.Artists) >  0 && track.Artists[0].Name != nil {
			artistName = *track.Artists[0].Name
		}
		songIds[i] = RespTrack{
			ID: track.ID,
			PreviewURL: previewURL,
			Image: track.Album.Images[0].URL,
			Name: track.Name,
			ArtistName: artistName,
		}
	}

	return songIds, nil
}