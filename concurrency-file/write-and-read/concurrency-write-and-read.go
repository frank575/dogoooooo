package main

import (
	"bufio"
	concurrency_file "dogoooooo/concurrency-file"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func checkOpen(err error) {
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
}

func checkRead(err error) {
	if err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
}

func readFile(path string, wg *sync.WaitGroup) {
	file, err := os.Open(path)
	checkOpen(err)
	defer file.Close()

	r := bufio.NewScanner(file)
	var num int

	for r.Scan() {
		text := r.Text()
		fmt.Printf("path: %s, txt: %s\n", path, text)
		intTxt, _ := strconv.Atoi(text)
		num = intTxt + 1
		go writeFile(path, num, &*wg)
	}
	checkRead(r.Err())
}

func writeFile(path string, num int, wg *sync.WaitGroup) {
	file, err := os.Create(path)
	checkOpen(err)
	defer func() {
		file.Close()
		wg.Done()
	}()

	w := bufio.NewWriter(file)
	w.WriteString(strconv.Itoa(num))
	w.Flush()
}

func main() {
	fileNameList := concurrency_file.CreatePathList()
	var wg sync.WaitGroup
	for _, path := range fileNameList.GetRandFilePathList(3) {
		wg.Add(1)
		go readFile(path, &wg)
	}
	wg.Wait()
	fmt.Println("Done!")
}
