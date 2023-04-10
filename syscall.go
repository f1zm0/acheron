package acheron

import (
	"errors"
	"fmt"
)

// Syscall executes a syscall with the given function hash and arguments.
// Returns the error code and an error if the syscall failed.
func (a *Acheron) Syscall(fnHash int64, args ...uintptr) error {
	sys, err := a.resolver.GetSyscall(fnHash)
	if err != nil {
		return err
	}
	if errCode := execIndirectSyscall(sys.SSN, sys.TrampolineAddr, args...); errCode != 0 {
		return errors.New(fmt.Sprintf("syscall failed with error code %d", errCode))
	}
	return nil
}

// execIndirectSyscall function signature for go-asm impelementation.
// returns 0 if the syscall was successful or an error code if the operation failed.
func execIndirectSyscall(ssn uint16, gateAddr uintptr, argh ...uintptr) (errcode uint32)
