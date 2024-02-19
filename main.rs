use std::fs::OpenOptions;
use std::io::prelude::*;
use std::sync::{Arc, Mutex};
use std::thread;

fn main() {
    let file = Arc::new(Mutex::new(
        OpenOptions::new()
            .write(true)
            .create(true)
            .append(true)
            .open("rustFile.txt")
            .unwrap(),
    ));

    let mut handles = vec![];

    for i in 0..1000 {
        let file = Arc::clone(&file);
        let handle = thread::spawn(move || {
            let mut file = file.lock().unwrap();
            writeln!(file, "This is line {}", i).unwrap();
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }
}