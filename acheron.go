package acheron

import (
	"github.com/f1zm0/acheron/internal/resolver"
	"github.com/f1zm0/acheron/internal/resolver/rvasort"
	"github.com/f1zm0/acheron/pkg/hashing"
)

// Acheron is the main struct of the acheron package.
type Acheron struct {
	resolver     resolver.Resolver
	hashFunction hashing.HashFunction
}

// Option is a configuration option to configure the Acheron instance.
type Option func(*options)

type options struct {
	hashFunction hashing.HashFunction
}

// New returns a new Acheron instance with the given options, or an error if initialization fails.
func New(opts ...Option) (*Acheron, error) {
	options := &options{
		hashFunction: hashing.XorDjb2Hash, // default
	}
	for _, o := range opts {
		o(options)
	}

	if r, err := rvasort.NewResolver(options.hashFunction); err != nil {
		return nil, err
	} else {
		return &Acheron{
			resolver:     r,
			hashFunction: options.hashFunction,
		}, nil
	}
}

// WithHashFunction returns an Option that sets a custom hashing or obfuscation function.
func WithHashFunction(f hashing.HashFunction) Option {
	return func(o *options) {
		o.hashFunction = f
	}
}

// HashString is a helper function to hash a string which can be used as first arg for Syscall.
func (a *Acheron) HashString(s string) uint64 {
	return a.hashFunction([]byte(s))
}
