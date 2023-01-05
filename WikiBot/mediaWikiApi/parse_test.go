package mwApi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMapperPrioritizesPageIdOverOtherIdentifiers(t *testing.T) {
	assert := assert.New(t)

	var testValue Parse
	testValue.Page = "TestPage"
	testValue.PageId = "12345"
	testValue.Title = "TestTitle"

	testResult := testValue.Map()

	assert.Contains(testResult, "title", "Parse Mapper should have contained [title] key")
	assert.NotContains(testResult, "page", "Parse Mapper was given [Title] value and still contains mutually exclusive [page] key")
	assert.NotContains(testResult, "pageid", "Parse Mapper was given [Title] value and still contains mutually exclusive [pageid] key")
}
