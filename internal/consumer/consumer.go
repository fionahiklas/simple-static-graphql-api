//go:generate mockgen -package consumer_test -destination=./mock_consumer_test.go -source $GOFILE
package consumer

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/jmespath/go-jmespath"
)

const (
	AllAlarmSystemQuery  = `{ "query": "{ alarmSystems { name } }" }`
	JmesPathForAlarmList = "data.alarmSystems[*].name"
)

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

func (con *consumer) GetAllAlarmNames() ([]string, error) {
	request := con.buildAllAlarmSystemRequest()

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

	jmesPathExpression, err := jmespath.Compile(JmesPathForAlarmList)
	if err != nil {
		con.log.Errorf("JMESPath '%s' doesn't compile: ", jmesPathExpression, err)
		return nil, err
	}

	alarmJsonList, err := jmesPathExpression.Search(bodyJson)
	if err != nil {
		con.log.Errorf("Couldn't find alarm list: %s", err)
		return nil, err
	}

	alarmList, err := convertToArrayOfStrings(alarmJsonList)
	if err != nil {
		con.log.Errorf("Couldn't convert to an array of strings")
		return nil, errors.New("not a string array result")
	}

	return alarmList, nil
}

// TODO: Cache the URL and request to save time
func (con *consumer) buildAllAlarmSystemRequest() *http.Request {
	url, err := url.Parse(con.graphqlUrl)
	if err != nil {
		con.log.Errorf("Error parsing URL: %s", err)
		return nil
	}
	return &http.Request{
		Method: http.MethodPost,
		URL:    url,
		Body:   io.NopCloser(strings.NewReader(AllAlarmSystemQuery)),
		Header: http.Header{"Content-Type": {"application/json"}},
	}
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
