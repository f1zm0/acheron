package types

import (
	"github.com/Binject/debug/pe"
)

// PEModule is a struct that contains the base address of a PE module and a pointer to the PE file.
type PEModule struct {
	BaseAddr uintptr
	File     *pe.File
}
