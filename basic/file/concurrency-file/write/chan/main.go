// toc#file:::併發(channel 版本)寫入檔案#toc
package main

import (
	"bufio"
	"dogoooooo/basic/file/util"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func writeFile(path string, ch *chan string) {
	file, err := os.Create(path)
	util.CheckOpen(err)
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
	fileNameList := util.CreatePathList()

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
