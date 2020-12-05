// toc#file:::併發(waitGroup 版本)寫入檔案#toc
package main

import (
	"bufio"
	"dogoooooo/basic/file/util"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func writeFile(path string, wg *sync.WaitGroup) {
	file, err := os.Create(path)
	util.CheckOpen(err)
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
	fileNameList := util.CreatePathList()
	var wg sync.WaitGroup

	for _, path := range fileNameList.GetFilePathList() {
		wg.Add(1)
		go writeFile(path, &wg)
	}

	wg.Wait()
}
