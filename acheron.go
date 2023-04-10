package acheron

import (
	"github.com/f1zm0/acheron/internal/resolver"
	"github.com/f1zm0/acheron/internal/resolver/rvasort"
	"github.com/f1zm0/acheron/pkg/hashing"
)

type Acheron struct {
	resolver resolver.Resolver
}

type (
	// Option is a configuration option to configure the Acheron instance.
	Option func(*options)
)

type options struct {
	hasher hashing.HashFunction
}

// New returns a new Acheron instance that can be used as a proxy to perform
// indirect syscalls for native api functions, or an error if the initialization fails.
func New(opts ...Option) (*Acheron, error) {
	options := &options{
		hasher: hashing.DJB2,
	}
	for _, o := range opts {
		o(options)
	}

	if r, err := rvasort.NewResolver(options.hasher); err != nil {
		return nil, err
	} else {
		return &Acheron{
			resolver: r,
		}, nil
	}
}

// WithHashFunction returns an Option that sets a custom hashing (or obfuscation)
// function that will be used when resolving native api procedures by hash.
func WithHashFunction(f hashing.HashFunction) Option {
	return func(o *options) {
		o.hasher = f
	}
}
