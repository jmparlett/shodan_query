//Author: Jonathan Parlett
// Program Name: shodan
// Created: 2021-December-27
package shodan

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type APIInfo struct {
	QueryCredits int    `json:"query_credits"`
	ScanCredits  int    `json:"scan_credits"`
	Telnet       bool   `json:"telnet"`
	Plan         string `json:"plan"`
	HTTPS        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
}

func (s *Client) APIInfo() (*APIInfo, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseUrl, s.apiKey))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var ret APIInfo
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *Client) PrintAPIInfo() {
	info, err := s.APIInfo()

	if err != nil {
		log.Fatal(err)
		return
	}
	//print info
	fmt.Printf("QueryCreds: %d\n", info.QueryCredits)
	fmt.Printf("ScanCreds: %d\n", info.ScanCredits)
	fmt.Printf("Telnet: %t\n", info.Telnet)
	fmt.Printf("Plan: %s\n", info.Plan)
	fmt.Printf("HTTPS: %t\n", info.HTTPS)
	fmt.Printf("Unlocked: %t\n", info.Unlocked)

}
