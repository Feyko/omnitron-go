package mwApi

import (
	"reflect"
	"strconv"
	"strings"
)

// A QueryMapper can have its properties output as a map, with defaults and prefix for keys.
type QueryMapper interface {
	Map() map[string]string
}

/*
RecentChanges is the parameter set for querying for a list of the recent changes from the wiki.

Defaults:

	Type filter to "edit"
	TopOnly to false
	Continue to be included
	Limit not set
	No start time
*/
type RecentChanges struct {
	action   string `default:"query"`
	list     string `default:"recentchanges"`
	Limit    int    `prefix:"rc" special:"max"`
	Type     string `default:"edit" prefix:"rc"`
	Continue string `prefix:"rc"`
	Start    string `prefix:"rc"` // Timestamp in ISO8601 format YYYY-MM-DDThh:mm:ssZ
}

/*
RecentChanges.Map outputs the parameters as a map for mwclient.Client.Get actions.
*/
func (rc RecentChanges) Map() map[string]string {
	fields, output := prepMap(struct{ RecentChanges }{})

	for _, field := range fields {

		if field.Name == "RecentChanges" {
			continue
		}
		// if can't find the prefix tag, its ok to be blank string
		getKeyAndValue(rc, field, output)

	}

	return output
}

type Parse struct {
	action string `default:"parse"`
	Page   string
}

func (pa Parse) Map() map[string]string {
	fields, output := prepMap(struct{ Parse }{})

	for _, field := range fields {

		if field.Name == "Parse" {
			continue
		}
		// if can't find the prefix tag, its ok to be blank string
		getKeyAndValue(pa, field, output)

	}

	return output
}

/*PrepMap takes a Type and generates the fields and a preparatory map for them.
 */
func prepMap(structType any) ([]reflect.StructField, map[string]string) {
	fields := reflect.VisibleFields(reflect.TypeOf(structType))

	output := make(map[string]string, len(fields))
	return fields, output
}

/*
GetKeyAndValue is a helper for converting a QueryMapper type fields to a map.
*/
func getKeyAndValue(q QueryMapper, field reflect.StructField, output map[string]string) {
	value, includeKey := getValueOrDefault(q, field)
	prefix, _ := field.Tag.Lookup("prefix")

	name := prefix + strings.ToLower(field.Name)

	if includeKey {
		output[name] = value
	}
}

/*
GetValueOrDefault is a helper for determining the value of a QueryMapper field or its default.
*/
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
