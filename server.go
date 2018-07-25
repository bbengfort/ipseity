package ipseity

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/bbengfort/ipseity/pb"
	"google.golang.org/grpc"
)

// Server is the interface for the many implementations of the identity server.
type Server interface {
	Listen(addr string) error
}

// Embedable server type that implements the listen method.
type server struct {
	sock net.Listener
	srv  *grpc.Server
}

// Listen on the specified address for identity requests.
func (s *server) Listen(addr string) (err error) {
	if s.sock, err = net.Listen("tcp", addr); err != nil {
		return fmt.Errorf("could not listen on %s", addr)
	}

	s.srv = grpc.NewServer()
	return nil
}

// SimpleServer returns unique numeric identities on demand to clients by
// locking the server itself and incrementing an internal sequence number
// before responding to the client (then unlocking).
type SimpleServer struct {
	sync.Mutex
	server
	sequence int64
}

// Listen on the specified address for identity requests.
func (s *SimpleServer) Listen(addr string) error {
	if err := s.server.Listen(addr); err != nil {
		return err
	}

	pb.RegisterIdentityServer(s.srv, s)
	return s.srv.Serve(s.sock)
}

// Next returns the next identity on the server
func (s *SimpleServer) Next(ctx context.Context, in *pb.IdentityRequest) (*pb.IdentityReply, error) {
	s.Lock()
	defer s.Unlock()

	s.sequence++
	return &pb.IdentityReply{
		Key: in.Key, Identity: s.sequence,
	}, nil
}
