package memory

func RVA2VA(moduleBase uintptr, rva uint32) uintptr

func ReadDwordAt(start uintptr, offset uint32) uint32

func ReadWordAt(start uintptr, offset uint32) uint16

func ReadByteAt(start uintptr, offset uint32) byte
