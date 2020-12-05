// toc#file:::併發寫入自訂檔案數的隨機檔案值(+1)#toc
package main

import (
	"bufio"
	"dogoooooo/basic/file/util"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		intTxt, _ := strconv.Atoi(text)
		num = intTxt + 1
		fmt.Printf("path: %s, txt: %s -> %d\n", path, text, num)
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

func applicationDo() {
	fileNameList := util.CreatePathList()
	var wg sync.WaitGroup
	for _, path := range fileNameList.GetRandFilePathList(3) {
		wg.Add(1)
		go readFile(path, &wg)
	}
	wg.Wait()
	fmt.Println("Done!")

	answer := ""
	fmt.Println("Please enter - reset? [Y/N (default)]")
	fmt.Scanln(&answer)

	if strings.ToUpper(answer) == "Y" {
		applicationDo()
	}
}

func main() {
	applicationDo()
}
