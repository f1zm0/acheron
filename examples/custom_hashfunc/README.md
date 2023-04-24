# Custom Hash

Acheron allows passing a custom hashing function to the constructor, so that it can be used to store and retrieve the syscall structs from their map for better OPSEC.

In this example the custom function XORes the string buffer with `0xdeadbeef` key, and runs the result into SHA1 hash function.

Compile with:

```bash
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o custom_hash.exe main.go
```
