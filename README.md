# cs3050-threading-processing

## Python

### 🧠 What is the GIL?

### GIL stands for Global Interpreter Lock.
It’s a mechanism used in the CPython implementation of Python to ensure that only one thread executes Python bytecode at a time, even on multi-core systems.

⸻

### 🔒 Why Does Python Have a GIL?
	•	CPython is not thread-safe by default.
	•	The GIL simplifies memory management in CPython by preventing concurrent access to Python objects.

⸻

### 🛑 How It Bottlenecks Threads

Even if you create multiple threads for CPU-bound tasks:
	•	Only one thread can execute Python code at any moment.
	•	Other threads are just waiting, switching in and out of the lock.
	•	You don’t get real parallelism—just concurrent context switching.

✅ I/O-bound tasks (like reading files, waiting for a network) benefit from threading.
❌ CPU-bound tasks (like computing primes) don’t benefit because of the GIL.


## Go

### Libraries in Go for Visualization: 
```go
go mod init runtimeviz
go get gonum.org/v1/plot/...
``` 



## Contrast between Go and Python

| Feature            | Python (CPython)                            | Go                                      |
|--------------------|----------------------------------------------|------------------------------------------|
| **Concurrency model** | Threads (limited by GIL)                   | Goroutines (green threads)               |
| **Parallelism**       | Requires `multiprocessing`                 | Achieved naturally via goroutines        |
| **Thread cost**       | OS threads (heavy)                         | Goroutines (lightweight, managed by Go runtime) |
| **Memory sharing**    | Shared, but needs locks or queues          | Shared by default, channels preferred    |
| **GIL**               | Yes (CPython only)                         | No GIL in Go                             |