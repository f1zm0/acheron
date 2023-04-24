//go:build windows
// +build windows

package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"

	"github.com/f1zm0/acheron"
)

// custom encoding/hashing function implementation that complies with the acheron.HashFunction type
func customXORSHA1(s []byte) uint64 {
	key := []byte{0xde, 0xad, 0xbe, 0xef}
	for i := 0; i < len(s); i++ {
		s[i] ^= key[i%len(key)]
	}
	hash := sha1.Sum(s)
	return binary.LittleEndian.Uint64(hash[:])
}

func main() {
	// creates Acheron instance, resolves SSNs, collects clean trampolines in ntdll.dlll, etc.
	ach, err := acheron.New(
		// Customize instance with fucntional options
		acheron.WithHashFunction(customXORSHA1),
	)
	if err != nil {
		panic(err)
	}

	// having used a custom func, you can now calc the hashes
	// using both <acheron_instance>.HashString() or customXorFn()

	ntqsi := ach.HashString("NtSetQueryInformationProcess")
	fmt.Printf("NtSetQueryInformationProcess: 0x%x\r\n", ntqsi)

	ntavm := customXORSHA1([]byte("NtAllocateVirtualMemory"))
	fmt.Printf("NtAllocateVirtualMemory: 0x%x\r\n", ntavm)
}
