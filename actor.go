package ipseity

import (
	"context"
	"fmt"
	"sync"

	"github.com/bbengfort/ipseity/pb"
)

// Buffer size to instantiate actor channels with
const actorEventBufferSize = 1024

// Actor objects listen for events (messages) and then can create more actors,
// send more messages or make local decisions that modify their own private
// state. Actors implement lockless concurrent operation (indeed, none of the
// structs in this package implement mutexes and are not thread safe
// independently). Concurrency here is based on the fact that only a single
// actor is initialized and reads event objects one at a time off of a
// buffered channel. All actor methods should be private as a result so they
// are not called from other threads.
type Actor interface {
	Server
	Dispatch(Event) error // Outside callers can dispatch events to the actor
	Handle(Event) error   // Handler method for each event in sequence
}

//===========================================================================
// Non-Blocking Actor
//===========================================================================

// ActorServer is a simple implementation of an actor object that orders
// events through a channel and responds to them on demand.
type ActorServer struct {
	server
	events   chan Event
	sequence int64
}

// Listen for events, handling them with the default callback handler. If the
// callback returns an error, then Listen will return with that error. If the
// actor is closed externally, then Listen will finish all remaining events
// and return nil.
func (a *ActorServer) Listen(addr string) error {
	a.events = make(chan Event, actorEventBufferSize)

	if err := a.server.Listen(addr); err != nil {
		return err
	}

	pb.RegisterIdentityServer(a.srv, a)
	go a.srv.Serve(a.sock)

	// Continue reading events off the channel until its closed
	for event := range a.events {
		if err := a.Handle(event); err != nil {
			return err
		}
	}

	return nil
}

// Next returns the next identity on the server
func (a *ActorServer) Next(ctx context.Context, in *pb.IdentityRequest) (*pb.IdentityReply, error) {
	identity := make(chan int64, 1)

	// Create an event to dispatch to the actor
	event := &event{
		etype: IdentityEvent, source: identity, value: in,
	}

	if err := a.Dispatch(event); err != nil {
		return nil, err
	}

	return &pb.IdentityReply{
		Key: in.Key, Identity: <-identity,
	}, nil
}

// Dispatch an event on the actor for the listener to handle.
func (a *ActorServer) Dispatch(e Event) error {
	a.events <- e
	return nil
}

// Handle each event by passing it to the callback function.
func (a *ActorServer) Handle(e Event) error {
	if e.Type() != IdentityEvent {
		return fmt.Errorf("cannot handle event type %s", e.Type())
	}

	a.sequence++
	source := e.Source().(chan int64)
	source <- a.sequence
	return nil
}

//===========================================================================
// Blocking Actor
//===========================================================================

// LockerServer is another mutex based, blocking server but does so with the
// added overhead of event handling.
type LockerServer struct {
	sync.RWMutex
	server
	sequence int64
}

// Listen on the specified address for identity requests.
func (a *LockerServer) Listen(addr string) error {
	if err := a.server.Listen(addr); err != nil {
		return err
	}

	pb.RegisterIdentityServer(a.srv, a)
	return a.srv.Serve(a.sock)
}

// Next returns the next identity on the server
func (a *LockerServer) Next(ctx context.Context, in *pb.IdentityRequest) (*pb.IdentityReply, error) {
	identity := make(chan int64, 1)

	// Create an event to dispatch to the actor
	event := &event{
		etype: IdentityEvent, source: identity, value: in,
	}

	if err := a.Dispatch(event); err != nil {
		return nil, err
	}

	return &pb.IdentityReply{
		Key: in.Key, Identity: <-identity,
	}, nil
}

// Dispatch an event to the locker actor
func (a *LockerServer) Dispatch(e Event) error {
	a.Lock()
	defer a.Unlock()
	return a.Handle(e)
}

// Handle the dispatched events to the locker actor
func (a *LockerServer) Handle(e Event) error {
	if e.Type() != IdentityEvent {
		return fmt.Errorf("cannot handle event type %s", e.Type())
	}

	a.sequence++
	source := e.Source().(chan int64)
	source <- a.sequence
	return nil
}
