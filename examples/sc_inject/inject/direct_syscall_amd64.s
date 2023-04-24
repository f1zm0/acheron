#define maxargs 16

// func execDirectSyscall(callID uint16, argh ...uintptr) (errcode uint32)
TEXT ·execDirectSyscall(SB), $0-56
	XORQ AX,AX
	MOVW callid+0(FP), AX
	PUSHQ CX

	//put variadic pointer into SI
	MOVQ argh_base+8(FP),SI

	//put variadic size into CX
	MOVQ argh_len+16(FP),CX

	// SetLastError(0).
	MOVQ	0x30(GS), DI
	MOVL	$0, 0x68(DI)
	SUBQ	$(maxargs*8), SP	// room for args

	// Fast version, do not store args on the stack.
	CMPL	CX, $4
	JLE	loadregs

	// Check we have enough room for args.
	CMPL	CX, $maxargs
	JLE	2(PC)
	INT	$3			// not enough room -> crash

	// Copy args to the stack.
	MOVQ	SP, DI
	CLD
	REP; MOVSQ
	MOVQ	SP, SI

	//move the stack pointer????? why????
	SUBQ	$8, SP

loadregs:
	// Load first 4 args into correspondent registers.
	MOVQ	0(SI), CX
	MOVQ	8(SI), DX
	MOVQ	16(SI), R8
	MOVQ	24(SI), R9

	// Floating point arguments are passed in the XMM
	// registers. Set them here in case any of the arguments
	// are floating point values. For details see
	//	https://msdn.microsoft.com/en-us/library/zthk2dkh.aspx
	MOVQ	CX, X0
	MOVQ	DX, X1
	MOVQ	R8, X2
	MOVQ	R9, X3

	MOVQ    CX, R10
	SYSCALL
	ADDQ	$((maxargs+1)*8), SP

	POPQ	CX
	MOVL	AX, errcode+32(FP)
	RET
