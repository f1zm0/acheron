// func GetInMemoryOrderModuleListPtr() uintptr
TEXT ·GetInMemoryOrderModuleListPtr(SB),$0-8
	// PEB
    MOVQ 0x60(GS), AX
	// PEB->Ldr
	MOVQ 0x18(AX), AX
	// PEB->Ldr->InMemoryOrderModuleList
	MOVQ 0x20(AX), AX
	MOVQ AX, ret+0(FP)
	RET

// func GetLdrTableEntryPtr(listptr uintptr, i int64) *LdrDataTableEntry
TEXT ·GetLdrTableEntryPtr(SB),$0-24

	MOVQ listptr+0(FP), AX

	XORQ R10, R10
next_entry:
	CMPQ R10, i+8(FP)
	JE endloop

	// next Flink
	MOVQ (AX), AX
	INCQ R10
	JMP next_entry

endloop:
	MOVQ AX, CX
	// start of LDR_DATA_TABLE_ENTRY struct
	SUBQ $0x10, CX
	MOVQ CX, ret+16(FP)
	RET
