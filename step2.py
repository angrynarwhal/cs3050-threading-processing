import threading
import time

def task(name):
    print(f"Starting task {name}")
    time.sleep(2)
    print(f"Task {name} completed")

if __name__ == "__main__":
    start = time.time()

    t1 = threading.Thread(target=task, args=("A",))
    t2 = threading.Thread(target=task, args=("B",))

    t1.start()
    t2.start()

    t1.join()
    t2.join()

    end = time.time()
    print(f"Total time: {end - start:.2f} seconds")