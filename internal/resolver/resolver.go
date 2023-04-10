package resolver

type Resolver interface {
	// GetSyscallSSN returns the syscall SSN.
	GetSyscall(funcNameHash uint64) (*Syscall, error)
}
