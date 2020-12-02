package main

import (
	"bufio"
	"dogoooooo/file/util"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func readFile(path string, wg *sync.WaitGroup) {
	file, err := os.Open(path)
	util.CheckOpen(err)
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
	util.CheckRead(r.Err())
}

func writeFile(path string, num int, wg *sync.WaitGroup) {
	file, err := os.Create(path)
	util.CheckOpen(err)
	defer func() {
		file.Close()
		wg.Done()
	}()

	w := bufio.NewWriter(file)
	w.WriteString(strconv.Itoa(num))
	w.Flush()
}

func main() {
	fileNameList := util.CreatePathList()
	var wg sync.WaitGroup
	for _, path := range fileNameList.GetRandFilePathList(3) {
		wg.Add(1)
		go readFile(path, &wg)
	}
	wg.Wait()
	fmt.Println("Done!")
}
