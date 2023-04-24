//go:build windows
// +build windows

package main

import (
	"unsafe"

	"github.com/f1zm0/acheron"
	"golang.org/x/sys/windows"
)

const nullptr = uintptr(0)

func main() {
	bufferSize := uint32(0)

	// creates Acheron instance, resolves SSNs, collects clean trampolines in ntdll.dlll, etc.
	ach, err := acheron.New()
	if err != nil {
		panic(err)
	}

	// make indirect syscall for NtQuerySystemInformation to figure out the buffer size
	_ = ach.Syscall(
		ach.HashString("NtQuerySystemInformation"), // function name hash
		0x5,                                  // _In_ SYSTEM_INFORMATION_CLASS SystemInformationClass
		0,                                    // _Out_ PVOID SystemInformation
		uintptr(bufferSize),                  // _In_ ULONG SystemInformationLength
		uintptr(unsafe.Pointer(&bufferSize)), // _Out_opt_ PULONG ReturnLength
	)

	buf := make([]byte, bufferSize)

	// make indirect syscall to get the process list (in buf)
	if err := ach.Syscall(
		ach.HashString("NtQuerySystemInformation"),
		0x5, // SystemProcessInformation
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(bufferSize),
		nullptr,
	); err != nil {
		panic(err)
	}

	// print each process ID and Name
	offset := 0
	for {
		p := (*windows.SYSTEM_PROCESS_INFORMATION)(unsafe.Pointer(&buf[offset]))
		if p.ImageName.String() != "" {
			println("PID: ", p.UniqueProcessID, " Name: ", p.ImageName.String())
		}
		if p.NextEntryOffset == 0 {
			break
		}
		offset += int(p.NextEntryOffset)
	}
}
