# simple_static_graphql_api

## Overview


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

### Resolvers

* For string fields which can be nullable (don't have `!`) the type is `*string` rather than `string`
* It seems that `[SomeGraphQLType!]!` requires a Go slice so `*[]*SomeGoStruct` doesn't seem to be correct
* Unless you allow struct resolvers (this is mentioned in the docs, __TODO:__ add link) each resolver needs 
to implement accessors for every field in the GraphQL type 



## References

### GraphQL

* 
