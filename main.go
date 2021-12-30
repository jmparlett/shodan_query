//Author: Jonathan Parlett
// Program Name: shodan
// Created: 2021-December-27
package main

import (
	"flag"
	"fmt"
	"shodan/shodan"
)

const apiKey string = "{Your API Key Here}"

func main() {

	//params
	var host string //ip of host to query shodan about
	var info bool   //switch to print api-info and exit

	flag.StringVar(&host, "s", "8.8.8.8", "server to query shodan about")
	flag.BoolVar(&info, "i", false, "print api-info and exit")

	flag.Parse()

	client := shodan.New(apiKey)

	if info {
		fmt.Println("\n--------API-INFO--------")
		client.PrintAPIInfo()
		fmt.Println()
		return
	}

	//print host specified
	fmt.Printf("\n--------Info-On-%s--------\n", host)
	client.PrintHost(host)
	fmt.Println()
}
