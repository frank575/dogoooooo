//##file:::(新)併發生成 README 目錄簡述##
package main

import (
	"bufio"
	"dogoooooo/file/util"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type FileInfo struct {
	typeId   int
	typeName string
	name     string
	path     string
}

func getFileInfo(mx *sync.Mutex, infoList *[]FileInfo, typeIdMap *map[string]int, path *string) {
	f, _ := os.Open(*path)
	r := bufio.NewScanner(f)
	reg, _ := regexp.Compile("##.+##")
	fileContext := ""

	for r.Scan() {
		fileContext += r.Text() + "\n"
	}

	info := reg.FindAllString(fileContext, 1)
	if len(info) > 0 {
		typeId := 0
		mapLen := len(*typeIdMap)
		hashReg, _ := regexp.Compile("\\s?##\\s?")
		info := hashReg.ReplaceAllString(info[0], "")
		sp := strings.Split(info, ":::")
		var name, typeName string
		if len(sp) > 1 {
			typeName = sp[0]
			name = sp[1]
		} else {
			typeName = "無分類"
			name = sp[0]
		}
		if id, ok := (*typeIdMap)[typeName]; ok {
			typeId = id
		} else {
			mx.Lock()
			_id := mapLen
			(*typeIdMap)[typeName] = _id
			typeId = _id
			mx.Unlock()
		}
		*infoList = append(*infoList, FileInfo{typeId, typeName, name, *path})
	}
}

func getTOCList(wg *sync.WaitGroup, mx *sync.Mutex, ignoreList *[]string, infoList *[]FileInfo, typeIdMap *map[string]int, path string) {
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
			go getTOCList(wg, mx, ignoreList, infoList, typeIdMap, newPath)
		} else if notDirAndNotIgnore {
			getFileInfo(mx, infoList, typeIdMap, &newPath)
		}
		wg.Done()
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex
	var infoList []FileInfo
	before, after := util.GetReadmeText()
	ignoreList := util.GetIgnoreFile()
	strTOC := ""
	typeIdMap := map[string]int{"無分類": 0}

	wg.Add(1)
	//TODO 可能會少資料
	go getTOCList(&wg, &mx, &ignoreList, &infoList, &typeIdMap, ".")
	wg.Wait()

	sort.Slice(infoList, func(i, j int) bool {
		return infoList[i].typeId < infoList[j].typeId
	})

	fmt.Println(len(infoList))
	fmt.Print(infoList)
	return
	util.WriteReadme(before, after, strTOC)
}
