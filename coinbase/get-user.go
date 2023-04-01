package coinbase

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	userEndpoint = "https://api.coinbase.com/v2/user"
)

func GetUser() (user any, err error) {
	req, _ := http.NewRequest("GET", userEndpoint, nil)
	err = signRequest(req)
	if err != nil {
		return
	}

	c := http.Client{}
	response, err := c.Do(req)
	if err != nil {
		return
	}

	// read body from response into a bytes array
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var cbResponse Response[User]

	// unmarshal the bytes array into a map
	err = json.Unmarshal(bytes, &cbResponse)
	if err != nil {
		return
	}

	user = cbResponse.Data
	return
}

type User struct {
	Id              string  `json:"id"`
	Name            string  `json:"name"`
	Username        *string `json:"username"`
	ProfileLocation *string `json:"profile_location"`
	ProfileBio      *string `json:"profile_bio"`
	ProfileUrl      *string `json:"profile_url"`
	AvatarUrl       string  `json:"avatar_url"`
	Resource        string  `json:"resource"`
	ResourcePath    string  `json:"resource_path"`
}
