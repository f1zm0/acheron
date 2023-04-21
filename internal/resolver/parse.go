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

// ParseNtdllModule returns a slice of Syscall structs for all Zw* syscalls from in-memory ntdll.dll.
func ParseNtdllModule(hashFn hashing.HashFunction) []*Syscall {
	baseAddr := getNtdllBaseAddr()
	exportsBaseAddr := getModuleExportsDirAddr(baseAddr)
	numberOfNames := getExportsNumberOfNames(exportsBaseAddr)
	addressOfFunctions := getExportsAddressOfFunctions(baseAddr, exportsBaseAddr)
	addressOfNames := getExportsAddressOfNames(baseAddr, exportsBaseAddr)
	addressOfNameOrdinals := getExportsAddressOfNameOrdinals(baseAddr, exportsBaseAddr)

	sysStubs := make([]*Syscall, 0, numberOfNames/4) // Zw* < 25% of all exports
	for i := uint32(0); i < numberOfNames; i++ {
		fn := memory.ReadCStringAt(baseAddr, memory.ReadDwordAtOffset(addressOfNames, i*4))
		if fn[0] == 'Z' && fn[1] == 'w' {
			fn[0] = 'N'
			fn[1] = 't'
			nameOrd := memory.ReadWordAtOffset(addressOfNameOrdinals, i*2)
			rva := memory.ReadDwordAtOffset(addressOfFunctions, uint32(nameOrd*4))

			sysStubs = append(sysStubs, &Syscall{
				NameHash: hashFn(fn),
				RVA:      rva,
				VA:       memory.RVA2VA(baseAddr, rva),
			})
		}
	}
	return sysStubs
}
