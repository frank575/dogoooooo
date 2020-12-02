package main

import (
	"bufio"
	concurrency_file "dogoooooo/concurrency-file"
	"fmt"
	"log"
	"os"
	"strconv"
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

func readFile(path string) {
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
		// TODO go
		writeFile(path, num)
	}
	checkRead(r.Err())
}

func writeFile(path string, num int) {
	file, err := os.Create(path)
	checkOpen(err)
	defer file.Close()

	w := bufio.NewWriter(file)
	w.WriteString(strconv.Itoa(num))
	w.Flush()
}

func main() {
	fileNameList := concurrency_file.CreatePathList()
	for _, path := range fileNameList.GetRandFilePathList(3) {
		// TODO go
		readFile(path)
	}
}
