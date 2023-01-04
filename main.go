package main

import (
	"fmt"

	"cgt.name/pkg/go-mwclient"
)

func main() {
	// Initialize a *Client with New(), specifying the wiki's API URL
	// and your HTTP User-Agent. Try to use a meaningful User-Agent.
	w, err := mwclient.New("https://en.wikipedia.org/w/api.php", "Omnitron-Go")
	if err != nil {
		panic(err)
	}

	// Log in.
	err = w.Login("USERNAME", "PASSWORD")
	if err != nil {
		panic(err)
	}

	// Specify parameters to send.
	parameters := map[string]string{
		"action":   "query",
		"list":     "recentchanges",
		"rclimit":  "2",
		"rctype":   "edit",
		"continue": "",
	}

	// Make the request.
	resp, err := w.Get(parameters)
	if err != nil {
		panic(err)
	}

	// Print the *jason.Object
	fmt.Println(resp)
}
