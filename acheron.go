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

// WithHashFunction returns an Option that sets a custom hashing (or obfuscation)
// function that will be used when resolving native api procedures by hash.
func WithHashFunction(f hashing.HashFunction) Option {
	return func(o *options) {
		o.hasher = f
	}
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

// Syscall executes a syscall with the given function hash and arguments.
// Returns the error code and an error if the syscall failed.
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

// execIndirectSyscall function signature for go-asm impelementation.
// returns 0 if the syscall was successful or an error code if the operation failed.
func execIndirectSyscall(ssn uint16, gateAddr uintptr, argh ...uintptr) (errcode uint32)
