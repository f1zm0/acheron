#include "textflag.h"

// func XorDjb2Hash(data []byte) uint64
TEXT ·XorDjb2Hash(SB),NOSPLIT,$0-24
    MOVQ    data+0(FP), DI
    MOVQ    DI, R9
    MOVQ    len+8(FP), SI
    MOVQ    SI, R10

    // djb2 magic number
    XORQ    AX, AX
    MOVL    $5381, AX

    // xor byte
    XORQ    R15, R15
    MOVB    $0xf1, R15

    // if data slice is empty, return 0
    CMPQ    SI, $0
    JEQ     done

loop_xor:
    // data[i] = data[i] ^ 0xf1
    XORB    R15, (DI)

    INCQ    DI
    DECQ    SI
    CMPQ    SI, $0
    JNE     loop_xor

    // restore start ptr and length
    MOVQ    R9, DI
    MOVQ    R10, SI

loop_hash:
    // move nth byte to BX
    MOVB    (DI), BX
    INCQ    DI

    // hash = hash * 33 + data[i]
    IMULQ   $33, AX
    ADDQ    BX, AX

    DECQ    SI
    CMPQ    SI, $0
    JNE     loop_hash

done:
    MOVQ    AX, ret+24(FP)
    RET
