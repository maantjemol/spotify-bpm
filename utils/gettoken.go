package utils

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

func GetToken(client string, secret string) (string, error) {
	sEnc := b64.StdEncoding.EncodeToString([]byte(client + ":" + secret))

	url := "https://accounts.spotify.com/api/token"

	payload := strings.NewReader("grant_type=client_credentials")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", "Basic " + sEnc)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	
	defer res.Body.Close()

	var accessToken AccessToken
	
	err := json.Unmarshal(body, &accessToken)

	if (err != nil){
		fmt.Println("can't decode spotify response")
	}

	if (accessToken.AccessToken == ""){
		return "", errors.New("accesstoken empty")
	}

	return accessToken.AccessToken, nil
}