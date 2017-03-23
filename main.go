//
// Copyright (c) 2017 SongCF <songcf_faith@foxmail.com>.
// https://github.com/golang/protobuf
//

package main

import (
	"container/list"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

const (
	SUFFIX    = ".proto"
	CMD_FLAG  = "// cmd:"
	MSG_FLAG  = "message "
	DESC_FLAG = "// "
	FILE_NAME = "msgcode"
)

type Msg struct {
	cmd  string
	name string
	desc string
}

// ./proto4go -i IN_DIR -o OUT_DIR
// ./proto4go -i . -o ./proto
func main() {
	argNum := len(os.Args)
	if argNum != 5 {
		fmt.Printf("Error args num: %v \nRight is: ./proto4go -i IN_DIR -o OUT_DIR", argNum)
		return
	}

	if !strings.EqualFold(os.Args[1], "-i") || !strings.EqualFold(os.Args[3], "-o") {
		fmt.Print("Error args format. \nRight is: ./proto4go -i IN_DIR -o OUT_DIR\n")
		return
	}

	inDir := os.Args[2]
	outDir, err := filepath.Abs(os.Args[4])
	check(err)
	outDir = strings.Replace(outDir, "\\", "/", -1)
	//fmt.Println("out dir1: ", os.Args[4])
	fmt.Println("out dir2: ", outDir)

	protoFileL, err := WalkDir(inDir, SUFFIX)
	if err != nil {
		fmt.Println("Error: walk dir failed!")
		return
	}

	parse(protoFileL, outDir)

	fmt.Println("Success!")
}

func parse(fileL *list.List, outDir string) {
	cmdMap := make(map[int]*Msg)
	for elem := fileL.Front(); elem != nil; elem = elem.Next() {
		// do something with elem.Value
		fileName := elem.Value.(string)
		fmt.Printf("parse file: %v\n", fileName)
		// gen *.pb.go
		workDir := filepath.Dir(fileName)
		curDir := getCurrentDirectory()
		err := os.Chdir(workDir)
		check(err)
		base := filepath.Base(fileName)
		command := exec.Command("protoc", "--go_out="+outDir, base)
		//fmt.Printf("command: %v\n", command)
		err = command.Run()
		check(err)
		err = os.Chdir(curDir)
		check(err)
		// parse file
		parseFile(fileName, &cmdMap)
	}
	keys := make([]int, 0)
	for k := range cmdMap {
		keys = append(keys, k)
		//fmt.Println("id:", k, ", name:", cmdMap[k].name, ", desc:", cmdMap[k].desc)
	}
	sort.Ints(keys)

	writeCode(keys, &cmdMap, outDir)
	writeCSV(keys, &cmdMap, outDir)
}


func check(err error) {
	if err != nil {
		panic(err)
	}
}


