package util

import (
	"fmt"
	"math/rand"
	"time"
)

type PathList struct {
	List []string
	IMethods
}

type IMethods interface {
	GetFilePathList() []string
	GetRandFilePathList(num int) []string
}

func CreatePathList() PathList {
	var pl PathList
	pl.List = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	return pl
}

func (pl *PathList) GetFilePathList() []string {
	var list []string
	for _, name := range pl.List {
		list = append(list, fmt.Sprintf("basic/concurrency-file/files/%s.txt", name))
	}
	return list
}

func checkIndexRepeat(list *[]int, r *int) bool {
	repeat := 0
	for _, i := range *list {
		if i == *r {
			repeat++
		}
	}
	return repeat > 0
}

func randTakeFileIndexList(num, size int) []int {
	var list []int
	for {
		if len(list) == num {
			break
		}
		rd := rand.New(rand.NewSource(time.Now().UnixNano()))
		r := rd.Intn(size)
		if len(list) > 0 {
			for {
				if checkIndexRepeat(&list, &r) {
					rd := rand.New(rand.NewSource(time.Now().UnixNano()))
					r = rd.Intn(size)
				} else {
					break
				}
			}
		}
		list = append(list, r)
	}
	return list
}

func (pl *PathList) GetRandFilePathList(num int) []string {
	var list []string
	for _, name := range pl.List {
		list = append(list, fmt.Sprintf("basic/file/concurrency-file/files/%s.txt", name))
	}
	randIndexList := randTakeFileIndexList(num, len(pl.List))
	return []string{list[randIndexList[0]], list[randIndexList[1]], list[randIndexList[2]]}
}
