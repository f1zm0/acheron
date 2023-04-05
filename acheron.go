package acheron

import (
	"github.com/f1zm0/acheron/internal/resolver"
	"github.com/f1zm0/acheron/internal/resolver/ssnsort"
	"github.com/f1zm0/acheron/pkg/hashing"
)

type Acheron struct {
	resolver resolver.Resolver
}

// New returns a new Acheron instance that can be used as a proxy to perform
// indirect syscalls for native api functions, or an error if the initialization fails.
func New(opts ...Option) (*Acheron, error) {
	options := &options{
		hasher: hashing.NewDjb2(),
	}
	for _, o := range opts {
		o(options)
	}

	if r, err := ssnsort.NewResolver(options.hasher); err != nil {
		return nil, err
	} else {
		return &Acheron{
			resolver: r,
		}, nil
	}
}
