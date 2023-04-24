// go:build windows

package inject

var (
	baseAddr uintptr
	hThread  uintptr
	hSelf    = uintptr(0xffffffffffffffff) // handle to current proc
)

const (
	nullptr = uintptr(0)
)

func execDirectSyscall(callID uint16, argh ...uintptr) (errcode uint32)
