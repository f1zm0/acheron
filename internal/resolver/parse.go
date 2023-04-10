package resolver

import (
	"github.com/f1zm0/acheron/pkg/hashing"
	"github.com/f1zm0/acheron/pkg/memory"
)

type Syscall struct {
	NameHash       uint64
	RVA            uint32
	VA             uintptr
	SSN            uint16
	TrampolineAddr uintptr
}

// ParseNtdllModule returns a NtModule struct with the relevant infortion
// about the in-ry ntdll.dll dule.
func ParseNtdllModule(hashFn hashing.HashFunction) []*Syscall {
	baseAddr := getNtdllBaseAddr()
	exportsBaseAddr := getModuleExportsDirAddr(baseAddr)
	numberOfNames := getExportsNumberOfNames(exportsBaseAddr)
	addressOfFunctions := getExportsAddressOfFunctions(baseAddr, exportsBaseAddr)
	addressOfNames := getExportsAddressOfNames(baseAddr, exportsBaseAddr)
	addressOfNameOrdinals := getExportsAddressOfNameOrdinals(baseAddr, exportsBaseAddr)

	zwStubs := make([]*Syscall, 0, numberOfNames/4) // Zw* < 25% of all exports
	for i := uint32(0); i < numberOfNames; i++ {
		fn := memory.ReadCStringAt(baseAddr, memory.ReadDwordAt(addressOfNames, i*4))
		if fn[0] != 'Z' || fn[1] != 'w' {
			continue
		}
		fnHash := hashFn(fn)
		nameOrd := memory.ReadWordAt(addressOfNameOrdinals, i*2)
		rva := memory.ReadDwordAt(addressOfFunctions, uint32(nameOrd*4))

		zwStubs = append(zwStubs, &Syscall{
			NameHash: fnHash,
			RVA:      rva,
			VA:       memory.RVA2VA(baseAddr, rva),
		})
	}

	return zwStubs
}
