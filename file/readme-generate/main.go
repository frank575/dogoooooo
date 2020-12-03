// 生成 README 的項目目錄結構腳本
package main

import (
	"bufio"
	"dogoooooo/file/util"
	"errors"
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
	var ignoreList []string
	f, err := os.Open(".gitignore")
	//util.CheckOpen(err)
	defer f.Close()

	if err != nil {
	} else {
		r := bufio.NewScanner(f)
		for r.Scan() {
			text := r.Text()
			ignoreList = append(ignoreList, text)
		}
		util.CheckRead(r.Err())
	}

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

func createTOCInfoList(ignoreList *[]string, path string, extension string) []FileInfo {
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

			children = append(children, FileInfo{fileName, info, createTOCInfoList(&*ignoreList, path, extension)})
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
		writeTOCList(tab+"  ", src, &*strTOC, &children)
	}
}

func main() {
	readmeName := "README.md"
	rf, err := os.Open(readmeName)
	util.CheckRead(err)
	defer rf.Close()

	r := bufio.NewScanner(rf)
	readmeText := ""

	for r.Scan() {
		txt := r.Text()
		readmeText += txt + "\n"
	}

	regTOC, _ := regexp.Compile("<!--TOC-->")
	TOCIndex := regTOC.FindAllIndex([]byte(readmeText), -1)
	if len(TOCIndex) < 2 {
		panic(errors.New("<!--TOC--> 魔法鎮遭破壞，請確認 README 格式是否正確"))
	}
	before, after := readmeText[:TOCIndex[0][1]], readmeText[TOCIndex[1][0]:]

	ignoreList := getIgnoreFile()
	extension := "go"
	if len(os.Args) > 1 {
		extension = os.Args[1]
	}
	var fileInfoList []FileInfo
	fileInfoList = createTOCInfoList(&ignoreList, ".", extension)
	strTOC := ""
	writeTOCList("", ".", &strTOC, &fileInfoList)

	wf, err := os.Create(readmeName)
	util.CheckOpen(err)
	defer wf.Close()

	w := bufio.NewWriter(wf)
	w.WriteString(fmt.Sprintf("%s\n%s%s", before, strTOC, after))
	w.Flush()
}
