package rvasort

import (
	"sort"

	"github.com/f1zm0/acheron/internal/resolver"
	"github.com/f1zm0/acheron/pkg/hashing"
)

type ssnSortResolver struct {
	hasher           hashing.HashFunction
	syscallStubs     map[uint64]*resolver.Syscall
	cleanTrampolines []uintptr
}

var _ resolver.Resolver = (*ssnSortResolver)(nil)

// NewResolver returns a new resolver that uses the given hash function to resolve syscalls SSNs.
func NewResolver(h hashing.HashFunction) (resolver.Resolver, error) {
	r := &ssnSortResolver{
		hasher: h,
	}
	ss := resolver.ParseNtdllModule(r.hasher) // returns a slice of Syscall structs

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].RVA < ss[j].RVA // sort stubs by RVA
	})

	// search clean syscall;ret gadgets to use as syscall trampolines in stubs memory range
	for _, st := range ss {
		if trampoline := getTrampoline(st.VA); trampoline != uintptr(0) {
			st.TrampolineAddr = trampoline
			r.cleanTrampolines = append(r.cleanTrampolines, trampoline)
		}
	}

	r.syscallStubs = make(map[uint64]*resolver.Syscall, len(ss))
	for idx, st := range ss {
		st.SSN = uint16(idx)

		// keep its default trampoline if it was unhooked, otherwise use one of the clean ones
		if st.TrampolineAddr == uintptr(0) {
			ss[idx].TrampolineAddr = r.cleanTrampolines[0] // pick random one?
		}

		r.syscallStubs[ss[idx].NameHash] = ss[idx]
	}

	return r, nil
}

// GetSyscall returns the Syscall struct for the given function hash, or nil if not found.
func (r *ssnSortResolver) GetSyscall(fnHash uint64) *resolver.Syscall {
	if v, ok := r.syscallStubs[fnHash]; ok {
		return v
	}
	return nil
}
