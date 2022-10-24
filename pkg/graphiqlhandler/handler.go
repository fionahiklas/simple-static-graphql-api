//go:generate mockgen -package graphiqlhandler_test -destination=./mock_handler_test.go -source $GOFILE
package graphiqlhandler

import (
	_ "embed"
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
	log logger
}

//go:embed graphiql.html
var graphiqlHtml []byte

func NewHandler(log logger) *handler {
	return &handler{
		log: log,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Debugf("GraphiQL handler called")
	if r.Method != http.MethodGet {
		h.log.Errorf("Tried to use method '%s'", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(graphiqlHtml)

	if err != nil {
		h.log.Errorf("Couldn't write graphiql response: %w", err)
	}
}
