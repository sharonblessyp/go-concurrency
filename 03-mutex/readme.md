**WaitGroup** coordinated completion.
**Mutex** protects shared data.

-------------------------
**The Problem: Race Condition**
when multiple threads or processes access shared data concurrently, and the final outcome depends on the unpredictable order of execution. Without proper synchronization (e.g., locks) leads to data corruption, crashes, or security vulnerabilities

**eg:**
G1 reads counter = 5
G2 reads counter = 5
G1 writes 6
G2 writes 6

lost one increment.
That’s a data race.
---------------------------
Detecting Race Conditions
Go has a built-in tool to tell you if you forgot a Mutex. When running your code or tests, use the -race flag:

Bash
go run -race main.go

---------------------------
Goroutines → share memory
Shared memory → needs protection
Protection → Mutex
---------------------------
A Mutex has two primary methods:
- Lock()
- Unlock()