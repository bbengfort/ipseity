package ipseity

import (
	"fmt"

	"github.com/bbengfort/ipseity/pb"
	"google.golang.org/grpc"
)

// NewClient creates an ipseity client and connects it to the ipseity server
// at the given address, allowing the user to make identity requests.
func NewClient(addr string) (pb.IdentityClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not dial %s: %s", addr, err)
	}

	return pb.NewIdentityClient(conn), nil
}

// NewStreamClient creates an ipseity streaming service client and connects
// it to the ipseity server at the given address.
func NewStreamClient(addr string) (pb.StreamIdentityClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not dial %s: %s", addr, err)
	}

	return pb.NewStreamIdentityClient(conn), nil
}
