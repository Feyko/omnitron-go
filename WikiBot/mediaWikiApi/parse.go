package mwApi

import (
	"golang.org/x/exp/slices"
)

type Parse struct {
	action       string `default:"parse"`
	PageId       string
	Title        string
	Page         string
	ContentModel string `default:"text"`
	Prop         string `default:"wikitext|categories"` // A pipe separated list of properties. see https://www.mediawiki.org/wiki/API:Parsing_wikitext#parse
}

func (pa Parse) Map() map[string]string {
	fields, output := prepMap(struct{ Parse }{})

	// these values are mutually exclusive to one another. If more than one are set we don't want
	// them all sent to the api parameters.
	contentIdentifiers := []string{"PageId", "Page", "Title"}

	alreadyHaveContentIdentifier := false

	for _, field := range fields {

		if field.Name == "Parse" {
			continue
		}

		// since reflection returns the slice of fields in the order in which they are defined on the struct
		// we can use that to define priority as well
		if slices.Contains(contentIdentifiers, field.Name) {

			if alreadyHaveContentIdentifier {
				continue
			}

			alreadyHaveContentIdentifier = !isFieldBlank(pa, field)
		}

		getKeyAndValue(pa, field, output)

	}

	//ok := verifyDependantKeysPresent(output, parseKeyDependencies)

	return output

}
