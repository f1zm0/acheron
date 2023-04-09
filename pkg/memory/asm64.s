#include "textflag.h"

// func RVA2VA(moduleBase uintptr, rva uint32) uintptr
TEXT ·RVA2VA(SB),NOSPLIT,$0-16
    MOVQ moduleBase+0(FP), AX
    XORQ DI, DI

    MOVL rva+8(FP), DI
    ADDQ DI, AX

    MOVQ AX, ret+16(FP)
    RET


// func ReadDwordAt(start uintptr, offset uint32) uint32
TEXT ·ReadDwordAt(SB),NOSPLIT,$0-16
    MOVQ start+0(FP), AX
    MOVL offset+8(FP), R8

    XORQ DI, DI
    ADDQ R8, AX
    MOVL (AX), DI

    MOVL DI, ret+16(FP)
    RET


// func ReadWordAt(start uintptr, offset uint32) uint16
TEXT ·ReadWordAt(SB),NOSPLIT,$0-16
    MOVQ start+0(FP), AX
    MOVL offset+8(FP), R8

    XORQ DI, DI
    ADDQ R8, AX
    MOVW (AX), DI

    MOVW DI, ret+16(FP)
    RET


// func ReadByteAt(start uintptr, offset uint32) uint8
TEXT ·ReadByteAt(SB),NOSPLIT,$0-16
    MOVQ start+0(FP), AX
    MOVL offset+8(FP), R8

    XORQ DI, DI
    ADDQ R8, AX
    MOVB (AX), DI

    MOVB DI, ret+16(FP)
    RET
