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
	zwStubs := resolver.ParseNtdllModule(r.hasher)

	sort.Slice(zwStubs, func(i, j int) bool {
		return zwStubs[i].RVA < zwStubs[j].RVA // sort stubs by RVA
	})

	// search clean syscall;ret gadgets to use as syscall trampolines in stubs memory range
	for _, st := range zwStubs {
		if trampoline := getTrampoline(st.VA); trampoline != uintptr(0) {
			st.TrampolineAddr = trampoline
			r.cleanTrampolines = append(r.cleanTrampolines, trampoline)
		}
	}

	r.zwStubs = make(map[uint64]*resolver.Syscall, len(zwStubs))
	for idx, st := range zwStubs {
		st.SSN = uint16(idx)

		// keep its default trampoline if it was unhooked, otherwise use one of the clean ones
		if st.TrampolineAddr == uintptr(0) {
			zwStubs[idx].TrampolineAddr = r.cleanTrampolines[0] // pick random one?
		}

		r.zwStubs[zwStubs[idx].NameHash] = zwStubs[idx]
		fmt.Printf(
			"NameHash: %d | VA: 0x%x | SSN: %d | Trampoline: 0x%x\r\n\r\n",
			zwStubs[idx].NameHash,
			zwStubs[idx].VA,
			zwStubs[idx].SSN,
			zwStubs[idx].TrampolineAddr,
		)
	}

	return r, nil
}

func (r *ssnSortResolver) GetSyscall(fnHash uint64) (*resolver.Syscall, error) {
	if v, ok := r.zwStubs[fnHash]; ok {
		return v, nil
	}
	return nil, errors.New("could not find stub with the provided hash")
}
