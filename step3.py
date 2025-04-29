import multiprocessing
import time

def task(name):
    print(f"Starting task {name}")
    time.sleep(2)
    print(f"Task {name} completed")

if __name__ == "__main__":
    start = time.time()

    p1 = multiprocessing.Process(target=task, args=("A",))
    p2 = multiprocessing.Process(target=task, args=("B",))

    p1.start()
    p2.start()

    p1.join()
    p2.join()

    end = time.time()
    print(f"Total time: {end - start:.2f} seconds")