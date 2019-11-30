package protocol

import (
	"context"
	"github.com/imulab/go-scim/src/core/prop"
)

type PersistenceProvider interface {
	// Insert the given resource into the database, or return any error.
	Insert(ctx context.Context, resource *prop.Resource) error
}