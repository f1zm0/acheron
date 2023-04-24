# Process Snapshot

Simple example that shows how to call `NtQuerySystemInformation` to get a snapshot of all running processes using indirect syscalls.

Compile with:

```bash
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o process_snapshot.exe main.go
```
