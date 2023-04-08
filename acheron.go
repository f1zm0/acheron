package acheron

import (
	"github.com/f1zm0/acheron/internal/resolver"
	"github.com/f1zm0/acheron/internal/resolver/ssnsort"
	"github.com/f1zm0/acheron/pkg/hashing"
)

type Acheron struct {
	resolver resolver.Resolver
}

type (
	// Option is a configuration option which can be used within
	// a call to the constructor to configure the Acheron instance.
	Option func(*options)
)

type options struct {
	hasher hashing.Hasher
}

// New returns a new Acheron instance that can be used as a proxy to perform
// indirect syscalls for native api functions, or an error if the initialization fails.
func New(opts ...Option) (*Acheron, error) {
	// defaults
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

// WithHashFunction returns an Option that sets a custom hashing (or obfuscation)
// function that will be used when resolving native api procedures by hash.
func WithHashFunction(f hashing.Hasher) Option {
	return func(o *options) {
		o.hasher = f
	}
}
