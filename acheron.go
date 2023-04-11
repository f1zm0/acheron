package acheron

import (
	"errors"
	"fmt"

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

// WithHashFunction returns an Option that sets a custom hashing or obfuscation function.
func WithHashFunction(f hashing.HashFunction) Option {
	return func(o *options) {
		o.hasher = f
	}
}

// New returns a new Acheron instance with the given options, or an error if initialization fails.
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

// Syscall executes an indirect syscall with the given function hash and arguments.
// Returns the error code returned by the syscall is something goes wrong.
func (a *Acheron) Syscall(fnHash uint64, args ...uintptr) error {
	sys, err := a.resolver.GetSyscall(fnHash)
	if err != nil {
		return err
	}
	if errCode := execIndirectSyscall(sys.SSN, sys.TrampolineAddr, args...); errCode != 0 {
		return errors.New(fmt.Sprintf("syscall failed with error code %d", errCode))
	}
	return nil
}

func execIndirectSyscall(ssn uint16, gateAddr uintptr, argh ...uintptr) (errcode uint32)
