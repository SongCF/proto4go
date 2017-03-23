package main

import (
	"fmt"
	"strings"
	"strconv"
	"io"
	"os"
	"bufio"
)

func parseFile(fileName string, cmdMap *map[int]*Msg) {
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Print("Error: parse error!")
		panic(err)
	}

	buf := bufio.NewReader(f)
	lastLine := ""
	tmpLine := lastLine
	for {
		tmpLine = lastLine
		// line
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Print("Error: read line error!")
			panic(err)
		}

		//parse cmd
		idx := strings.Index(line, CMD_FLAG)
		if idx < 0 {
			lastLine = line
			continue
		}
		cmd := line[idx+len(CMD_FLAG) : len(line)-1]
		id, err := strconv.Atoi(cmd)
		if err != nil {
			fmt.Println("Error format!")
			panic(err)
		}
		//parse message
		line, err = buf.ReadString('\n')
		if err != nil {
			fmt.Println("Error format!")
			panic(err)
		}
		idx = strings.Index(line, MSG_FLAG)
		if idx < 0 {
			fmt.Println("Error format!")
			panic(line)
		}
		idx2 := strings.Index(line, " {")
		msg := line[idx+len(MSG_FLAG):idx2]
		//parse describe
		idx = strings.Index(tmpLine, DESC_FLAG)
		des := ""
		if idx >= 0 {
			des = tmpLine[idx+len(DESC_FLAG) : len(tmpLine)-1]
		}

		// check duplicate
		if _, ok := (*cmdMap)[id]; ok {
			fmt.Println("Error: duplicate cmd!")
			panic("Error: duplicate cmd!")
		}
		(*cmdMap)[id] = &Msg{
			cmd:  cmd,
			name: msg,
			desc: des,
		}

		lastLine = line
	}
}

