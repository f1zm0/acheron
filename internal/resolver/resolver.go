package resolver

type Resolver interface {
	// GetSyscallSSN returns the syscall SSN.
	GetSyscallSSN(funcNameHash int64) (uint16, error)

	// GetSafeGate returns the address of an unhooked syscall;ret gadget in ntdll.dll
	GetSafeGate() uintptr
}
