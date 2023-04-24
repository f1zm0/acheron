# Inject

Extremely simple shellcode injector PoC, that injects calc shellcode using `NtAllocateVirtualMemory`+`NtWriteVirtualMemory`+`NtCreateThreadEx`.

Using build tags, you can compile both the direct and indirect syscall versions of the injector, if you want to run them against defensive tools and see the detection for each technique.

```bash
# indirect syscall version (default)
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o sc_inject_indirect.exe

# direct syscall version
GOOS=windows GOARCH=amd64 go build -tags='direct' -ldflags "-s -w" -o sc_inject_direct.exe
```
