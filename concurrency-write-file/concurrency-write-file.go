package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func checkOpen(err error) {
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
}

func writeFile(path string) {
	file, err := os.Create(path)
	checkOpen(err)
	defer file.Close()

	duration := time.Millisecond * time.Duration(rand.Intn(1000))
	fmt.Printf("path: %s, duration: %d", path, duration)
	time.Sleep(duration)

	writer := bufio.NewWriter(file)
	writer.WriteString("hello world!")
	writer.Flush()
}

func main() {
	fileNameList := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	var pathList []string

	for _, name := range fileNameList {
		pathList = append(pathList, "concurrency-write-file/files/"+name+".txt")
	}

	for _, path := range pathList {
		go writeFile(path)
	}
}
