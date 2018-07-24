package ipseity

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/bbengfort/ipseity/pb"
	"google.golang.org/grpc"
)

// PackageVersion for the Ipseity server
const PackageVersion = "0.1"

// New returns an identity server.
func New() *IdentityServer {
	return new(IdentityServer)
}

// IdentityServer returns unique numeric identities on demand to clients.
type IdentityServer struct {
	sync.Mutex
	sequence int64
}

// Listen on the specified address for identity requests.
func (s *IdentityServer) Listen(addr string) (err error) {
	var sock net.Listener
	if sock, err = net.Listen("tcp", addr); err != nil {
		return fmt.Errorf("could not listen on %s", addr)
	}
	defer sock.Close()

	srv := grpc.NewServer()
	pb.RegisterIdentityServer(srv, s)

	return srv.Serve(sock)
}

// Next returns the next identity on the server
func (s *IdentityServer) Next(ctx context.Context, in *pb.IdentityRequest) (*pb.IdentityReply, error) {
	s.Lock()
	defer s.Unlock()

	s.sequence++
	return &pb.IdentityReply{
		Key: in.Key, Identity: s.sequence,
	}, nil
}
