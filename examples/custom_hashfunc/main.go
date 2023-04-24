//go:build windows
// +build windows

package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"

	"github.com/f1zm0/acheron"
)

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
	acheron, err := acheron.New(
		// Customize instance with fucntional options
		acheron.WithHashFunction(customXORSHA1),
	)
	if err != nil {
		panic(err)
	}

	// you can calc the hashes using both acheron.HashString or customXorFn
	ntqsi := acheron.HashString("NtSetQueryInformationProcess")
	fmt.Printf("NtSetQueryInformationProcess: 0x%x\r\n", ntqsi)
}
