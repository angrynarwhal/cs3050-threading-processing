import time

def task(name):
    print(f"Starting task {name}")
    time.sleep(2)
    print(f"Task {name} completed")

if __name__ == "__main__":
    start = time.time()

    task("A")
    task("B")

    end = time.time()
    print(f"Total time: {end - start:.2f} seconds")