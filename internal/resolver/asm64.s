#include "textflag.h"

// func getNtdllBaseAddr() uintptr
TEXT ·getNtdllBaseAddr(SB),NOSPLIT,$0

    // TEB->ProcessEnvironmentBlock
    XORQ AX, AX
    MOVQ 0x30(GS), AX
    MOVQ 0x60(AX), AX

    // PEB->Ldr
    MOVQ 0x18(AX), AX

    // PEB->Ldr->InMemoryOrderModuleList
    MOVQ 0x20(AX), AX

    // PEB->Ldr->InMemoryOrderModuleList->Flink (ntdll.dll)
    MOVQ (AX), AX

    // PEB->Ldr->InMemoryOrderModuleList->Flink DllBase
    MOVQ 0x20(AX), AX

    MOVQ AX, ret+0(FP)
    RET


// func getModuleExportsDirAddr (moduleBase uintptr) uintptr
TEXT ·getModuleExportsDirAddr(SB),NOSPLIT,$0-8
    MOVQ moduleBase+0(FP), AX

    XORQ R15, R15
    XORQ R14, R14

    // AX = IMAGE_DOS_HEADER->e_lfanew offset
    MOVB 0x3C(AX), R15

    // R15 = ntdll base + R15
    ADDQ AX, R15

    // R15 = R15 + OptionalHeader + DataDirectory offset
    ADDQ $0x88, R15

    // AX = ntdll base + IMAGE_DATA_DIRECTORY.VirtualAddress
    ADDL (R15), R14
    ADDQ R14, AX

    MOVQ AX, ret+8(FP)
    RET


// func getExportsNumberOfNames(exportsBase uintptr) uint32
TEXT ·getExportsNumberOfNames(SB),NOSPLIT,$0-8
    MOVQ exportsBase+0(FP), AX

    XORQ R15, R15

    // R15 = exportsBase + IMAGE_EXPORT_DIRECTORY.NumberOfNames
    MOVL 0x18(AX), R15

    MOVL R15, ret+8(FP)
    RET


// func getExportsAddressOfFunctions(moduleBase,exportsBase uintptr) uintptr
TEXT ·getExportsAddressOfFunctions(SB),NOSPLIT,$0-16
    MOVQ moduleBase+0(FP), AX
    MOVQ exportsBase+8(FP), R8

    XORQ SI, SI

    // R15 = exportsBase + IMAGE_EXPORT_DIRECTORY.AddressOfFunctions
    MOVL 0x1c(R8), SI

    // AX = exportsBase + AddressOfFunctions offset
    ADDQ SI, AX

    MOVQ AX, ret+16(FP)
    RET


// func getExportsAddressOfNames(moduleBase,exportsBase uintptr) uintptr
TEXT ·getExportsAddressOfNames(SB),NOSPLIT,$0-16
    MOVQ moduleBase+0(FP), AX
    MOVQ exportsBase+8(FP), R8

    XORQ SI, SI

    // SI = exportsBase + IMAGE_EXPORT_DIRECTORY.AddressOfNames
    MOVL 0x20(R8), SI

    // AX = exportsBase + AddressOfNames offset
    ADDQ SI, AX

    MOVQ AX, ret+16(FP)
    RET


// func getExportsAddressOfNameOrdinals(moduleBase, exportsBase uintptr) uintptr
TEXT ·getExportsAddressOfNameOrdinals(SB),NOSPLIT,$0-16
    MOVQ moduleBase+0(FP), AX
    MOVQ exportsBase+8(FP), R8

    XORQ SI, SI

    // SI = exportsBase + IMAGE_EXPORT_DIRECTORY.AddressOfNameOrdinals
    MOVL 0x24(R8), SI

    // AX = exportsBase + AddressOfNames offset
    ADDQ SI, AX

    MOVQ AX, ret+16(FP)
    RET

