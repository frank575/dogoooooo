// toc#file:::(舊 - 廢棄)生成 README 目錄簡述#toc
package main

import (
	"bufio"
	"dogoooooo/basic/file/util"
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

func createTOCInfoList(ignoreList *[]string, path string, extension string) []FileInfo {
	//f, err := os.Getwd()
	//util.CheckGetwd(err)

	dirList, err := ioutil.ReadDir(path)
	util.CheckReadDir(err)

	var children []FileInfo

	for i := range dirList {
		file := dirList[i]
		fileName := file.Name()
		if file.IsDir() && !util.CheckIgnore(&*ignoreList, &fileName) {
			path := fmt.Sprintf("%s/%s", path, fileName)
			info := ""

			if mainFile, err := os.Open(path + "/" + "main." + extension); err == nil {
				r := bufio.NewReader(mainFile)
				line, _, err := r.ReadLine()
				util.CheckRead(err)

				reg, _ := regexp.Compile("//\\s.+")
				strLine := string(line)
				if reg.MatchString(strLine) {
					info = strLine[3:]
				}

				mainFile.Close()
			}

			children = append(children, FileInfo{fileName, info, createTOCInfoList(ignoreList, path, extension)})
			//fmt.Println(path, fileName)
		}
	}

	return children
}

func writeTOCList(tab string, path string, strTOC *string, fileInfoList *[]FileInfo) {
	for _, info := range *fileInfoList {
		name := info.name
		src := path + "/" + name
		description := info.info
		children := info.children
		if description == "" {
			if len(children) > 0 {
				*strTOC += fmt.Sprintf("%s- [%s](%s) %s\n", tab, name, src, description)
			}
		} else {
			*strTOC += fmt.Sprintf("%s- [%s](%s) %s\n", tab, name, src, description)
		}
		writeTOCList(tab+"  ", src, strTOC, &children)
	}
}

func main() {
	before, after := util.GetReadmeText()
	ignoreList := util.GetGitIgnoreFile()

	extension := "go"
	if len(os.Args) > 1 {
		extension = os.Args[1]
	}
	var fileInfoList []FileInfo
	fileInfoList = createTOCInfoList(&ignoreList, ".", extension)
	strTOC := ""
	writeTOCList("", ".", &strTOC, &fileInfoList)

	util.WriteReadme(before, after, strTOC)
}
