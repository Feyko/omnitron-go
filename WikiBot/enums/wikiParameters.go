package wikiApi

type Action string

const (
	NilAction  Action = ""
	Query      Action = "query"
	Edit       Action = "edit"
	Parse      Action = "parse"
	ShortenUrl Action = "shortenurl"
	Tag        Action = "tag"
)

type Format string

const (
	Json       Format = "json"
	PrettyJson Format = "jsonfm"
)

type Assert string

const (
	Bot  Assert = "bot"
	Anon Assert = "anon"
	User Assert = "user"
)

type List string

const (
	RecentChanges List = "recentchanges"
)
