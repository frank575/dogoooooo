package main

import (
	"bufio"
	"dogoooooo/file/util"
	"fmt"
	"os"
)

/*
# 狗語言練習範例

## 目錄結構

- [file](./file)
    - [util](./util) 公共方法
    - [concurrency-file](./file/concurrency-file) 併發檔案處理
        - [files]()
*/
func main() {
	f, err := os.Create("README.md")
	util.CheckOpen(err)

	w := bufio.NewWriter(f)
	w.WriteString(fmt.Sprintf("# 狗語言練習範例\n" +
		"```command\n" +
		"# 生成 README.md 指令\n" +
		"go run file/create-readme/create-readme.go\n" +
		"```\n" +
		"## 目錄結構\n" +
		"\n" +
		"- [file](./file)\n"))
	w.Flush()
}
