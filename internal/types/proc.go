package types

import (
	"bytes"
	"encoding/binary"
	"io"

	rrd "github.com/f1zm0/acheron/pkg/rawreader"
)

// InMemProc is a struct that contains the name, base address and SSN of a function.
type InMemProc struct {
	Name     string
	BaseAddr uintptr
	GateAddr uintptr
	SSN      int
}

func (p *InMemProc) IsHooked() bool {
	safeBytes := []byte{0x4c, 0x8b, 0xd1, 0xb8}
	stub := make([]byte, len(safeBytes))

	rr := rrd.NewRawReader(p.BaseAddr, len(safeBytes))

	sr := io.NewSectionReader(rr, 0, 1<<63-1)
	binary.Read(sr, binary.LittleEndian, &stub)

	if bytes.Compare(stub, safeBytes) == 0 {
		return true
	}
	return false
}
