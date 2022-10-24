//go:generate mockgen -package versionhandler_test -destination=./mock_handler_test.go -source $GOFILE
package versionhandler

import (
	"encoding/json"
	"net/http"
)

//lint:ignore U1000 Just here so that it can be mocked
type responseWriter interface {
	http.ResponseWriter
}

type logger interface {
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type handler struct {
	log        logger
	version    string
	commitHash string
}

func NewHandler(log logger, version string, commitHash string) *handler {
	return &handler{
		log:        log,
		version:    version,
		commitHash: commitHash,
	}
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	h.log.Debugf("Version handler called")
	if request.Method != http.MethodGet {
		h.log.Errorf("Tried to use method '%s'", request.Method)
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	responseJson := h.marshalVersion()
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	_, err := responseWriter.Write(responseJson)

	if err != nil {
		h.log.Errorf("Couldn't write version response, error: %w", err)
	}
}

// TODO: Cache the bytes so we don't keep marshalling the same thing
func (h *handler) marshalVersion() []byte {
	// Don't forget, fields must be exported to be
	// marshalled (I forgot, hence this comment :) )
	type versionJson struct {
		Version string `json:"version"`
		Commit  string `json:"commit"`
	}

	toMarshal := versionJson{
		Version: h.version,
		Commit:  h.commitHash,
	}

	// I can't find a way that this can actually fail given the input
	// is strings.  The tests should handle catching anything really bizarre
	// or the app will panic fairly quickly when running is the version
	// endpoint is used as a health check in Docker/k8s
	marshalledVersion, _ := json.Marshal(&toMarshal)
	return marshalledVersion
}
