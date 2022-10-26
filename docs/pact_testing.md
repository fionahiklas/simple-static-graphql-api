# Pact Testing

## Overview

Testing out the Pact testing framework

## Get Started

* Set up the Ruby code for Pact testing locally

```
export RUBY_VERSION_FOR_PATH=`ruby --version | cut -d' ' -f2 | cut -d. -f1-2`.0

bundle config set path $PWD/.gem
bundle install

export PATH=$PWD/.gem/ruby/$RUBY_VERSION_FOR_PATH/bin:$PATH
export GEM_PATH=`gem env gempath`:$PWD/.gem/ruby/$RUBY_VERSION_FOR_PATH
```


## Notes

### Initial Setup

__NOTE:__ These were steps carried out when creating the project.  You don't need to repeat 
these locally.  Instead follow the [Getting started](#GettingStarted) section above


#### Installing Ruby

The version of Ruby on Mac is quite old - 2.6.8 when latest is 3.x install using homebrew

``` 
brew install ruby
```

#### Installing Ruby Tools

Using bundler this codebase was setup using the following the following commands 

``` 
bundle config set path $PWD/.gem
```

Adding dependencies

``` 
bundle add pact
bundle add pact-provider-verifier
bundle add pact_broker-client
```

I think the last one is only needed for connecting to a Pact broker like [pactflow](pactflow.io)

Adding tools into the path

``` 
export PATH=$PWD/.gem/ruby/3.1.0/bin:$PATH
export GEM_PATH=`gem env gempath`:$PWD/.gem/ruby/3.1.0
```

__NOTE:__ Only run the `GEM_PATH` setup once otherwise you'll keep adding to that environment variable: 
the `gem env gempath` picks up the path from `GEM_PATH`


#### Adding Go Implementation

Add dependency

```
go get github.com/pact-foundation/pact-go
```


## Troubleshooting 

### CLI tools are out of date

Running the Pact consumer unit test via the IDE see these errors

``` 
/usr/local/go/bin/go tool test2json -t /private/var/folders/zf/4fx6p54x2c1dqs27xl1b2vnm0000gp/T/GoLand/___TestConsumer_GetAllAlarmNames_pact_consumer_test_in_github_com_fionahiklas_simple_static_graphql_api_internal_consumer.test -test.v -test.paniconexit0 -test.run ^\QTestConsumer_GetAllAlarmNames\E$/^\Qpact_consumer_test\E$
=== RUN   TestConsumer_GetAllAlarmNames
=== RUN   TestConsumer_GetAllAlarmNames/pact_consumer_test
2022/10/24 22:12:41 [INFO] checking pact-provider-verifier within range >= 1.36.1, < 2.0.0
2022/10/24 22:12:41 [ERROR] CLI tools are out of date, please upgrade before continuing
```
 
or this one on a subsequent run

``` 
/usr/local/go/bin/go tool test2json -t /private/var/folders/zf/4fx6p54x2c1dqs27xl1b2vnm0000gp/T/GoLand/___TestConsumer_GetAllAlarmNames_pact_consumer_test_in_github_com_fionahiklas_simple_static_graphql_api_internal_consumer.test -test.v -test.paniconexit0 -test.run ^\QTestConsumer_GetAllAlarmNames\E$/^\Qpact_consumer_test\E$
=== RUN   TestConsumer_GetAllAlarmNames
=== RUN   TestConsumer_GetAllAlarmNames/pact_consumer_test
2022/10/24 22:14:15 [INFO] checking pact-mock-service within range >= 3.5.0, < 4.0.0
2022/10/24 22:14:15 [ERROR] CLI tools are out of date, please upgrade before continuing
```

In the above case the `test2json` is just the way that the IDE is running a single test 

Seems I missed a step and need to download the CLI tools as per the [instructions](https://github.com/pact-foundation/pact-go#installation)

Looks like these tools are [ruby applications](https://github.com/pact-foundation/pact-ruby-standalone/releases)

### Error Adding Ruby Gems

After I did `bundle init` the `Gemfile` was created read only and I got these errors trying to add gems

```
Errno::EACCES: Permission denied @ rb_sysopen - /Users/user/wd/simple-static-graphql-api/Gemfile
```

Used `chmod +w Gemfile` and all working again.  Very weird