// toc#file:::基本讀取檔案#toc
package main

import (
	"bufio"
	"dogoooooo/basic/file/util"
	"fmt"
	"os"
)

// 如何高效讀取檔案
// https://www.delftstack.com/zh-tw/howto/go/how-to-read-a-file-line-by-line-in-go/

// bufio
// https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter01/01.4.html
func main() {
	path := "basic/file/basic/read-file/a.txt"
	file, err := os.Open(path)
	defer file.Close()

	util.CheckOpen(err)
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	util.CheckOpen(fileScanner.Err())
}
