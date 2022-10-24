# simple_static_graphql_api

## Overview

Simple test application for GraphQL API in Go

![Main build](../../actions/workflows/build-and-test.yaml/badge.svg?branch=main)

## Notes

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


### Github Actions

* [Setup Go action](https://github.com/actions/setup-go)
* [Go tutorial post](https://medium.com/swlh/setting-up-github-actions-for-go-project-ea84f4ed3a40)
* [Versioning for Go setup](https://github.com/npm/node-semver)
* [Actions build badge](https://docs.github.com/en/actions/monitoring-and-troubleshooting-workflows/adding-a-workflow-status-badge)


### Pact testing 


