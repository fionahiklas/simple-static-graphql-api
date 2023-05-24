//go:generate mockgen -package consumer_test -destination=./mock_consumer_test.go -source $GOFILE
package consumer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/fionahiklas/simple-static-graphql-api/pkg/alarmstorage"
	"github.com/jmespath/go-jmespath"
	"github.com/mitchellh/mapstructure"
)

const (
	AllAlarmSystemQuery   = `{ "query": "{ alarmSystems { name } }" }`
	OneAlarmSystemQuery   = `{"query":"query($identifier: ID!){ alarmSystem(alarmSystemId: $identifier){ identifier name description  } }","variables":{"identifier": "%s"}}`
	AllAlarmNamesJmesPath = "data.alarmSystems[*].name"
	OneAlarmJmesPath      = "data.alarmSystem"
)

var allAlarmNamesJmesPathExpression = jmespath.MustCompile(AllAlarmNamesJmesPath)
var oneAlarmJmesPathExpression = jmespath.MustCompile(OneAlarmJmesPath)

type logger interface {
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type httpClient interface {
	Do(request *http.Request) (*http.Response, error)
}

type consumer struct {
	log        logger
	client     httpClient
	graphqlUrl string
}

func NewConsumer(log logger, client httpClient, graphqlUrl string) *consumer {
	return &consumer{
		log:        log,
		client:     client,
		graphqlUrl: graphqlUrl,
	}
}

func (con *consumer) GetOneAlarm(identifier string) (*alarmstorage.Alarm, error) {
	specificAlarmSystemQuery := fmt.Sprintf(OneAlarmSystemQuery, identifier)
	oneAlarmJsonObject, err := con.makeGraphQLRequest(specificAlarmSystemQuery, oneAlarmJmesPathExpression)
	if err != nil {
		return nil, err
	}

	oneAlarm, err := convertToAlarmObject(oneAlarmJsonObject)
	if err != nil {
		con.log.Errorf("Couldn't convert to an array of strings")
		return nil, errors.New("not a string array result")
	}
	return oneAlarm, nil
}

func convertToAlarmObject(jsonNodeObject interface{}) (*alarmstorage.Alarm, error) {
	alarmResult := &alarmstorage.Alarm{}
	err := mapstructure.Decode(jsonNodeObject, alarmResult)
	return alarmResult, err
}

func (con *consumer) GetAllAlarmNames() ([]string, error) {
	alarmJsonList, err := con.makeGraphQLRequest(AllAlarmSystemQuery, allAlarmNamesJmesPathExpression)
	if err != nil {
		return nil, err
	}

	alarmList, err := convertToArrayOfStrings(alarmJsonList)
	if err != nil {
		con.log.Errorf("Couldn't convert to an array of strings")
		return nil, errors.New("not a string array result")
	}

	return alarmList, nil
}

func convertToArrayOfStrings(jmesPathResult interface{}) ([]string, error) {
	interfaceArray, ok := jmesPathResult.([]interface{})
	if !ok {
		return nil, errors.New("not an array")
	}
	resultArray := make([]string, 0, len(interfaceArray))
	for _, element := range interfaceArray {
		elementAsString, ok := element.(string)
		if !ok {
			return nil, errors.New("elements aren't strings")
		}
		resultArray = append(resultArray, elementAsString)
	}
	return resultArray, nil
}

func (con *consumer) makeGraphQLRequest(queryString string, jmesPathExpression *jmespath.JMESPath) (interface{}, error) {
	request := con.buildHttpRequestFromGraphQLQueryString(queryString)
	response, err := con.client.Do(request)

	if err != nil {
		con.log.Errorf("HTTP Request failed: %s", err)
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		con.log.Errorf("Incorrect status code: %d", response.StatusCode)
		return nil, err
	}

	var bodyJson interface{}
	defer response.Body.Close()

	jsonDecoder := json.NewDecoder(response.Body)
	err = jsonDecoder.Decode(&bodyJson)

	if err != nil {
		con.log.Errorf("Couldn't decode JSON: %s", err)
		return nil, err
	}

	payloadData, err := jmesPathExpression.Search(bodyJson)

	if err != nil {
		con.log.Errorf("Couldn't extract data from response: %s", err)
		return nil, err
	}
	return payloadData, nil
}

// TODO: Cache the URL and request to save time
func (con *consumer) buildHttpRequestFromGraphQLQueryString(queryString string) *http.Request {
	gqlUrl, err := url.Parse(con.graphqlUrl)
	if err != nil {
		con.log.Errorf("Error parsing URL: %s", err)
		return nil
	}
	return &http.Request{
		Method: http.MethodPost,
		URL:    gqlUrl,
		Body:   io.NopCloser(strings.NewReader(queryString)),
		Header: http.Header{"Content-Type": {"application/json"}},
	}
}
