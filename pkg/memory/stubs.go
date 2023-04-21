package memory

func RVA2VA(moduleBase uintptr, rva uint32) uintptr

func ReadDwordAtOffset(start uintptr, offset uint32) uint32

func ReadWordAtOffset(start uintptr, offset uint32) uint16

func ReadByteAtOffset(start uintptr, offset uint32) byte
