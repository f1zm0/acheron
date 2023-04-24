# sc_inject

Extremely simple shellcode injector PoC, that injects calc shellcode using syscalls for `NtAllocateVirtualMemory`+`NtWriteVirtualMemory`+`NtCreateThreadEx`.

Using build tags, you can compile both the direct and indirect syscall versions of the injector, if you want to run them against defensive tools to check out the detection and compare the IOCs of each technique.

```bash
# indirect syscall version (default)
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o sc_inject_indirect.exe

# direct syscall version
GOOS=windows GOARCH=amd64 go build -tags='direct' -ldflags "-s -w" -o sc_inject_direct.exe
```
