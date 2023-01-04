package environment

import (
	"fmt"
	"os"

	"github.com/pkg/errors"

	"github.com/lynkfox/omnitron-go/bot"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

/*
	GetEnvValue returns an environment value from a given key.

This is wrapped for ease of use in regards to errors and for anticipation of
transferring from .env files to a more secure method, such as parameter store
or secrets manager
*/
func getEnvValue(key string) (string, error) {

	val := os.Getenv(key)

	if val == "" {
		return val, errors.New(fmt.Sprintf("EnvironmentKey: %s", key))
	}

	return val, nil
}

/* LoadEnv is a temp function to load .env file. Expected to be replaced when
moving to more secure method of env variables.

Performs some basic checking for necessary values.
*/

func loadEnv() error {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Could not load .env file.")
		return err
	}

	return nil
}

/* InitEnvironment will get all the necessary values for the bot and return
them as a BotOperatingValues object
*/

func InitEnvironment() (bot.WikiBot, error) {

	var bot bot.WikiBot

	err := loadEnv()
	if err != nil {
		return bot, err
	}

	envValues := map[string]*string{
		"WIKI_URL":      &bot.wikiUrl,
		"WIKI_USERNAME": &bot.username,
		"WIKI_PASSWORD": &bot.password,
	}
	for k, v := range envValues {
		envValue, err := getEnvValue(k)
		if err != nil {
			return bot, errors.Wrap(err, "Missing value")
		}
		*v = envValue
	}
	return bot, nil
}
