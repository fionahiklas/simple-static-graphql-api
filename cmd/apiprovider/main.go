package main

import (
	"context"
	"github.com/fionahiklas/simple-static-graphql-api/internal/graphapi"
	"github.com/fionahiklas/simple-static-graphql-api/internal/version"
	"github.com/fionahiklas/simple-static-graphql-api/pkg/graphiqlhandler"
	"github.com/fionahiklas/simple-static-graphql-api/pkg/teststorage"
	"github.com/fionahiklas/simple-static-graphql-api/pkg/versionhandler"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log := logrus.New()

	log.SetLevel(logrus.DebugLevel)

	dummyData := teststorage.NewStorage(log)

	graphApi := graphapi.NewGraphQLAPI(log, dummyData)
	graphApiHandler := graphApi.GetHandler()
	versionHandler := versionhandler.NewHandler(log, version.CodeVersion(), version.CommitHash())
	graphiqlHandler := graphiqlhandler.NewHandler(log)

	mux := http.NewServeMux()
	mux.Handle("/graphql", graphApiHandler)
	mux.Handle("/version", versionHandler)
	mux.Handle("/iql", graphiqlHandler)

	log.Infof("Starting provider version '%s', commit: '%s' on localhost:7777",
		version.CodeVersion(), version.CommitHash())

	if err := http.ListenAndServe("localhost:7777", mux); err != nil {
		log.Error(context.Background(), "Error from HTTP server: %s", err.Error())
	}
}
