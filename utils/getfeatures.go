package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SongData struct {
	Acousticness     float64 `json:"acousticness"`
	AnalysisURL      string  `json:"analysis_url"`
	Danceability     float64 `json:"danceability"`
	DurationMS       int64   `json:"duration_ms"`
	Energy           float64 `json:"energy"`
	ID               string  `json:"id"`
	Instrumentalness float64 `json:"instrumentalness"`
	Key              int64   `json:"key"`
	Liveness         float64 `json:"liveness"`
	Loudness         float64 `json:"loudness"`
	Mode             int64   `json:"mode"`
	Speechiness      float64 `json:"speechiness"`
	Tempo            float64 `json:"tempo"`
	TimeSignature    int64   `json:"time_signature"`
	TrackHref        string  `json:"track_href"`
	Type             string  `json:"type"`
	URI              string  `json:"uri"`
	Valence          float64 `json:"valence"`
}

func GetSongFeatures(songId string, token string) (SongData, error) {

	url := "https://api.spotify.com/v1/audio-features/" + songId

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Bearer " + token)

	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	var songData SongData
	
	err := json.Unmarshal(body, &songData)

	if (err != nil){
		fmt.Printf("can't decode spotify response")
	}

	return songData, nil
}