package resolver

import (
	"fmt"

	wt "github.com/f1zm0/acheron/internal/types"
)

// GetNtdlloduleHandle returns a PEModule struct with information about in-memory
// ntdll.dll's module, or an error if for some reason an error occur while reading it.
func GetNtdllModuleHandle() (*wt.PEModule, error) {
	moduleBase := getNtdllBaseAddr()
	fmt.Printf("ntdll base: 0x%x\r\n\r\n", moduleBase)

	exportsBase := getModuleEATAddr(moduleBase)
	fmt.Printf("EAT base: 0x%x\r\n\r\n", exportsBase)

	numberOfFunctions := getEATNumberOfFunctions(exportsBase)
	fmt.Printf("Number of functions: %d\r\n\r\n", numberOfFunctions)

	addressOfFunctions := getEATAddressOfFunctions(moduleBase, exportsBase)
	fmt.Printf("Address of functions: 0x%x\r\n\r\n", addressOfFunctions)

	addressOfNames := getEATAddressOfNames(moduleBase, exportsBase)
	fmt.Printf("Address of names: 0x%x\r\n\r\n", addressOfNames)

	// rr := rrd.NewRawReader(modBaseAddr, modSize)

	// p, err := pe.NewFileFromMemory(rr)
	// if err != nil {
	// 	return nil, errors.New("error reading module from memory")
	// }

	// return &wt.PEModule{
	// 	BaseAddr: modBaseAddr,
	// 	File:     p,
	// }, nil

	return &wt.PEModule{}, nil
}
