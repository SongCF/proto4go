package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"container/list"
	"path/filepath"
)



// walk dir and sub-dir, get all files with suffix
func WalkDir(dirPth, suffix string) (*list.List, error) {
	files := list.New()
	suffix = strings.ToUpper(suffix)
	err := filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			absName := getAbsName(filename)
			files.PushBack(absName)
		}
		return nil
	})
	return files, err
}

// walk dir without sub-dir, get all files with suffix
func ListDir(dirPth string, suffix string) (*list.List, error) {
	files := list.New()
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //ignore
	for _, fi := range dir {
		if fi.IsDir() { // ignore sub-dir
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files.PushBack(dirPth + PthSep + fi.Name())
		}
	}
	return files, nil
}



func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Error: get current directory failed!")
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func getAbsName(file string) string {
	abs, err := filepath.Abs(file)
	check(err)
	return strings.Replace(abs, "\\", "/", -1)
}

