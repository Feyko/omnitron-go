package mwApi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecentChangesMapperDefaultsActionFieldToQuery(t *testing.T) {
	assert := assert.New(t)

	var testValue RecentChanges

	testResult := testValue.Map()

	assert.Equal(testResult["action"], "query", "Mapping RecentChanges param did not set [action] to [query]")
}

func TestRecentChangesMapperDefaultsListToRecentChanges(t *testing.T) {
	assert := assert.New(t)

	var testValue RecentChanges

	testResult := testValue.Map()

	assert.Equal(testResult["list"], "recentchanges", "Mapping RecentChanges param did not set [list] to [recentchanges]")
}
