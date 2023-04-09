package resolver

func getNtdllBaseAddr() uintptr

func getModuleExportsDirAddr(modBaseAddr uintptr) uintptr

func getExportsNumberOfNames(exportsBase uintptr) uint32

func getExportsAddressOfFunctions(moduleBase, exportsBase uintptr) uintptr

func getExportsAddressOfNames(moduleBase, exportsBase uintptr) uintptr

func getExportsAddressOfNameOrdinals(moduleBase, exportsBase uintptr) uintptr
