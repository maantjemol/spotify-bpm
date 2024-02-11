package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	spotify "github.com/maantjemol/spotify-bpm/api/spotify"
)

type Response struct {
	Tracks []spotify.RespTrack `json:"tracks"`
}


func Spotify(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	vals := r.URL.Query()

	playlistID := vals.Get("id")

	if (playlistID == ""){
		resp := make(map[string]string)
		resp["error"] = "missing playlist id"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}

	var resp Response 

	spotify_client := os.Getenv("SPOTIFY_CLIENT")
	spotify_secret := os.Getenv("SPOTIFY_SECRET")

	// Check if string exists
	if (spotify_client == "" || spotify_secret == ""){
		resp := make(map[string]string)
		resp["error"] = "missing client or secret"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}

	ACCESS_TOKEN, err := spotify.GetToken(spotify_client, spotify_secret)

	if (err != nil){
		resp := make(map[string][]spotify.Tracks)
		// resp["error"] = "can't get accesstoken"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}

	var playlist, playErr = spotify.GetPlaylist(playlistID, ACCESS_TOKEN)

	if (playErr != nil){
		fmt.Println(playErr)
	}

	resp.Tracks = make([]spotify.RespTrack, len(playlist))

	for i := 0; i < len(playlist); i++ {
		song := playlist[i]
		songFeatures, err := spotify.GetSongFeatures(song.ID, ACCESS_TOKEN)
		if (err != nil){
			println("couln't get id")
		}
		song.BPM = int(songFeatures.Tempo)
		song.Key = int(songFeatures.Key)
		song.Url = "https://open.spotify.com/track/" + song.ID
		song.TimeSignature = int(songFeatures.TimeSignature)
		resp.Tracks[i] = song
	}
	
	jsonResp, err := json.Marshal(resp)
	
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(jsonResp)
		return
	}

}