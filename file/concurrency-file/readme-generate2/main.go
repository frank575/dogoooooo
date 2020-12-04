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
	name string
	path string
}

func getFileInfo(mx *sync.Mutex, infoListMap *map[string][]FileInfo, path *string) {
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
			typeName = "無分類"
			name = sp[0]
		}
		mx.Lock()
		if infoList, ok := (*infoListMap)[typeName]; ok {
			(*infoListMap)[typeName] = append(infoList, FileInfo{name, *path})
		} else {
			(*infoListMap)[typeName] = []FileInfo{{name, *path}}
		}
		mx.Unlock()
	}
}

func getTOCList(wg *sync.WaitGroup, mx *sync.Mutex, ignoreList *[]string, infoListMap *map[string][]FileInfo, path string) {
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
			go getTOCList(wg, mx, ignoreList, infoListMap, newPath)
		} else if notDirAndNotIgnore {
			getFileInfo(mx, infoListMap, &newPath)
		}
		wg.Done()
	}

	wg.Done()
}

func createTOC(infoListMap *map[string][]FileInfo, toc *string) {
	var keys []string
	for k := range *infoListMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for i, k := range keys {
		index := i + 1
		*toc += fmt.Sprintf("- **%d. %s**\n", index, k)
		for j, info := range (*infoListMap)[k] {
			*toc += fmt.Sprintf("  - [%d-%d. %s](%s)\n", index, j+1, info.name, info.path)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex
	before, after := util.GetReadmeText()
	ignoreList := util.GetIgnoreFile(".gitignore")
	ignoreList = append(ignoreList, util.GetIgnoreFile(".custom-ignore")...)
	strTOC := ""
	infoListMap := map[string][]FileInfo{}

	wg.Add(1)
	go getTOCList(&wg, &mx, &ignoreList, &infoListMap, ".")
	wg.Wait()

	createTOC(&infoListMap, &strTOC)
	util.WriteReadme(before, after, strTOC)
}
