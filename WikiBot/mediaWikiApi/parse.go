package mwApi

import (
	"golang.org/x/exp/slices"
)

type Parse struct {
	action string `default:"parse"`
	PageId string
	Page   string
	Title  string
}

func (pa Parse) Map() map[string]string {
	fields, output := prepMap(struct{ Parse }{})

	// these values are mutually exclusive to one another. If more than one are set we dont want
	// them all sent to the api parameters
	contentIdentifiers := []string{"PageId", "Page", "Title"}

	alreadyHaveContentIdentifier := false

	for _, field := range fields {

		if field.Name == "Parse" {
			continue
		}

		// since reflection returns the slice of fields in the order in which they are defined on the struct
		// we can use that to define priority as well
		if slices.Contains(contentIdentifiers, field.Name) {
			alreadyHaveContentIdentifier = !isFieldBlank(pa, field)
		}

		if alreadyHaveContentIdentifier {
			continue
		}
		getKeyAndValue(pa, field, output)

	}

	//ok := verifyDependantKeysPresent(output, parseKeyDependencies)

	return output

}

type ParseProps struct {
}
