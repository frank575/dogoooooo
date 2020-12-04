// toc#file:::(新)併發生成 README 目錄簡述#toc
package main

import (
	"bufio"
	"dogoooooo/file/util"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type FileInfo struct {
	name string
	path string
}

func getFileInfo(mx *sync.Mutex, wg *sync.WaitGroup, infoListMap *map[string][]FileInfo, path *string) {
	f, _ := os.Open(*path)
	r := bufio.NewScanner(f)
	reg, _ := regexp.Compile(`(?i)toc#.*#toc.*`)
	fileContext := ""

	for r.Scan() {
		fileContext += r.Text() + "\n"
	}

	info := reg.FindAllString(fileContext, 1)
	if len(info) > 0 {
		hashReg, _ := regexp.Compile(`(?i)(toc#)|(#toc.*)`)
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

	wg.Done()
}

func getTOCList(wg *sync.WaitGroup, mx *sync.Mutex, ignoreList *[]string, infoListMap *map[string][]FileInfo, path string) {
	dirList, err := ioutil.ReadDir(path)
	util.CheckReadDir(err)

	for i := range dirList {
		file := dirList[i]
		fileName := file.Name()
		isDir := file.IsDir()
		notIgnore := !util.CheckIgnore(ignoreList, &fileName)
		isDirAndNotIgnore := isDir && notIgnore
		notDirAndNotIgnore := !isDir && notIgnore
		newPath := path + "/" + fileName

		exec.Command("clear")
		fmt.Println(newPath)

		if isDirAndNotIgnore {
			wg.Add(1)
			go getTOCList(wg, mx, ignoreList, infoListMap, newPath)
		} else if notDirAndNotIgnore {
			wg.Add(1)
			go getFileInfo(mx, wg, infoListMap, &newPath)
		}
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
	//TODO .custom-ignore -> .custom-generate 要更新 ROOT_PATH= 及 IGNORE=
	ignoreList = append(ignoreList, util.GetIgnoreFile(".custom-ignore")...)
	strTOC := ""
	infoListMap := map[string][]FileInfo{}

	wg.Add(1)
	getTOCList(&wg, &mx, &ignoreList, &infoListMap, ".")
	wg.Wait()

	createTOC(&infoListMap, &strTOC)
	util.WriteReadme(before, after, strTOC)
}
