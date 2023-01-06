package WikiBot

import (
	"github.com/pkg/errors"

	mwApi "github.com/lynkfox/omnitron-go/WikiBot/mediaWikiApi"

	"cgt.name/pkg/go-mwclient"
	"github.com/antonholmquist/jason"
)

// Core contains the central parts of this WikiBot connections.
type Core struct {
	Username string
	Password string
	WikiUrl  string
	client   mwclient.Client
}

/* Connect generates a client and logs in, storing the client on the Core struct.
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

func (bot *Core) Close() error {
	err := bot.client.Logout()
	return err
}

/*
WikiGet sends a GET request to the wiki client logged in.
*/
func (bot *Core) WikiGet(parameters mwApi.QueryMapper) (*jason.Object, error) {
	resp, err := bot.client.Get(parameters.Map())
	if err != nil {
		return nil, errors.Wrap(err, "GetFailed")
	}
	return resp, nil
}

/*
GetRecentChanges is a wrapper around WikiGet for a default or custom RecentChanges query.

Pass nil to use default RecentChanges query of:

	Type filter to "edit"
	TopOnly to false
	Continue to be included
	Limit not set
	No start time
*/
func (bot *Core) GetRecentChanges(parameters *mwApi.RecentChanges) (*jason.Object, error) {
	if parameters == nil {
		var defaultParam mwApi.RecentChanges
		parameters = &defaultParam
	}

	resp, err := bot.WikiGet(parameters)
	if err != nil {
		return nil, errors.Wrap(err, "RecentChanges: %v")
	}

	/* TODO: Handle nextToken loop

	nextToken, err := resp.GetString("continue", "rcontinue")
	*/

	return resp, nil
}
