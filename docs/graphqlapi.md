# GraphQL API

## Overview

Details and examples for querying the GraphQL API provide by the API


## Examples

### Retrieve all available alarm systems

``` 
curl \
-X POST \
-H "Content-Type: application/json" \
--data '{ "query": "{ alarmSystems { name id } }" }' \
--silent http://localhost:7777/graphql | json_pp 
```
