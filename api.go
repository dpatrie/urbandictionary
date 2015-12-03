package urbandictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SearchResult struct {
	Type    string `json:"result_type"`
	Tags    []string
	Results []Result `json:"list"`
	Sounds  []string
}

type Result struct {
	Author     string
	Word       string
	Definition string
	Example    string
	Permalink  string
	Upvote     int `json:"thumbs_up"`
	Downvote   int `json:"thumbs_down"`
}

const API_URL = "http://api.urbandictionary.com/v0/define?term="

func Query(searchTerm string) (*SearchResult, error) {
	resp, err := http.Get(API_URL + url.QueryEscape(searchTerm))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Response was not a 200: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &SearchResult{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
