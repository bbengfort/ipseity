package ipseity

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// PackageVersion for the Ipseity server
const PackageVersion = "0.3"

// New returns an identity server based on the specified type. By default a
// SimpleServer is returned if stype="". Other options include simple,
// sequence, stream, actor, or locker.
func New(stype string) (Server, error) {
	stype = strings.ToLower(stype)
	stype = strings.TrimSpace(stype)

	switch stype {
	case "", "simple":
		return new(SimpleServer), nil
	case "sequence":
		return new(SequenceServer), nil
	case "stream":
		return new(StreamServer), nil
	case "actor":
		return new(ActorServer), nil
	case "locker":
		return new(LockerServer), nil
	default:
		return nil, fmt.Errorf("unknown stype '%s' use simple, sequence, actor, or locker", stype)
	}
}
