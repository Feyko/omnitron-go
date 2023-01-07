package mwApi

/*
	Parse is the parameter set for parsing a page of the wiki.

It implements the QueryMapper interface.

It defaults to:

	pase action
	content model of text
	retrieving the properties for wikitext and categories.
	It does NOT have a PageId/Title/Page set and so will need this set before being called.
*/
type Parse struct {
	action       string `default:"parse"`
	PageId       string
	Title        string
	Page         string
	ContentModel string `default:"text"`
	Prop         string `default:"wikitext|categories|templates"` // A pipe separated list of properties. see https://www.mediawiki.org/wiki/API:Parsing_wikitext#parse
}

/*
	Parse.Map() outputs a parameter map for the Parse Query.

It will only allow one of the following: PageId, Title, Page.
It will take the first it encounters in that order, discarding the others.
*/
func (pa *Parse) Transform(in map[string]string) {
	// I would probably just add a "unique" tag to the map utils
}
