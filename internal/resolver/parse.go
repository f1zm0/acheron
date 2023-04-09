package resolver

import (
	"github.com/f1zm0/acheron/pkg/hashing"
	"github.com/f1zm0/acheron/pkg/memory"
)

type NtModule struct {
	BaseAddr              uintptr
	ExportsBaseAddr       uintptr
	NumberOfNames         uint32
	AddressOfFunctions    uintptr
	AddressOfNames        uintptr
	AddressOfNameOrdinals uintptr
	ZwStubs               map[int64]*ZwStub
}

type ZwStub struct {
	RVA uint32
	VA  uintptr
	SSN uint16
}

// ParseNtdllModule returns a NtModule struct with the relevant information
// about the in-memory ntdll.dll module.
func ParseNtdllModule(hashFn hashing.Hasher) *NtModule {
	var m NtModule

	m.BaseAddr = getNtdllBaseAddr()
	m.ExportsBaseAddr = getModuleExportsDirAddr(m.BaseAddr)
	m.NumberOfNames = getExportsNumberOfNames(m.ExportsBaseAddr)
	m.AddressOfFunctions = getExportsAddressOfFunctions(m.BaseAddr, m.ExportsBaseAddr)
	m.AddressOfNames = getExportsAddressOfNames(m.BaseAddr, m.ExportsBaseAddr)
	m.AddressOfNameOrdinals = getExportsAddressOfNameOrdinals(m.BaseAddr, m.ExportsBaseAddr)

	m.ZwStubs = make(
		map[int64]*ZwStub,
		m.NumberOfNames/4, // Zw* functions are less than ~25% of the total so we can save some memory
	)

	for i := uint32(0); i < m.NumberOfNames; i++ {
		fn := memory.ReadCStringAt(m.BaseAddr, memory.ReadDwordAt(m.AddressOfNames, i*4))
		if fn[0] != 'Z' || fn[1] != 'w' {
			continue
		}
		fnHash := hashFn.HashByteString(fn)

		nameOrd := memory.ReadWordAt(m.AddressOfNameOrdinals, i*2)
		rva := memory.ReadDwordAt(m.AddressOfFunctions, uint32(nameOrd*4))

		m.ZwStubs[fnHash] = &ZwStub{
			RVA: rva,
			VA:  memory.RVA2VA(m.BaseAddr, uint32(rva)),
		}
	}

	return &m
}
