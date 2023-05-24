# simple_static_graphql_api

## Overview

Simple test application for GraphQL API in Go

![Main build](../../actions/workflows/build-and-test.yaml/badge.svg?branch=main)


## Quickstart

Follow the instructions in [docs/pact_testing.md](docs/pact_testing.md) to setup the 
Pact tools.


## Development Notes

Details of how this project was created.  You don't need to follow these 
instructions to get things up and running

### Setup 

* Created initial module file

```
go mod init github.com/fionahiklas/simple-static-graphql-api
```

* Adding dependency 

``` 
go get github.com/graph-gophers/graphql-go
```

* Created basic Makefile to be able to run tests and linting 


### Resolvers

* For string fields which can be nullable (don't have `!`) the type is `*string` rather than `string`
* To set pointer to empty string as a return value needs some work as per 
[this post](https://stackoverflow.com/questions/42594789/initialize-string-pointer-in-struct)
* It seems that `[SomeGraphQLType!]!` requires a Go slice so `*[]*SomeGoStruct` doesn't seem to be correct
* Unless you allow struct resolvers (this is mentioned in the docs, __TODO:__ add link) each resolver needs 
to implement accessors for every field in the GraphQL type 



## References

### GraphQL

* [GraphQL implementations for Go](https://graphql.org/code/#go)
* [Go graphQL framework from graph-gophers](https://github.com/graph-gophers/graphql-go)
* [GraphQL by example](https://github.com/tonyghita/graphql-go-example)
* [GraphIQL](https://github.com/graphql/graphiql/tree/main/packages/graphiql)
* [Apollo Server Resolvers](https://www.apollographql.com/docs/apollo-server/data/resolvers/)
* [GraphQL Schema](https://graphql.org/learn/schema/)
* [Serving GraphQL over HTTP](https://graphql.org/learn/serving-over-http/)
* [GraphQL inspector](https://www.the-guild.dev/graphql/inspector/docs/essentials/introspect)
* [GraphQL Playground](https://github.com/graphql/graphql-playground)
* [GraphQL schema basics](https://www.apollographql.com/docs/apollo-server/schema/schema/)
* [Introspection](https://graphql.org/learn/introspection/)
* [GraphQL spec](https://github.com/graphql/graphql-spec/blob/main/README.md)

### Other Go

* [JMESPath Go implementation](https://github.com/jmespath/go-jmespath)
* [JMESPath Tutorial](https://jmespath.org/tutorial.html)
* [Converting from arrays of interface](https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces)
* [Trim whitespace](https://yourbasic.org/golang/trim-whitespace-from-string/)
* [Listen on random/free port](https://stackoverflow.com/questions/43424787/how-to-use-next-available-port-in-http-listenandserve)
* [Read whole file](https://stackoverflow.com/questions/13514184/how-can-i-read-a-whole-file-into-a-string-variable)
* [Using go ldflags to set version values](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications)
* [Decode maps to structs](https://github.com/mitchellh/mapstructure)

### Github Actions

* [Setup Go action](https://github.com/actions/setup-go)
* [Go tutorial post](https://medium.com/swlh/setting-up-github-actions-for-go-project-ea84f4ed3a40)
* [Versioning for Go setup](https://github.com/npm/node-semver)
* [Actions build badge](https://docs.github.com/en/actions/monitoring-and-troubleshooting-workflows/adding-a-workflow-status-badge)


### Pact testing 

* [Pact testing homepage](https://pact.io)
* [Pact documentation](https://docs.pact.io)
* [Pact testing tutorial](https://medium.com/nerd-for-tech/the-ultimate-guide-for-contract-testing-with-pact-and-go-177b4af13700)
* [Pact go workshop](https://github.com/pact-foundation/pact-workshop-go)
* [Pact go implementation](https://github.com/pact-foundation/pact-go)
* [Pact testing with GraphQL](https://pactflow.io/blog/contract-testing-a-graphql-api/)

### Ruby

* [Bundler](https://bundler.io/v2.3/man/bundle.1.html)
* [Adding a Gem path](https://stackoverflow.com/questions/44293904/add-a-gem-path)

### UNIX

* [Getting fields from input](https://stackoverflow.com/questions/16136943/how-to-get-the-second-column-from-command-output)