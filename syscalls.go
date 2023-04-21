package acheron

import (
	"errors"
	"fmt"
)

func execIndirectSyscall(ssn uint16, gateAddr uintptr, argh ...uintptr) (errcode uint32)

func execDirectSyscall(ssn uint16, argh ...uintptr) (errcode uint32)

// Syscall executes an indirect syscall with the given function hash and arguments. Returns the error code if it fails.
func (a *Acheron) Syscall(fnHash uint64, args ...uintptr) error {
	sys, err := a.resolver.GetSyscall(fnHash)
	if err != nil {
		return err
	}
	if errCode := execIndirectSyscall(sys.SSN, sys.TrampolineAddr, args...); errCode != 0 { // !NT_SUCCESS
		return errors.New(fmt.Sprintf("syscall failed with error code %d", errCode))
	}
	return nil
}

// DirectSyscall executes a direct syscall with the given function hash and arguments. Returns the error code if it fails.
func (a *Acheron) DirectSyscall(fnHash uint64, args ...uintptr) error {
	sys, err := a.resolver.GetSyscall(fnHash)
	if err != nil {
		return err
	}
	if errCode := execDirectSyscall(sys.SSN, args...); errCode != 0 { // !NT_SUCCESS
		return errors.New(fmt.Sprintf("syscall failed with error code %d", errCode))
	}
	return nil
}
