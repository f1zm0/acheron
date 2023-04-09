package resolver

// FindSyscallRetGadgets finds syscall;ret gadgets in ntdll.dll
// that can be "recycled" to ensure syscalls goes through ntdll.
func FindSyscallRetGadgets(stubs map[int64]*ZwStub) []uintptr {
	// TODO: do implementation
	return []uintptr{}
}
