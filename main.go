package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex = &sync.Mutex{}

func writeToFile(i int) {
	defer wg.Done()
	mutex.Lock()
	f, err := os.OpenFile("./files/goFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("This is line %d\n", i)); err != nil {
		fmt.Println(err)
	}
	mutex.Unlock()
}

func originalWrite() {
	start := time.Now()
	for i := 0; i < 90000; i++ {
		wg.Add(1)
		go writeToFile(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Time taken for the first version: %s\n", elapsed)
}

func main() {
	fmt.Println("Starting original write...")
	originalWrite()
	fmt.Println("Starting improved write...")
	improveWrite()
}

// ➜  writeFile git:(master) ✗ go build -o file
// ➜  writeFile git:(master) ✗ ./file
// Time taken for the first version: 2.40659975s
// Starting improved write...
// Time taken for the improved version: 9.366375ms

// run go run main.go twice, new lines are added to the original file
