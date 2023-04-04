package rawreader

import (
	"errors"
	"io"
	"reflect"
	"unsafe"
)

// RawReader struct and functions below are taken from:
// https://github.com/awgh/rawreader/blob/master/rawreader.go

// RawReader struct uses reflect to read data from underlying memory
type RawReader struct {
	sliceHeader *reflect.SliceHeader
	rawPtr      uintptr
	Data        []byte
	Length      int
}

// NewRawReader returns a reference to a new populated RawReader
func NewRawReader(start uintptr, length int) *RawReader {
	sh := &reflect.SliceHeader{
		Data: start,
		Len:  length,
		Cap:  length,
	}
	data := *(*[]byte)(unsafe.Pointer(sh))
	return &RawReader{sliceHeader: sh, rawPtr: start, Data: data, Length: length}
}

// ReadAt func reads a file with a seek offset
func (f *RawReader) ReadAt(p []byte, off int64) (n int, err error) {
	if off < 0 {
		return 0, errors.New("RawReader.ReadAt: negative offset")
	}
	reqLen := len(p)
	buffLen := int64(f.Length)
	if off >= buffLen {
		return 0, io.EOF
	}

	n = copy(p, f.Data[off:])
	if n < reqLen {
		err = io.EOF
	}
	return n, err
}
