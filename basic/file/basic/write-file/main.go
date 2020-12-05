// toc#file:::基本寫入檔案#toc
package main

import (
	"bufio"
	"dogoooooo/basic/file/util"
	"os"
)

// bufio
// https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter01/01.4.html

// 如何在 golang 讀取/寫入 檔案
// https://blog.iphpo.com/blog/2017/03/%E5%A6%82%E4%BD%95%E5%9C%A8-golang-%E8%AE%80%E5%8F%96/%E5%AF%AB%E5%85%A5%E6%AA%94%E6%A1%88/
func main() {
	path := "basic/file/basic/write-file/a.txt"
	file, err := os.Create(path)
	defer file.Close()

	util.CheckOpen(err)
	fileWriter := bufio.NewWriter(file)
	fileWriter.WriteString("hello world\n123")
	fileWriter.Flush() // 將文字寫入
}
