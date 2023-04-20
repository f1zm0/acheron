package acheron

import (
	"errors"
	"fmt"

	"github.com/f1zm0/acheron/internal/resolver"
	"github.com/f1zm0/acheron/internal/resolver/rvasort"
	"github.com/f1zm0/acheron/pkg/hashing"
)

// Acheron is the main struct of the acheron package.
type Acheron struct {
	resolver     resolver.Resolver
	hashFunction hashing.HashFunction
}

type (
	// Option is a configuration option to configure the Acheron instance.
	Option func(*options)
)

type options struct {
	hashFunction hashing.HashFunction
}

// WithHashFunction returns an Option that sets a custom hashing or obfuscation function.
func WithHashFunction(f hashing.HashFunction) Option {
	return func(o *options) {
		o.hashFunction = f
	}
}

// New returns a new Acheron instance with the given options, or an error if initialization fails.
func New(opts ...Option) (*Acheron, error) {
	options := &options{
		hashFunction: hashing.DJB2,
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

// HashString is a helper function to hash a string which can be used as first arg for Syscall.
func (a *Acheron) HashString(s string) uint64 {
	return a.hashFunction([]byte(s))
}

// Syscall executes an indirect syscall with the given function hash and arguments.
// Returns the error code returned by the syscall is something goes wrong.
func (a *Acheron) Syscall(fnHash uint64, args ...uintptr) error {
	sys, err := a.resolver.GetSyscall(fnHash)
	if err != nil {
		return err
	}
	if errCode := execIndirectSyscall(sys.SSN, sys.TrampolineAddr, args...); errCode != 0 { // !NT_SUCCESS
		return errors.New(fmt.Sprintf("syscall failed with error code %d", errCode))
	}
	return nil
}

func execIndirectSyscall(ssn uint16, gateAddr uintptr, argh ...uintptr) (errcode uint32)
