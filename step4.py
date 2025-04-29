import time
import matplotlib.pyplot as plt
import concurrent.futures
import multiprocessing
import threading

def cpu_heavy(x):
    start = time.time()
    total = 0
    for i in range(10_000_000):
        total += i * i
    duration = time.time() - start
    return x, duration

def run_parallel_tasks(mode="threading", num_tasks=8):
    times = []

    if mode == "threading":
        executor_class = concurrent.futures.ThreadPoolExecutor
    elif mode == "multiprocessing":
        executor_class = concurrent.futures.ProcessPoolExecutor
    else:
        raise ValueError("mode must be 'threading' or 'multiprocessing'")

    with executor_class(max_workers=num_tasks) as executor:
        futures = [executor.submit(cpu_heavy, i) for i in range(num_tasks)]
        for future in concurrent.futures.as_completed(futures):
            idx, duration = future.result()
            times.append((idx, duration))

    times.sort()  # ensure consistent order
    return [t[1] for t in times]

def visualize_comparison():
    tasks = 20
    threading_times = run_parallel_tasks("threading", num_tasks=tasks)
    multiprocessing_times = run_parallel_tasks("multiprocessing", num_tasks=tasks)

    plt.figure(figsize=(10, 6))
    plt.plot(threading_times, label="Multithreading", marker='o')
    plt.plot(multiprocessing_times, label="Multiprocessing", marker='s')
    plt.title("Task Duration: Threads vs Processes")
    plt.xlabel("Task Index")
    plt.ylabel("Time (seconds)")
    plt.legend()
    plt.grid(True)
    plt.tight_layout()
    plt.savefig("threads_vs_processes.png")
    plt.show()

if __name__ == "__main__":
    multiprocessing.set_start_method("spawn", force=True)
    visualize_comparison()