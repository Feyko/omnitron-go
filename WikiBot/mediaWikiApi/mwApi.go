/*
	The mwApi Package contains the mediaWiki API connection information.

In addition it contains the basic core functionality that either the Sentinels or SpiritIsland packages will implement
*/
package mwApi

import (
	"reflect"
	"strconv"
	"strings"
)

/* A QueryMapper can have its properties output as a map, with defaults and prefix for keys.

Use tags when implementing a QueryMapper:

	default:"[default value]" - the default value to include if the property is "", 0, or nil.
	special:"[special value]" - a specific value to use for the property if it is set to -1.
		useful for integer properties that also can accept strings like "max"
	prefix: "[prefix]" a prefix for the key, such as 'rc' on the Limit property, leading to 'rclimit' as the output key.
	ignore:"any value" if the ignore tag is present, this property is ignored and not included in the Map output.
*/

type ParameterSet interface {
	Transform(map[string]string)
}

/* Map takes a parameter set and makes it into a map suited for use as parameters for the API
 */
func Map(parameterSet ParameterSet) map[string]string {
	// Using ValueOf and Elem might not be necessary, I don't exactly remember the semantics. Try it!
	fields := reflect.VisibleFields(reflect.ValueOf(parameterSet).Elem().Type())
	output := make(map[string]string, len(fields))

	for _, field := range fields {
		GetKeyAndValue(parameterSet, field, output)
	}

	return output
}

/*
GetKeyAndValue is a helper for converting a QueryMapper type fields to a map.
*/
func GetKeyAndValue(q any, field reflect.StructField, output map[string]string) {
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
func getValueOrDefault(q any, field reflect.StructField) (value string, ok bool) {

	value = extractString(q, field)

	ignoreValue, ok := field.Tag.Lookup("ignore")
	if ignoreValue != "" {
		return "", false
	}

	// -1 is used for special situations where defaulting may not be the right call
	// i.e.: max for rclimit
	if value == "-1" {
		specialValue, ok := field.Tag.Lookup("special")
		return specialValue, ok
	}

	if isFieldBlank(q, field) {
		defaultValue, ok := field.Tag.Lookup("default")
		return defaultValue, ok
	}
	return value, true
}

// reflect.Value.String() already does this
/*
	extractString gets the value of the field back as a string, for use in the QueryMapper.Map() functions
*/
func extractString(q any, field reflect.StructField) (value string) {
	fieldValue := reflect.ValueOf(q).FieldByName(field.Name)

	if fieldValue.Kind() == reflect.Int {
		value = strconv.FormatInt(fieldValue.Int(), 10)
	}
	if fieldValue.Kind() == reflect.String {
		value = fieldValue.String()
	}
	return value
}

// Probably use pointers to determine if something was set. Checking the zero-value isn't safe as they are probably a valid value.
// Also, reflect has a function for that
/* isFieldBlank is a helper function for determining if fields have 0 or "" indicating they were not set
 */
func isFieldBlank(q any, field reflect.StructField) bool {
	value := extractString(q, field)
	// 0 and "" indicate the field was never set or that the user wants to use
	// the default.
	return value == "0" || value == ""
}
