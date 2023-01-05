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

	assert.NotContains(testResult, "page", "Parse Mapper was given [PageId] value and still contains mutually exclusive [page] key")
	assert.NotContains(testResult, "title", "Parse Mapper was given [PageId] value and still contains mutually exclusive [title] key")
}
