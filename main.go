package main

import (
	"fmt"
	"os"

	"github.com/lynkfox/omnitron-go/WikiBot/environment"
	mwApi "github.com/lynkfox/omnitron-go/WikiBot/mediaWikiApi"
)

func main() {

	bot, err := environment.InitEnvironment()
	defer bot.Close()
	if err != nil {
		fmt.Print("error on env setup")
		fmt.Print(err)
		os.Exit(1)
	}

	err = bot.Connect()

	if err != nil {
		fmt.Print("error on connect")
		fmt.Print(err)
		os.Exit(1)
	}

	var parse mwApi.Parse
	parse.Page = "Push"
	parseResp, err := bot.WikiGet(parse)

	//parseResp.GetString()

	if err != nil {
		fmt.Print("error on wikiget")
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Print(parseResp)
	os.Exit(0)
}
