package ipseity

import (
	"context"
	"sync"

	"github.com/bbengfort/ipseity/pb"
)

// Sequence is a thread-safe monotonically increasing identity container.
type Sequence struct {
	sync.RWMutex
	current int64
}

// Next returns the next value in the sequence with a write lock.
func (s *Sequence) Next() int64 {
	s.Lock()
	defer s.Unlock()
	s.current++
	return s.current
}

// Current returns the current value in the sequence with a read lock.
func (s *Sequence) Current() int64 {
	s.RLock()
	defer s.RUnlock()
	return s.current
}

// SequenceServer uses a Sequence to return the identity rather than a server
// mutex to protect the variable being read.
type SequenceServer struct {
	server
	sequence *Sequence
}

// Listen on the specified address for identity requests.
func (s *SequenceServer) Listen(addr string) error {
	s.sequence = new(Sequence)

	if err := s.server.Listen(addr); err != nil {
		return err
	}

	pb.RegisterIdentityServer(s.srv, s)
	return s.srv.Serve(s.sock)
}

// Next returns the next identity on the server
func (s *SequenceServer) Next(ctx context.Context, in *pb.IdentityRequest) (*pb.IdentityReply, error) {
	return &pb.IdentityReply{
		Key: in.Key, Identity: s.sequence.Next(),
	}, nil
}
