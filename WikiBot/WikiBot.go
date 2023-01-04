package WikiBot

import (
	"cgt.name/pkg/go-mwclient"
)

type Core struct {
	Username string
	Password string
	WikiUrl  string
	client   mwclient.Client
}

func (bot *Core) Connect() {
	// Initialize a *Client with New(), specifying the wiki's API URL
	// and your HTTP User-Agent. Try to use a meaningful User-Agent.
	client, err := mwclient.New(bot.WikiUrl, bot.Username)
	if err != nil {
		panic(err)
	}

	bot.client = *client

	// Log in.
	err = bot.client.Login(bot.Username, bot.Password)
	if err != nil {
		panic(err)
	}

}

/*	// Specify parameters to send.
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
*/
