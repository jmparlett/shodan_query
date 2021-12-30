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

// {
type HostLocation struct {
	RegionCode  string  `json:"region_code"`  // "region_code": null,
	PostalCode  string  `json:"postal_code"`  // "postal_code": null,
	CountryCode string  `json:"country_code"` // "country_code": "US",
	City        string  `json:"city"`         // "city": null,
	CountryName string  `json:"country_name"` // "country_name": "United States",
	Lat         float32 `json:"latitude"`     // "latitude": 37.751,
	Long        float32 `json:"longitude"`    // "longitude": -97.822,
}

type Host struct {
	Ip         int64        `json:"ip"`          // "ip": 134744072,
	DmaCode    string       `json:"dma_code"`    // "dma_code": null,
	LastUpdate string       `json:"last_update"` // "last_update": "2021-01-22T08:49:35.190817",
	Hostnames  []string     `json:"hostnames"`   // "hostnames": ["dns.google"],
	Org        string       `json:"org"`         // "org": "Google",
	Location   HostLocation `json:"location"`
	// Data       []string     `json:"data"` // "data": [
	Port     int      `json:"port"`
	IPString string   `json:"ip_str"`
	Ports    []int    `json:"ports"`
	Os       string   `json:"os"`
	Domains  []string `json:"domains"`
	Asn      string   `json:"asn"`
	Isp      string   `json:"isp"`
}

type HostSearch struct {
	Matches []Host `json:"matches"`
}

func (s *Client) Host(hostIp string) (*Host, error) {
	//https://api.shodan.io/shodan/host/{ip}?key=P6z6oEIao8SkbN0TqGodLxwhVQHqRhsm

	resp, err := http.Get(fmt.Sprintf("%s/shodan/host/%s?key=%s", BaseUrl, hostIp, s.apiKey))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ret Host
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *Client) PrintHost(hostIp string) {
	host, err := s.Host(hostIp)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("IP: %s\n", host.IPString)
	fmt.Printf("Hostnames: %v\n", host.Hostnames)
	fmt.Printf("Domains: %v\n", host.Domains)
	fmt.Printf("ORG: %s\n", host.Org)
	fmt.Printf("OS: %s\n", host.Os)
	fmt.Printf("Ports: %v\n", host.Ports)
	fmt.Printf("ASN: %s\n", host.Asn)
	fmt.Printf("ISP: %s\n", host.Isp)
}
