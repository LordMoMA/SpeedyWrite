package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func writeToFile2() {
	f, err := os.Create("./files/goFile2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	for i := 0; i < 90000; i++ {
		_, err := w.WriteString(fmt.Sprintf("This is line %d\n", i))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Ensure all data has been written to the underlying writer
	if err := w.Flush(); err != nil {
		fmt.Println(err)
		return
	}
}

func improveWrite() {
	start := time.Now()

	writeToFile2()

	elapsed := time.Since(start)
	fmt.Printf("Time taken for the improved version: %s\n", elapsed)
}
