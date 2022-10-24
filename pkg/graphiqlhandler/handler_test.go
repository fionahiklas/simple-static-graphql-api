package graphiqlhandler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fionahiklas/simple-static-graphql-api/pkg/graphiqlhandler"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {

	var mockResponseWriter *MockresponseWriter
	var mockLogger *Mocklogger
	var testLogger *logrus.Logger

	resetTest := func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockResponseWriter = NewMockresponseWriter(ctrl)
		mockLogger = NewMocklogger(ctrl)
		testLogger = logrus.New()
	}

	t.Run("create new handler", func(t *testing.T) {
		resetTest(t)
		handler := graphiqlhandler.NewHandler(testLogger)
		require.NotNil(t, handler)
	})

	t.Run("bad request", func(t *testing.T) {
		resetTest(t)
		handler := graphiqlhandler.NewHandler(testLogger)
		testRequest := httptest.NewRequest("POST", "/wibble", nil)
		testResponseRecorder := httptest.NewRecorder()

		handler.ServeHTTP(testResponseRecorder, testRequest)
		require.Equal(t, http.StatusBadRequest, testResponseRecorder.Result().StatusCode)
	})

	t.Run("get request", func(t *testing.T) {
		resetTest(t)
		handler := graphiqlhandler.NewHandler(testLogger)
		testRequest := httptest.NewRequest("GET", "/wibble", nil)
		testResponseRecorder := httptest.NewRecorder()

		handler.ServeHTTP(testResponseRecorder, testRequest)
		require.Equal(t, http.StatusOK, testResponseRecorder.Result().StatusCode)
		require.Equal(t, "text/html", testResponseRecorder.Result().Header.Get("Content-Type"))
	})

	t.Run("response write fails", func(t *testing.T) {
		resetTest(t)
		handler := graphiqlhandler.NewHandler(mockLogger)
		testRequest := httptest.NewRequest("GET", "/wibble", nil)

		testError := errors.New("test error")
		mockResponseWriter.EXPECT().Header().Return(http.Header{})
		mockResponseWriter.EXPECT().WriteHeader(http.StatusOK)
		mockResponseWriter.EXPECT().Write(gomock.Any()).Return(0, testError)
		mockLogger.EXPECT().Errorf(gomock.Any(), testError).Times(1)
		mockLogger.EXPECT().Debugf(gomock.Any())

		handler.ServeHTTP(mockResponseWriter, testRequest)
	})

}
