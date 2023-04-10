package resolver

type Resolver interface {
	// GetSyscallSSN returns the syscall SSN.
	GetSyscall(funcNameHash int64) (*Syscall, error)
}
