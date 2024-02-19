## Experiment Results

The fastest way to write 90000 lines into a file would be to use a single thread with buffered writing. This is because creating and managing multiple threads can add overhead, especially for I/O-bound tasks like writing to a file. Buffered writing minimizes the number of write calls to the underlying writer, which can significantly improve performance for workloads that involve many small writes.

## How to run the Node file

```bash
➜  writeFile git:(master) ✗ node main.js
writeStream: 815.381ms
writeFile: 828.453ms
```

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
