package WikiBot

import (
	"github.com/pkg/errors"

	"cgt.name/pkg/go-mwclient"
	"github.com/antonholmquist/jason"
)

type Core struct {
	Username string
	Password string
	WikiUrl  string
	client   mwclient.Client
}

/* Connect generates a client and logs in, storing the client on the Core struct
 */
func (bot *Core) Connect() error {
	client, err := mwclient.New(bot.WikiUrl, bot.Username)
	if err != nil {
		return errors.Wrap(err, "Could not generate Client")
	}

	bot.client = *client

	err = bot.client.Login(bot.Username, bot.Password)
	if err != nil {
		return errors.Wrap(err, "Failed Login")
	}

	return nil

}

/*
GetRecentChanges queries the Recent Changes List from the api and returns a
list of the changes, or nil and an error if it fails.
*/
func (bot *Core) GetRecentChanges() (*jason.Object, error) {
	parameters := map[string]string{
		"action":   "query",
		"list":     "recentchanges",
		"rclimit":  "2",
		"rctype":   "edit",
		"continue": "",
	}

	resp, err := bot.client.Get(parameters)
	if err != nil {
		return nil, errors.Wrap(err, "Recent Changes failed to get")
	}

	return resp, nil
}
