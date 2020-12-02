// !!生成 README 目錄腳本
package main

import (
	"bufio"
	"dogoooooo/file/util"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type FileInfo struct {
	name     string
	info     string
	children []FileInfo
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

func createTOCInfoList(ignoreList *[]string, path string) []FileInfo {
	//f, err := os.Getwd()
	//util.CheckGetwd(err)

	dirList, err := ioutil.ReadDir(path)
	util.CheckReadDir(err)

	var children []FileInfo

	for i := range dirList {
		file := dirList[i]
		fileName := file.Name()
		if file.IsDir() && !checkIgnore(&*ignoreList, &fileName) {
			path := fmt.Sprintf("%s/%s", path, fileName)
			info := ""

			if mainFile, err := os.Open(path + "/" + "main.go"); err == nil {
				r := bufio.NewReader(mainFile)
				line, _, err := r.ReadLine()
				util.CheckRead(err)

				reg, _ := regexp.Compile("//\\s!!.+")
				strLine := string(line)
				if reg.MatchString(strLine) {
					info = strLine[5:]
				}

				mainFile.Close()
			}

			children = append(children, FileInfo{fileName, info, createTOCInfoList(&*ignoreList, path)})
			//fmt.Println(path, fileName)
		}
	}

	return children
}

func writeTOCList(tab string, path string, title *string, fileInfoList *[]FileInfo) {
	for _, info := range *fileInfoList {
		name := info.name
		src := path + "/" + name
		description := info.info
		children := info.children
		if description == "" {
			if len(children) > 0 {
				*title += fmt.Sprintf("%s- [%s](%s) %s\n", tab, name, src, description)
			}
		} else {
			*title += fmt.Sprintf("%s- [%s](%s) %s\n", tab, name, src, description)
		}
		writeTOCList(tab+"  ", src, &*title, &children)
	}
}

func main() {
	f, err := os.Create("README.md")
	util.CheckOpen(err)

	w := bufio.NewWriter(f)
	title := "# 狗語言練習範例\n" +
		"```command\n" +
		"# 生成 README.md 指令(開發中)\n" +
		"go run file/readme-generate/main.go\n" +
		"```\n" +
		"## 項目結構\n" +
		"\n"

	ignoreList := getIgnoreFile()
	var fileInfoList []FileInfo
	fileInfoList = createTOCInfoList(&ignoreList, ".")

	writeTOCList("", ".", &title, &fileInfoList)

	w.WriteString(fmt.Sprintf(title))
	w.Flush()
}
