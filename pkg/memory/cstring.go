package memory

// ReadCStringAt reads a null-terminated ANSI string from memory.
func ReadCStringAt(start uintptr, offset uint32) []byte {
	var buf []byte
	for {
		ch := ReadByteAtOffset(start, offset)
		if ch == 0 {
			break
		}
		buf = append(buf, ch)
		offset++
	}
	return buf
}
