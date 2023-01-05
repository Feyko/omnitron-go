package mwApi

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
	NextPage string
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

		if field.Name == "NextPage" {
			output["continue"] = output["nextpage"]
			delete(output, "nextpage")
		}

	}

	return output
}
