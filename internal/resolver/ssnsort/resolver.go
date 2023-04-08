package ssnsort

import (
	"errors"

	"github.com/f1zm0/acheron/internal/resolver"
	wt "github.com/f1zm0/acheron/internal/types"
	"github.com/f1zm0/acheron/pkg/hashing"
)

type ssnSortResolver struct {
	// hashing provider
	hasher hashing.Hasher

	// map of Zw* InMemProc structs indexed by their name's hash
	zwStubs map[int64]wt.InMemProc

	// slice with addresses of clean gadgets
	safeGates []uintptr
}

var _ resolver.Resolver = (*ssnSortResolver)(nil)

func NewResolver(h hashing.Hasher) (resolver.Resolver, error) {
	r := &ssnSortResolver{}
	if err := r.init(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *ssnSortResolver) init() error {
	// var zwStubs []wt.InMemProc

	_, err := resolver.GetNtdllModuleHandle()
	if err != nil {
		return err
	}

	// exports, err := hNtdll.File.Exports()
	// if err != nil {
	// 	return err
	// }
	// for _, exp := range exports {
	// 	memAddr := int64(hNtdll.BaseAddr) + int64(exp.VirtualAddress)
	// 	r.safeGates = resolver.FindSyscallRetGadgets(hNtdll)
	// 	if len(r.safeGates) == 0 {
	// 		return errors.New("could not found syscall;ret gadgets")
	// 	}

	// 	if strings.HasPrefix(exp.Name, "Zw") {
	// 		zwStubs = append(zwStubs, wt.InMemProc{
	// 			Name:     exp.Name,
	// 			BaseAddr: uintptr(memAddr),
	// 		})
	// 	}
	// }

	// sort.Slice(zwStubs, func(i, j int) bool {
	// 	return zwStubs[i].BaseAddr < zwStubs[j].BaseAddr
	// })

	// for idx := range zwStubs {
	// 	zwStubs[idx].SSN = idx
	// 	r.zwStubs[r.hasher.HashString(zwStubs[idx].Name)] = zwStubs[idx]
	// }

	return nil
}

// GetSyscallSSN returns the syscall ID of a native API function by its djb2 hash.
// If the function is not found, 0 is returned.
func (r *ssnSortResolver) GetSyscallSSN(fnHash int64) (uint16, error) {
	if v, ok := r.zwStubs[fnHash]; ok {
		return uint16(v.SSN), nil
	}
	return 0, errors.New("could not find SSN")
}

func (r *ssnSortResolver) GetSafeGate() uintptr {
	// FIXME: this panics as safeGates is empty
	return r.safeGates[0]
}
