import time

def write_to_file():
    with open('pythonFile.txt', 'w') as f:
        for i in range(90000):
            f.write(f'This is line {i}\n')

def calc_write_time():
    start = time.time()

    write_to_file()

    elapsed = time.time() - start
    print(f'Time taken for the python version: {elapsed * 1000} milliseconds')

calc_write_time()