package consumer

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	allAlarmSystemQuery = `{ "query": "{ alarmSystems { name } }" }`
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

	_, err := con.client.Do(request)
	if err != nil {
		con.log.Errorf("HTTP Request failed: %s", err)
		return nil, err
	}
	return []string{}, nil
}

// TODO: Cache the URL and request to save time
func (con *consumer) buildAllAlarmSystemRequest() *http.Request {
	url, err := url.Parse(con.graphqlUrl)
	if err != nil {
		con.log.Errorf("Error parsing URL: %s", err)
		return nil
	}
	return &http.Request{
		Method: "POST",
		URL:    url,
		Body:   io.NopCloser(strings.NewReader(allAlarmSystemQuery)),
		Header: http.Header{"Content-Type": {"application/json"}},
	}
}
