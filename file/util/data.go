package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

const readmeName = "README.md"

func GetIgnoreFile(fileName string) []string {
	var ignoreList []string
	f, err := os.Open(fileName)
	//CheckOpen(err)
	defer f.Close()

	if err != nil {
	} else {
		r := bufio.NewScanner(f)
		for r.Scan() {
			text := r.Text()
			ignoreList = append(ignoreList, text)
		}
		CheckRead(r.Err())
	}

	return ignoreList
}

func GetReadmeText() (before string, after string) {
	readmeName := readmeName
	rf, err := os.Open(readmeName)
	CheckRead(err)
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
	before, after = readmeText[:TOCIndex[0][1]], readmeText[TOCIndex[1][0]:]
	return
}

func WriteReadme(before, after, toc string) {
	wf, err := os.Create(readmeName)
	CheckOpen(err)
	defer wf.Close()

	w := bufio.NewWriter(wf)
	w.WriteString(fmt.Sprintf("%s\n%s%s", before, toc, after))
	w.Flush()
}

func CheckIgnore(ignoreList *[]string, fileName *string) bool {
	isReplace := false
	for _, ignoreName := range *ignoreList {
		if *fileName == ignoreName {
			isReplace = true
			break
		}
	}
	return isReplace
}