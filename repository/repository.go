package repository

import (
	"context"

	"github.com/agilefoxHQ/go-function-template/config"
)

type Repositories struct {
	//Auth       svc.SomeExternalAPI
	//Users         svc.SomeStore
}

// NewRepository will load the underlying data stores. It abstracts away all the logic that we
// need to handle. This also lets us swap out the underlying stores if they implement the
// methods defined by the stores here.
func NewRepository(ctx context.Context, c config.Configuration) (*Repositories, error) {
	return &Repositories{}, nil
}

func CloseConnections() error {
	// some code to close connections
	return nil
}
