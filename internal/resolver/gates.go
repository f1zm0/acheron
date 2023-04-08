package resolver

import (
	wt "github.com/f1zm0/acheron/internal/types"
)

const SYSCALL_STUB_SIZE = 0x20

// FindSyscallRetGadgets finds syscall;ret gadgets in ntdll.dll
// that can be "recycled" to ensure syscalls goes through ntdll.
func FindSyscallRetGadgets(hNtdll *wt.PEModule) []uintptr {
	// TODO: do implementation
	return []uintptr{}
}
