// syscall implementation is a mix of:
// - https://golang.org/src/runtime/sys_windows_amd64.s
// - https://github.com/C-Sto/BananaPhone/blob/master/pkg/BananaPhone/asm_x64.s#L96
// with added support for indirect syscall inspired by:
// - https://github.com/thefLink/RecycledGate

// func execSyscall(ssn uint16, gateAddr uintptr, argh ...uintptr) (errcode uint32)
#define maxargs 16
TEXT ·execIndirectSyscall(SB), $0-56
	XORQ    AX,AX
	MOVW    ssn+0(FP), AX
	PUSHQ   CX

        XORQ    BX,BX
        MOVQ    gateAddr+8(FP),BX

	//put variadic pointer into SI
	MOVQ    argh_base+16(FP),SI
	//put variadic size into CX
	MOVQ    argh_len+24(FP),CX

	// SetLastError(0).
	MOVQ	0x30(GS), DI
	MOVL	$0, 0x68(DI)
	SUBQ	$(maxargs*8), SP	// room for args

	// Fast version, do not store args on the stack.
	CMPL	CX, $4
	JLE	    loadregs

	// Check we have enough room for args.
	CMPL	CX, $maxargs
	JLE	    2(PC)
	INT	    $3	// not enough room -> crash

	// Copy args to the stack.
	MOVQ	SP, DI
	CLD
	REP; MOVSQ
	MOVQ	SP, SI

	//move the stack pointer????? why????
	// SUBQ	$8, SP

loadregs:
	// Load first 4 args into correspondent registers.
	MOVQ	0(SI), CX
	MOVQ	8(SI), DX
	MOVQ	16(SI), R8
	MOVQ	24(SI), R9

	// Floating point arguments are passed in the XMM
	MOVQ	CX, X0
	MOVQ	DX, X1
	MOVQ	R8, X2
	MOVQ	R9, X3

	MOVQ    CX, R10

        // jump to gate instead of direct syscall
        CALL BX

	ADDQ	$((maxargs+1)*8), SP

	// Return result.
	POPQ	CX
	MOVL	AX, errcode+32(FP)
	RET
