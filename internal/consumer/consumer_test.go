//go:build !pact

package consumer_test

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/fionahiklas/simple-static-graphql-api/internal/consumer"
	"github.com/fionahiklas/simple-static-graphql-api/pkg/alarmstorage"
	"github.com/golang/mock/gomock"
	"github.com/jmespath/go-jmespath"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestJmesPathOutputForObjects(t *testing.T) {

	t.Run("for one alarm", func(t *testing.T) {
		// Test data refers to the character Granny Weatherwax in the Terry Pratchett Discworld books
		const testJsonData = `{"data":{"alarmSystem":{"name":"Weatherwax", "description": "Witch", "identifier": "Granny"}}}`
		jmesPathExpression := jmespath.MustCompile(consumer.OneAlarmJmesPath)

		var jsonData interface{}
		err := json.Unmarshal([]byte(testJsonData), &jsonData)
		require.NoError(t, err)

		alarmJsonResult, err := jmesPathExpression.Search(jsonData)
		require.NoError(t, err)
		require.NotNil(t, alarmJsonResult)

		result := alarmstorage.Alarm{}
		err = mapstructure.Decode(alarmJsonResult, &result)
		require.NoError(t, err)
		require.Equal(t, "Granny", result.Identifier)
		require.Equal(t, "Weatherwax", result.Name)
		require.Equal(t, "Witch", result.Description)
	})
}

func TestConsumer_GetAllAlarmNames(t *testing.T) {
	const (
		testCallUrl       = "https://ankh.morpork.discworld"
		testHappyResponse = `{"data":{"alarmSystems":[{"name":"Alarm One"},{"name":"Alarm Two"},{"name":"Alarm Three"}]}}`
	)

	// For now don't bother with using a mock for the logger
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	var mockHttpClient *MockhttpClient

	resetMocks := func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockHttpClient = NewMockhttpClient(ctrl)
	}

	buildResponse := func(status int, body string) *http.Response {
		return &http.Response{
			StatusCode: status,
			Body:       io.NopCloser(strings.NewReader(body)),
		}
	}

	t.Run("happy path", func(t *testing.T) {
		resetMocks(t)
		consumerToTest := consumer.NewConsumer(log, mockHttpClient, testCallUrl)

		mockHttpClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(request *http.Request) (*http.Response, error) {
			require.Equal(t, testCallUrl, request.URL.String())
			return buildResponse(http.StatusOK, testHappyResponse), nil
		})

		alarmResult, err := consumerToTest.GetAllAlarmNames()
		require.NoError(t, err)
		require.Equal(t, 3, len(alarmResult))
	})
}

func TestConsumer_GetOneAlarm(t *testing.T) {
	const (
		testCallUrl       = "https://ankh.morpork.discworld"
		testHappyResponse = `{"data":{"alarmSystem":{"name":"Aching", "description": "Witch", "identifier": "Tiffany"}}}`
	)

	// For now don't bother with using a mock for the logger
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	var mockHttpClient *MockhttpClient

	resetMocks := func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockHttpClient = NewMockhttpClient(ctrl)
	}

	buildResponse := func(status int, body string) *http.Response {
		return &http.Response{
			StatusCode: status,
			Body:       io.NopCloser(strings.NewReader(body)),
		}
	}

	t.Run("happy path", func(t *testing.T) {
		resetMocks(t)
		consumerToTest := consumer.NewConsumer(log, mockHttpClient, testCallUrl)

		mockHttpClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(request *http.Request) (*http.Response, error) {
			require.Equal(t, testCallUrl, request.URL.String())
			return buildResponse(http.StatusOK, testHappyResponse), nil
		})

		alarmResult, err := consumerToTest.GetOneAlarm("Tiffany")
		require.NoError(t, err)
		require.NotNil(t, alarmResult)
		require.Equal(t, "Tiffany", alarmResult.Identifier)
		require.Equal(t, "Aching", alarmResult.Name)
		require.Equal(t, "Witch", alarmResult.Description)
	})
}
