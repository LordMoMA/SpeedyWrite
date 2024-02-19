use std::fs::OpenOptions;
use std::io::prelude::*;
use std::io::BufWriter;
use std::time::Instant;

fn main() {
    let start = Instant::now();

    let file = OpenOptions::new()
        .write(true)
        .create(true)
        .append(true)
        .open("./files/rustFile.txt")
        .unwrap();

    let mut writer = BufWriter::new(file);

    let num_lines = 90000;

    for i in 0..num_lines {
        writeln!(writer, "This is line {}", i).unwrap();
    }

    let duration = start.elapsed();
    println!("Time elapsed in function is: {:?}", duration);
}