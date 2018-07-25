package ipseity

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/bbengfort/ipseity/pb"
	"golang.org/x/sync/errgroup"
)

// NewBenchmark creates the benchmark for the specified number of clients and
// number of messages per client, then runs the benchmark.
func NewBenchmark(nClients, msgsPerClient int, addr, stype string) (*Benchmark, error) {
	bench := &Benchmark{stype: stype, addr: addr, nClients: nClients, messages: msgsPerClient}
	err := bench.Run()
	return bench, err
}

// Benchmark implements several go routines sending chat messages in their
// own connections concurrently for a fixed number of messages, then returns
// the observed throughput from the client side.
type Benchmark struct {
	stype    string        // The type of ipseity server
	addr     string        // The address of the server to connect to
	nClients int           // Number of concurrent clients
	messages int           // Number of messages per client
	duration time.Duration // Total amount of time it took to send all messages
}

// Run the benchmark for the specified number of clients and messages.
func (b *Benchmark) Run() error {

	group := new(errgroup.Group)
	start := time.Now()

	for i := 0; i < b.nClients; i++ {
		group.Go(func() error {
			key := fmt.Sprintf("%04X", rand.Intn(10000))
			client, err := NewClient(b.addr)
			if err != nil {
				return err
			}

			for j := 0; j < b.messages; j++ {
				ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
				if _, err := client.Next(ctx, &pb.IdentityRequest{Key: key}); err != nil {
					cancel()
					return err
				}
				cancel()
			}

			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return err
	}
	b.duration = time.Since(start)
	return nil
}

// Throughput returns the number of operations per second.
func (b *Benchmark) Throughput() float64 {
	if b.duration == 0 {
		return 0.0
	}

	return float64(b.NumMessages()) / b.duration.Seconds()
}

// NumClients returns the number of concurrent clients.
func (b *Benchmark) NumClients() uint64 {
	return uint64(b.nClients)
}

// NumMessages returns the total number of messages sent.
func (b *Benchmark) NumMessages() uint64 {
	return b.NumClients() * uint64(b.messages)
}

// Duration returns the amount of time it took to send all messages.
func (b *Benchmark) Duration() time.Duration {
	return b.duration
}

// String returns a CSV string of the benchmark data.
// server,n_clients,n_messages,duration,throughput
func (b *Benchmark) String() string {
	return fmt.Sprintf("%s,%d,%d,%s,%0.4f",
		b.stype, b.NumClients(), b.NumMessages(), b.Duration(), b.Throughput(),
	)
}
