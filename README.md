# cs3050-threading-processing

## 🧠 What is the GIL?

### GIL stands for Global Interpreter Lock.
It’s a mechanism used in the CPython implementation of Python to ensure that only one thread executes Python bytecode at a time, even on multi-core systems.

⸻

## 🔒 Why Does Python Have a GIL?
	•	CPython is not thread-safe by default.
	•	The GIL simplifies memory management in CPython by preventing concurrent access to Python objects.

⸻

## 🛑 How It Bottlenecks Threads

Even if you create multiple threads for CPU-bound tasks:
	•	Only one thread can execute Python code at any moment.
	•	Other threads are just waiting, switching in and out of the lock.
	•	You don’t get real parallelism—just concurrent context switching.

✅ I/O-bound tasks (like reading files, waiting for a network) benefit from threading.
❌ CPU-bound tasks (like computing primes) don’t benefit because of the GIL.