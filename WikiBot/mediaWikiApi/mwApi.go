package mwApi

import (
	"reflect"
	"strconv"
	"strings"
)

type Property string

const (
	queryAction            Property = "query"
	editAction             Property = "edit"
	parseAction            Property = "parse"
	botAssert              Property = "bot"
	jsonFormat             Property = "json"
	queryListRecentChanges Property = "recentchanges"
)

type RecentChangesProp string

const (
	rcEditType RecentChangesProp = "edit"
	rcNewType  RecentChangesProp = "new"
	rcLimitMax RecentChangesProp = "max"
)

/*

	NilAction  Action = ""
	Query      Action = "query"
	EditPage   Action = "edit"
	Parse      Action = "parse"
	ShortenUrl Action = "shortenurl"
	Tag        Action = "tag"


	Json       Format = "json"
	PrettyJson Format = "jsonfm"

	Bot  Assert = "bot"
	Anon Assert = "anon"
	User Assert = "user"

	RecentChanges List = "recentchanges"

	RcEditFilter RcType = "edit"

*/

/* QueryMapper interfaces output MediaWiki query parameters in a map format */
type QueryMapper interface {
	Map() map[string]string
}

/*
	RecentChanges is the parameter set for querying for a list of the recent changes

from the wiki.

Defaults to:
Limit = "max" (can also set to -1 for "max")
Type filter to "edit"
TopOnly to false
Continue to be included
No start time
*/
type RecentChanges struct {
	action   string `default:"query"`
	list     string `default:"recentchanges"`
	Limit    int    `prefix:"rc" special:"max"`
	Type     string `default:"edit" prefix:"rc"`
	Continue bool   `default:"" prefix:"rc"`
	Start    string `prefix:"rc"` // Timestamp in ISO8601 format YYYY-MM-DDThh:mm:ssZ
}

/*
	Map creates a query action for getting list of the RecentChanges in a map

format for use with mwclient.Get (or the WikiBot.Core.client)
*/
func (rc RecentChanges) Map() map[string]string {
	fields := reflect.VisibleFields(reflect.TypeOf((struct{ RecentChanges }{})))

	output := make(map[string]string, len(fields))

	for _, field := range fields {

		if field.Name == "RecentChanges" {
			continue
		}

		value, includeKey := getValueOrDefault(rc, field)
		prefix, _ := field.Tag.Lookup("prefix")
		// if can't find the prefix tag, its ok to be blank string
		name := prefix + strings.ToLower(field.Name)

		if includeKey {
			output[name] = value
		}

	}

	return output
}

func getValueOrDefault(q QueryMapper, field reflect.StructField) (value string, ok bool) {

	fieldValue := reflect.ValueOf(q).FieldByName(field.Name)

	if fieldValue.Kind() == reflect.Int {
		value = strconv.FormatInt(fieldValue.Int(), 10)
	}
	if fieldValue.Kind() == reflect.String {
		value = fieldValue.String()
	}

	// -1 is used for special situations where defaulting may not be the right call
	// i.e.: max for rclimit
	if value == "-1" {
		specialValue, ok := field.Tag.Lookup("special")
		return specialValue, ok
	}

	// 0 and "" indicate the field was never set or that the user wants to use
	// the default.
	if value == "0" || value == "" {
		defaultValue, ok := field.Tag.Lookup("default")
		return defaultValue, ok
	}
	return value, true
}
