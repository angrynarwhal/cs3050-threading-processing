import time
import threading
import multiprocessing

# Function for CPU-heavy task
def cpu_heavy_task(id):
    total = 0
    for i in range(10000000):
        total += i * i
    return total

# Function to run tasks with threading in Python
def thread_example():
    print("Starting threads (Python)")
    threads = []
    start_time = time.time()
    for i in range(5):
        thread = threading.Thread(target=cpu_heavy_task, args=(i,))
        threads.append(thread)
        thread.start()

    for thread in threads:
        thread.join()
    
    print(f"Threads completed in: {time.time() - start_time:.4f} seconds")

# Function to run tasks with multiprocessing in Python
def process_example():
    print("Starting processes (Python)")
    processes = []
    start_time = time.time()
    for i in range(5):
        process = multiprocessing.Process(target=cpu_heavy_task, args=(i,))
        processes.append(process)
        process.start()

    for process in processes:
        process.join()
    
    print(f"Processes completed in: {time.time() - start_time:.4f} seconds")

# Go code simulation explanation (Step by step)
def explain_python_vs_go():
    print("\nStep 1: Threading in Python")
    thread_example()
    print("\nStep 2: Multiprocessing in Python")
    process_example()
    print("\nStep 3: Explanation of Python's GIL and Threading")
    print("Python's Global Interpreter Lock (GIL) prevents true parallel execution of threads.")
    print("This means threads cannot fully utilize multiple cores for CPU-bound tasks.")
    print("While threads are useful for I/O-bound tasks (like networking), they are not ideal for CPU-bound tasks.")
    
    print("\nStep 4: Multiprocessing in Python")
    print("Multiprocessing avoids the GIL and utilizes multiple processes, each with its own memory space.")
    print("This is useful for CPU-bound tasks because it allows each process to run on a separate core.")
    print("However, multiprocessing comes with more overhead, such as memory duplication and inter-process communication.")
    
    print("\nStep 5: How Go Handles Concurrency")
    print("In Go, concurrency is handled using Goroutines, which are more lightweight than Python threads.")
    print("Go's runtime schedules Goroutines onto available threads, and there is no GIL that prevents concurrent tasks.")
    print("Goroutines share memory safely through channels, and Go can scale with multiple processes without the overhead of multiprocessing.")
    
    print("\nStep 6: How Python and Go Differ")
    print("Python has a Global Interpreter Lock (GIL), limiting true parallelism in multi-threaded programs.")
    print("In contrast, Go can handle true concurrency with Goroutines and avoids issues related to the GIL.")
    print("Python is better for I/O-bound tasks using threads, while Go excels at parallel computation with Goroutines.")

# Main function to start the explanation
def main():
    explain_python_vs_go()

if __name__ == "__main__":
    main()