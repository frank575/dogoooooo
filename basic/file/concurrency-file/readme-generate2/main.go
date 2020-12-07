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

func setFileInfo(mx *sync.Mutex, infoListMap *map[string][]FileInfo, class string, name *string, path *string) {
	mx.Lock()
	if infoList, ok := (*infoListMap)[class]; ok {
		(*infoListMap)[class] = append(infoList, FileInfo{*name, *path})
	} else {
		(*infoListMap)[class] = []FileInfo{{*name, *path}}
	}
	mx.Unlock()
}

func getFileInfo(mx *sync.Mutex, wg *sync.WaitGroup, infoListMap *map[string][]FileInfo, path *string, name *string) {
	f, _ := os.Open(*path)
	r := bufio.NewScanner(f)
	reg, _ := regexp.Compile(`(?i)toc#.*#toc.*`)

	for r.Scan() {
		text := r.Text()
		line := reg.FindString(text)
		if *name == "README.md" {
			titleReg, _ := regexp.Compile("^#\\s*.+")
			if titleReg.MatchString(text) {
				hashReg, _ := regexp.Compile("^#\\s*")
				title := hashReg.ReplaceAllString(text, "")
				setFileInfo(mx, infoListMap, "project", &title, path)
			}
			break
		}
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
			setFileInfo(mx, infoListMap, typeName, &name, path)
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
		//ext := filepath.Ext(fileName) 取 extension
		newPath := path + "/" + fileName
		notIgnore := !util.CheckIgnore(ignoreList, &newPath)
		isDirAndNotIgnore := isDir && notIgnore
		notDirAndNotIgnore := !isDir && notIgnore

		exec.Command("clear")

		if isDirAndNotIgnore {
			wg.Add(1)
			fmt.Println("dir:", newPath)
			go getTOCList(wg, mx, ignoreList, infoListMap, newPath)
		} else if notDirAndNotIgnore {
			wg.Add(1)
			fmt.Println("file:", newPath)
			go getFileInfo(mx, wg, infoListMap, &newPath, &fileName)
		}
	}

	wg.Done()
}

func writeProjectTOCText(infoListMap *map[string][]FileInfo, toc *string) {
	const projectKey = "project"
	if projectList, ok := (*infoListMap)[projectKey]; ok {
		delete(*infoListMap, projectKey)
		*toc += fmt.Sprintf("- **⭐%s**\n", projectKey)
		for i, title := range projectList {
			*toc += fmt.Sprintf("  - [%d. %s](%s)\n", i+1, title.name, title.path)
		}
		*toc += "---\n"
	}
}

func createTOC(infoListMap *map[string][]FileInfo, toc *string) {
	var keys []string

	writeProjectTOCText(infoListMap, toc)

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
			reg, _ := regexp.Compile("\\s*,\\s*")
			sp := strings.Replace(line, ig, "", 1)
			if sp != "" {
				list := reg.Split(sp, -1)
				ignoreList = list
			}
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
