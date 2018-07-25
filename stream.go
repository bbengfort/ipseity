package ipseity

import (
	"io"

	"github.com/bbengfort/ipseity/pb"
)

// StreamServer implements the SequenceServer with a bidirection stream
// instead of using the standard RPC format.
type StreamServer struct {
	server
	sequence *Sequence
}

// Listen on the specified address for identity requests.
func (s *StreamServer) Listen(addr string) error {
	s.sequence = new(Sequence)

	if err := s.server.Listen(addr); err != nil {
		return err
	}

	pb.RegisterStreamIdentityServer(s.srv, s)
	return s.srv.Serve(s.sock)
}

// Next opens a stream from the client that can request identities.
func (s *StreamServer) Next(stream pb.StreamIdentity_NextServer) (err error) {
	for {
		var in *pb.IdentityRequest
		if in, err = stream.Recv(); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		out := &pb.IdentityReply{Key: in.Key, Identity: s.sequence.Next()}
		if err = stream.Send(out); err != nil {
			return err
		}
	}
}
