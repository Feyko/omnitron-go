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

func TestRecentChangesMapperCorrectsKeysNeedingRcPrefix(t *testing.T) {
	assert := assert.New(t)

	var testValue RecentChanges
	testValue.Limit = 2

	testResult := testValue.Map()

	assert.Contains(testResult, "rclimit", "Mapping Recent Changes did not lower and prefix [Limit] to [rclimit]")
}

func TestRecentChangesMapperIncludesLimitValueIfNotBlank(t *testing.T) {
	assert := assert.New(t)

	var testValue RecentChanges
	testValue.Limit = 2

	testResult := testValue.Map()

	assert.Equal(testResult["rclimit"], "2", "Mapping Recent Changes param did not set [rclimit:2] to [2]")
}

func TestRecentChangesMapperSetsRcLimitToMaxIfLimitIsNeg1(t *testing.T) {
	assert := assert.New(t)

	var testValue RecentChanges
	testValue.Limit = -1

	testResult := testValue.Map()

	assert.Equal(testResult["rclimit"], "max", "Mapping Recent Changes param did not set [rclimit:-1] to [max]")
}

func TestRecentChangesMapperDoesNotOutputStartIfNotSet(t *testing.T) {
	assert := assert.New(t)

	var testValue RecentChanges
	testValue.Limit = 2

	testResult := testValue.Map()

	assert.NotContains(testResult, "rcstart", "Mapping Recent Changes still included [rcstart] in output keys")
}
