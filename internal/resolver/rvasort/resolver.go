package rvasort

import (
	"errors"
	"fmt"
	"sort"

	"github.com/f1zm0/acheron/internal/resolver"
	"github.com/f1zm0/acheron/pkg/hashing"
)

type ssnSortResolver struct {
	hasher           hashing.HashFunction
	zwStubs          map[uint64]*resolver.Syscall
	cleanTrampolines []uintptr
}

var _ resolver.Resolver = (*ssnSortResolver)(nil)

func NewResolver(h hashing.HashFunction) (resolver.Resolver, error) {
	r := &ssnSortResolver{
		hasher: h,
	}
	zws := resolver.ParseNtdllModule(r.hasher) // returns a slice of Syscall structs

	sort.Slice(zws, func(i, j int) bool {
		return zws[i].RVA < zws[j].RVA // sort stubs by RVA
	})

	// search clean syscall;ret gadgets to use as syscall trampolines in stubs memory range
	for _, st := range zws {
		if trampoline := getTrampoline(st.VA); trampoline != uintptr(0) {
			st.TrampolineAddr = trampoline
			r.cleanTrampolines = append(r.cleanTrampolines, trampoline)
		}
	}

	r.zwStubs = make(map[uint64]*resolver.Syscall, len(zws))
	for idx, st := range zws {
		st.SSN = uint16(idx)

		// keep its default trampoline if it was unhooked, otherwise use one of the clean ones
		if st.TrampolineAddr == uintptr(0) {
			zws[idx].TrampolineAddr = r.cleanTrampolines[0] // pick random one?
		}

		// add to zwStubs map
		r.zwStubs[zws[idx].NameHash] = zws[idx]
	}

	return r, nil
}

func (r *ssnSortResolver) GetSyscall(fnHash uint64) (*resolver.Syscall, error) {
	if v, ok := r.zwStubs[fnHash]; ok {
		return v, nil
	}
	return nil, errors.New(fmt.Sprintf("syscall with hash %d not found", fnHash))
}
