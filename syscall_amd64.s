#include "textflag.h"

#define maxargs 16

// syscall implementation is a mix of:
// - https://golang.org/src/runtime/sys_windows_amd64.s
// - https://github.com/C-Sto/BananaPhone/blob/master/pkg/BananaPhone/asm_x64.s#L96
// with custom modifications to support indirect syscall execution 
// via a trampoline (syscall;ret instruction) in ntdll.dll

// func execIndirectSyscall(ssn uint16, trampoline uintptr, argh ...uintptr) (errcode uint32)
TEXT ·execIndirectSyscall(SB),NOSPLIT, $0-32
    XORQ    AX, AX
    MOVW    ssn+0(FP), AX
	
    XORQ    R11, R11
    MOVQ    trampoline+8(FP), R11
	
    //put variadic pointer into SI
    MOVQ    argh_base+16(FP),SI

    //put variadic size into CX
    MOVQ    argh_len+24(FP),CX
	
    // SetLastError(0).
    MOVQ    0x30(GS), DI
    MOVL    $0, 0x68(DI)

    // room for args
    SUBQ    $(maxargs*8), SP	
	
    PUSHQ   CX

    //no parameters, special case
    CMPL    CX, $0
    JLE     jumpcall
	
    // Fast version, do not store args on the stack.
    CMPL    CX, $4
    JLE	    loadregs

    // Check we have enough room for args.
    CMPL    CX, $maxargs
    JLE	    2(PC)

    // not enough room -> crash
    INT	    $3			

    // Copy args to the stack.
    MOVQ    SP, DI
    CLD
    REP; MOVSQ
    MOVQ    SP, SI
	
loadregs:

    // Load first 4 args into correspondent registers.
    MOVQ	0(SI), CX
    MOVQ	8(SI), DX
    MOVQ	16(SI), R8
    MOVQ	24(SI), R9
	
    // Floating point arguments are passed in the XMM registers
    // Set them here in case any of the arguments are floating point values. 
    // For details see: https://msdn.microsoft.com/en-us/library/zthk2dkh.aspx
    MOVQ	CX, X0
    MOVQ	DX, X1
    MOVQ	R8, X2
    MOVQ	R9, X3
	
jumpcall:
    MOVQ    CX, R10

    //jump to syscall;ret gadget address instead of direct syscall
    CALL    R11

    ADDQ	$((maxargs)*8), SP

    // Return result
    POPQ	CX
    MOVL	AX, errcode+40(FP)
    RET
