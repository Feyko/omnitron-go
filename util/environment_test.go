package environment

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvReturnsErrIfNotFound(t *testing.T) {
	assert := assert.New(t)

	_, err := getEnvValue("DoesntExist")

	assert.NotNil(err, "getEnvValue: Environment Variable Key [DoesntExist] should return an error")
}

func TestGetEnvReturnsValue(t *testing.T) {
	assert := assert.New(t)
	expectedResult := "TestValue"
	os.Setenv("TEST_ENV", expectedResult)

	value, _ := getEnvValue("TEST_ENV")

	assert.Equal(expectedResult, value, "getEnvValue: Environment Variable Key [TEST_ENV] should return TestValue")
}

func TestLoadEnvReturnsNilWhenSuccessfullyLoading(t *testing.T) {
	assert := assert.New(t)

	testResponse := loadEnv()

	assert.Nil(testResponse, "loadEnv: cannot find .env file")
}

func TestInitEnvironmentReturnsBotStruct(t *testing.T) {
	assert := assert.New(t)

	testResponse, _ := InitEnvironment()

	assert.Equal("Omnitron-Go", testResponse.username)
}
