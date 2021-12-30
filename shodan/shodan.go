//Author: Jonathan Parlett
// Program Name: shodan
// Created: 2021-December-27
// Purpose: provides constants and structs for use in calls to the shodan api
package shodan

const BaseUrl = "http://api.shodan.io" //remember anything public should be captilized

type Client struct {
	apiKey string
}

func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
