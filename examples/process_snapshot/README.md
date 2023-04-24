# Process Snapshot

Simple example that shows how to call `NtQuerySystemInformation` to get a snapshot of all running processes using indirect syscalls.

Compile with:

```bash
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o process_snapshot.exe
```

Output:

```
PS C:\> .\process_snapshot.exe
PID:  4    Name:  System
PID:  72   Name:  Registry
PID:  312  Name:  smss.exe
PID:  408  Name:  csrss.exe
PID:  476  Name:  wininit.exe
PID:  484  Name:  csrss.exe
PID:  544  Name:  winlogon.exe
PID:  568  Name:  services.exe
...
```
