package acheron

// Ref: https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-erref/87fba13e-bf06-450e-83b1-9241dc81e781
const (
	ErrSyscallNotFound = ((3 << 30) | (1 << 29) | 1) // Severity: STATUS_SEVERITY_ERROR, Custom: 1, Code: 1
)

func NT_SUCCESS(x uint32) bool {
	return (x)&(1<<31) == 0
}
