//##file:::(新)併發生成 README 目錄簡述##
package main

import (
	"bufio"
	"dogoooooo/file/util"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"sync"
)

type FileInfo struct {
	typeId   int
	typeName string
	name     string
	path     string
}

func getFileInfo(infoList *[]FileInfo, path *string) {
	f, _ := os.Open(*path)
	r := bufio.NewScanner(f)
	reg, _ := regexp.Compile("##.+##")
	fileContext := ""

	for r.Scan() {
		fileContext += r.Text() + "\n"
	}

	info := reg.FindAllString(fileContext, 1)
	if len(info) > 0 {
		hashReg, _ := regexp.Compile("\\s?##\\s?")
		info := hashReg.ReplaceAllString(info[0], "")
		sp := strings.Split(info, ":::")
		var name, typeName string
		if len(sp) > 1 {
			typeName = sp[0]
			name = sp[1]
		} else {
			name = sp[0]
		}
		//TODO typeId 還沒撈
		*infoList = append(*infoList, FileInfo{1, typeName, name, *path})
	}
}

func getTOCList(wg *sync.WaitGroup, ignoreList *[]string, infoList *[]FileInfo, path string) {
	dirList, err := ioutil.ReadDir(path)
	util.CheckReadDir(err)

	for i := range dirList {
		wg.Add(1)
		file := dirList[i]
		fileName := file.Name()
		isDir := file.IsDir()
		notIgnore := !util.CheckIgnore(ignoreList, &fileName)
		isDirAndNotIgnore := isDir && notIgnore
		notDirAndNotIgnore := !isDir && notIgnore
		newPath := path + "/" + fileName

		if isDirAndNotIgnore {
			wg.Add(1)
			go getTOCList(wg, ignoreList, infoList, newPath)
		} else if notDirAndNotIgnore {
			getFileInfo(infoList, &newPath)
		}
		wg.Done()
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	//before, after := util.GetReadmeText()
	ignoreList := util.GetIgnoreFile()
	var infoList []FileInfo

	wg.Add(1)
	go getTOCList(&wg, &ignoreList, &infoList, ".")
	wg.Wait()

	fmt.Print(infoList)
	//util.WriteReadme(before, after, "")
}
