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
Time taken for the improved version: 8.970291ms
```

## How to run the Rust file

```bash
➜  SpeedyWrite git:(main) ✗ cargo run main.rs
   Compiling speedy_write v0.1.0 
    Finished dev [unoptimized + debuginfo] target(s) in 0.18s
     Running `target/debug/speedy_write main.rs`
Time elapsed in function is: 5.450542ms
```
