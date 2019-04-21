package dgraph

import (
	"context"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/dolan-in/dgman"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Connection isitialize a DGraph client connection
func Connection(ctx context.Context, addr string) (*dgo.Dgraph, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Return dgraph client
	return dgo.NewDgraphClient(
		api.NewDgraphClient(conn),
	), nil
}

func Setup(ctx context.Context, addr string) error {

	c, err := Connection(ctx, addr)
	if err != nil {
		return errors.Wrap(err, "unable to initialize dgraph connection")
	}

	// Initialize schema from models
	schema, err := dgman.MutateSchema(c,
		&dgoUser{},
	)
	if err != nil {
		return errors.Wrap(err, "unable to initialize dgraph schema")
	}

	return nil
}
