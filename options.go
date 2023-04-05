package acheron

import (
	"github.com/f1zm0/acheron/pkg/hashing"
)

// Option is a configuration option which can be used within a
// call to the formatter constructor.
type (
	Option  func(*options)
	Options []Option
)

type options struct {
	hasher hashing.Hasher
}

// WithHashFunction returns an Option that sets a custom hashing (or obfuscation)
// function that will be used when resolving native api procedures by hash.
func WithHashFunction(f hashing.Hasher) Option {
	return func(o *options) {
		o.hasher = f
	}
}
