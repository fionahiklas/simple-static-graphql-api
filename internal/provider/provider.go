//go:generate mockgen -package provider_test -destination=./mock_provider_test.go -source $GOFILE
package provider

import "github.com/fionahiklas/simple-static-graphql-api/pkg/alarmstorage"

type ReadAndWrite interface {
	alarmstorage.ReadAndWrite
}
