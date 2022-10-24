package main

import (
	"context"
	"github.com/fionahiklas/simple-static-graphql-api/internal/graphapi"
	"github.com/fionahiklas/simple-static-graphql-api/pkg/graphiqlhandler"
	"github.com/fionahiklas/simple-static-graphql-api/pkg/teststorage"
	"github.com/fionahiklas/simple-static-graphql-api/pkg/versionhandler"
	"github.com/sirupsen/logrus"
	"net/http"
)

// These values are set by "-X" options to the ldflags for Go, see the Makefile
var (
	commitHash  = "null"
	codeVersion = "null"
)

func main() {
	log := logrus.New()

	log.SetLevel(logrus.DebugLevel)

	dummyData := teststorage.NewStorage(log)

	graphApi := graphapi.NewGraphQLAPI(log, dummyData)
	graphApiHandler := graphApi.GetHandler()
	versionHandler := versionhandler.NewHandler(log, codeVersion, commitHash)
	graphiqlHandler := graphiqlhandler.NewHandler(log)

	mux := http.NewServeMux()
	mux.Handle("/graphql", graphApiHandler)
	mux.Handle("/version", versionHandler)
	mux.Handle("/iql", graphiqlHandler)

	log.Infof("Starting provider on localhost:7777")
	if err := http.ListenAndServe("localhost:7777", mux); err != nil {
		log.Error(context.Background(), "Error from HTTP server: %s", err.Error())
	}
}
