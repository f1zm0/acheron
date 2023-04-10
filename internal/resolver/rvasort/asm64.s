#include "textflag.h"

// getTrampoline check if the export has a clean syscall;ret gadget within its 32 bytes.
// Returns the trampoline address if clean, nullptr if not.

// func getTrampoline(stubAddr uintptr) uintptr
TEXT ·getTrampoline(SB),NOSPLIT,$0-8
    MOVQ stubAddr+0(FP), AX
    MOVQ AX, R10

    // stub_length-gadget_length bytes of the stub (32-3)
    ADDQ $29, AX

loop:
    XORQ DI, DI

    // check for 0x0f05c3 byte sequence
    MOVB $0x0f, DI
    CMPB DI, 0(AX)
    JNE nope

    MOVB $0x05, DI
    CMPB DI, 1(AX)
    JNE nope

    MOVB $0xc3, DI
    CMPB DI, 2(AX)
    JNE nope

    // if we are here, we found a clean syscall;ret gadget
    MOVQ AX, ret+8(FP)
    RET

nope:
    // if AX is equal to R10, we have reached the start of the stub
    // which means we could not find a clean syscall;ret gadget
    CMPQ AX, R10
    JE not_found

    DECQ AX
    JMP loop

not_found:
    // returning nullptr
    XORQ AX, AX
    MOVQ AX, ret+0(FP)
    RET
