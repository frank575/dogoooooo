package main

import (
	"bufio"
	concurrency_file "dogoooooo/concurrency-file"
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

func writeFile(path string, ch *chan string) {
	file, err := os.Create(path)
	checkOpen(err)
	defer func() {
		file.Close()
		*ch <- path
	}()

	duration := time.Millisecond * time.Duration(rand.Intn(3000))
	time.Sleep(duration)
	fmt.Printf("path: %s, duration: %d\n", path, duration)

	writer := bufio.NewWriter(file)
	writer.WriteString("1")
	writer.Flush()
}

func checkChannel(ch *chan string, i *int, size int) {
	*i++
	if *i == size {
		close(*ch)
	}
}

func main() {
	fileNameList := concurrency_file.CreatePathList()

	ch := make(chan string, len(fileNameList.List))

	for _, path := range fileNameList.GetFilePathList() {
		go writeFile(path, &ch)
	}

	i := 0
	for path := range ch {
		checkChannel(&ch, &i, len(fileNameList.List))
		fmt.Printf("%s done!\n", path)
	}
}
