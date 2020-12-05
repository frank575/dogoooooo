// toc#file:::(新)併發生成 README 目錄簡述#toc
package main

import (
	"bufio"
	"dogoooooo/basic/file/util"
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

	for r.Scan() {
		line := reg.FindString(r.Text())
		if line != "" {
			hashReg, _ := regexp.Compile(`(?i)(toc#)|(#toc.*)`)
			info := hashReg.ReplaceAllString(line, "")
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
			break
		}
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

func getGTOCSetting() (ignoreList []string, rootPath string) {
	f, err := os.Open("gtoc.txt")
	if err != nil {
		fmt.Println("gtoc not found!")
	}
	defer f.Close()

	rootPath = "."

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		const rp = "ROOT_PATH="
		const ig = "IGNORE="
		if rRoot, _ := regexp.Compile(rp); rRoot.MatchString(line) {
			sp := strings.Split(line, rp)
			rootPath = strings.TrimSpace(sp[1])
		} else if rIgnore, _ := regexp.Compile(ig); rIgnore.MatchString(line) {
			reg, _ := regexp.Compile(",\\s*")
			sp := strings.Replace(line, ig, "", 1)
			list := reg.Split(sp, -1)
			ignoreList = list
		}
	}

	return
}

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex
	before, after := util.GetReadmeText()
	ignoreList := util.GetGitIgnoreFile()
	sIgnoreList, sRootPath := getGTOCSetting()
	ignoreList = append(ignoreList, sIgnoreList...)
	strTOC := ""
	infoListMap := map[string][]FileInfo{}

	wg.Add(1)
	getTOCList(&wg, &mx, &ignoreList, &infoListMap, sRootPath)
	wg.Wait()

	createTOC(&infoListMap, &strTOC)
	util.WriteReadme(before, after, strTOC)
}