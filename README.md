## Experiment Results

The fastest way to write 90000 lines into a file would be to use a single thread with buffered writing. This is because creating and managing multiple threads can add overhead, especially for I/O-bound tasks like writing to a file. Buffered writing minimizes the number of write calls to the underlying writer, which can significantly improve performance for workloads that involve many small writes.

## How to run the Node file

```bash
➜  writeFile git:(master) ✗ node main.js
writeStream: 815.381ms
writeFile: 828.453ms
```

## How to run the Python file

```bash
➜  SpeedyWrite git:(main) ✗ python main.py
Time taken for the python version: 12.826204299926758 milliseconds
```

Python's built-in file I/O functions are highly optimized. When you open a file for writing in Python using the open function with mode 'w', Python creates a buffered binary file object, which means it uses a buffer to batch write operations, reducing the number of system calls and thus increasing performance.

## How to run the Go file

```bash
➜  writeFile git:(master) ✗ go build -o file
➜  writeFile git:(master) ✗ ./file          
Starting original write...
Time taken for the first version: 2.455568375s
Starting improved write...
Time taken for the improved version: 7.690667ms
```

## How to run the Rust file

```bash
➜  SpeedyWrite git:(main) ✗ cargo run main.rs
   Compiling speedy_write v0.1.0 
    Finished dev [unoptimized + debuginfo] target(s) in 0.18s
     Running `target/debug/speedy_write main.rs`
Time elapsed in function is: 6.700958ms
```

## How to run the C file

```bash
➜  cFile git:(main) ✗ gcc main.c -o cFile
➜  cFile git:(main) ✗ ./cFile            
Time taken: 20.605000 milliseconds
```

The performance of a program can be influenced by many factors, including the efficiency of the code, the capabilities of the language, the performance of the compiler, and the specifics of the system it's running on.

In this case, the C code might be slower than the Rust and Go code due to the way it handles file I/O operations. The C standard library's file I/O functions, such as fputs, are generally less efficient than the file I/O libraries used by Rust and Go.

Rust and Go have more modern standard libraries that are designed with performance in mind. They use buffering and other techniques to minimize the number of system calls, which can significantly improve the speed of file I/O operations.

In addition, the Rust and Go compilers have more advanced optimization capabilities than the gcc compiler used for C. They can automatically apply a variety of optimizations to the code that can make it run faster.

## Why NodeJS is not even catching up with Python

1. Asynchronous Overhead: 

Node.js is inherently asynchronous and uses an event-driven architecture. While this is great for handling many operations concurrently, it can add some overhead due to the setup and teardown of callbacks and event loops, especially for I/O-bound tasks.

2. V8 Engine: 

Node.js runs on the V8 JavaScript engine, which, while highly optimized, might not be as fast as Python's CPython interpreter for certain tasks, especially those that are I/O-bound rather than CPU-bound.

3. Buffering: 

Python's file I/O operations are buffered by default, meaning they collect many small write operations into a larger one, reducing the number of system calls and thus increasing performance. While Node.js's fs.createWriteStream also uses buffering, the specifics of how it's implemented could potentially make it slower than Python's implementation.