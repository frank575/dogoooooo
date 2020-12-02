package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

// 如何高效讀取檔案
// https://www.delftstack.com/zh-tw/howto/go/how-to-read-a-file-line-by-line-in-go/

// bufio
// https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter01/01.4.html
func main() {
	path := "read-file/a.txt"
	file, err := os.Open(path)
	defer file.Close()

	checkOpen(err)
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	checkRead(fileScanner.Err())
}
