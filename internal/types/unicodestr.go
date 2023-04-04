package types

import "golang.org/x/sys/windows"

// UnicodeString is a struct that represents a Windows Unicode string.
type UnicodeString struct {
	Length        uint16
	MaximumLength uint16
	Buffer        *uint16
}

// String returns the string representation of the UnicodeString.
func (s UnicodeString) String() string {
	return windows.UTF16PtrToString(s.Buffer)
}
