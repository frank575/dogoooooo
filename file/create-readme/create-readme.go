package main

import (
	"bufio"
	"dogoooooo/file/util"
	"fmt"
	"io/ioutil"
	"os"
)

type Info struct {
	path string
	info string
}

func getIgnoreFile() []string {
	f, err := os.Open(".gitignore")
	util.CheckOpen(err)
	r := bufio.NewScanner(f)

	var ignoreList []string

	for r.Scan() {
		text := r.Text()
		ignoreList = append(ignoreList, text)
	}
	util.CheckRead(r.Err())

	return ignoreList
}

func checkIgnore(ignoreList *[]string, fileName *string) bool {
	isReplace := false
	for _, ignoreName := range *ignoreList {
		if *fileName == ignoreName {
			isReplace = true
			break
		}
	}
	return isReplace
}

func createTOC(ignoreList *[]string, title *string, path string, tab string, infoList *[]Info) {
	//f, err := os.Getwd()
	//util.CheckGetwd(err)

	dirList, err := ioutil.ReadDir(path)
	util.CheckReadDir(err)

	info := ""

	for i := range dirList {
		file := dirList[i]
		fileName := file.Name()
		if fileName == "info.txt" {
			*infoList = append(*infoList, Info{path, "2123132132"})
		}
		if file.IsDir() && !checkIgnore(&*ignoreList, &fileName) {
			path := fmt.Sprintf("%s/%s", path, fileName)
			*title += fmt.Sprintf("%s- [%s](%s) %s\n", tab, fileName, path, info)
			createTOC(&*ignoreList, &*title, path, tab+"  ", &*infoList)
			//fmt.Println(path, fileName)
		}
	}
}

func main() {
	f, err := os.Create("README.md")
	util.CheckOpen(err)

	w := bufio.NewWriter(f)
	title := "# 狗語言練習範例\n" +
		"```command\n" +
		"# 生成 README.md 指令(開發中)\n" +
		"go run file/create-readme/create-readme.go\n" +
		"```\n" +
		"## 目錄結構\n" +
		"\n"

	ignoreList := getIgnoreFile()
	var infoList []Info
	createTOC(&ignoreList, &title, ".", "", &infoList)

	fmt.Println(infoList)

	w.WriteString(fmt.Sprintf(title))
	w.Flush()
}
