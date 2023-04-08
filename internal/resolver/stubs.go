package resolver

// getNtdllBase signature.
func getNtdllBaseAddr() uintptr

// getModuleEATAddr signature.
func getModuleEATAddr(modBaseAddr uintptr) uintptr

// getEATNumberOfFunctions signature.
func getEATNumberOfFunctions(exportsBase uintptr) uint32

// getEATAddressOfFunctions signature.
func getEATAddressOfFunctions(moduleBase, exportsBase uintptr) uintptr

// getEATAddressOfNames signature.
func getEATAddressOfNames(moduleBase, exportsBase uintptr) uintptr
