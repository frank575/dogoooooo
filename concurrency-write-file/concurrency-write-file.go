package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

func checkOpen(err error) {
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
}

func writeFile(path string, wg *sync.WaitGroup) {
	file, err := os.Create(path)
	checkOpen(err)
	defer func() {
		file.Close()
		wg.Done()
	}()

	duration := time.Millisecond * time.Duration(rand.Intn(3000))
	time.Sleep(duration)
	fmt.Printf("path: %s, duration: %d\n", path, duration)

	writer := bufio.NewWriter(file)
	writer.WriteString("1")
	writer.Flush()
}

func main() {
	fileNameList := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	var pathList []string

	for _, name := range fileNameList {
		pathList = append(pathList, fmt.Sprintf("concurrency-write-file/files/%s.txt", name))
	}

	var wg sync.WaitGroup

	for _, path := range pathList {
		wg.Add(1)
		go writeFile(path, &wg)
	}

	wg.Wait()
}
