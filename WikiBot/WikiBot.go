package WikiBot

import (
	"fmt"

	"cgt.name/pkg/go-mwclient"
	"github.com/lynkfox/omnitron-go/WikiBot/wikiApi"
)

type Core struct {
	Username string
	Password string
	WikiUrl  string
	client   mwclient.Client
}

func (bot *Core) Connect() {
	client, err := mwclient.New(bot.WikiUrl, bot.Username)
	if err != nil {
		panic(err)
	}

	bot.client = *client

	err = bot.client.Login(bot.Username, bot.Password)
	if err != nil {
		panic(err)
	}

}

func (bot *Core) GetRecentChanges() {
	parameters := map[string]string{
		"action":   wikiApi.Query,
		"list":     wikiApi.RecentChanges,
		"rclimit":  "2",
		"rctype":   "edit",
		"continue": "",
	}

	resp, err := w.Get(parameters)
	if err != nil {
		panic(err)
	}

	// Print the *jason.Object
	fmt.Println(resp)
}
