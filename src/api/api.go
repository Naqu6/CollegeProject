package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SessionInformation struct {
	Add     bool   `json:"add"`
	Value   string `json:"value"`
	Expires string `json:"expires"`
	Name    string `json:"name"`
}

type APIResponse struct {
	Result   string                 `json:"result"`
	Data     map[string]interface{} `json:"data"`
	Sessions []SessionInformation   `json:"sessions"`
}

func ContactAPI(w http.ResponseWriter, r *http.Request) {
	info := r.URL.Query()

	for _, cookie := range r.Cookies() {
		info[cookie.Name] = []string{cookie.Value}
	}

	resp, err := http.PostForm("http://localhost:8000/api/", info)

	if err != nil {
		fmt.Printf("Error Contacting API Server %s", err)
	}

	var bytes []byte
	bytes, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error Reading API Response %s", err)
	}

	apiResponse := APIResponse{}

	err = json.Unmarshal(bytes, &apiResponse)

	if err != nil {
		fmt.Printf("Error Parsing API Data %s\n", err)
	}

	for _, sessionInfo := range apiResponse.Sessions {

		var cookie http.Cookie

		if sessionInfo.Add {
			expires, err := time.Parse(time.RFC3339, sessionInfo.Expires)

			if err != nil {
				fmt.Printf("Error Parsing API Resposne Session Time %s", err)
			}

			cookie = http.Cookie{Name: sessionInfo.Name, Value: sessionInfo.Value, Expires: expires}
		} else {
			cookie = http.Cookie{Name: sessionInfo.Name, Value: "", Expires: time.Now()}
		}

		http.SetCookie(w, &cookie)
	}

	w.Write(bytes)
}
