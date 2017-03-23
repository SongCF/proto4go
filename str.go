package main

import "strings"


// login_req -> LoginReq
func UpFirstChar(str string) string {
	s := []rune(str)
	for i:=0; i<len(s); i++ {
		if s[i] == '_' && i+1 < len(s) {
			s[i+1] = upChar(s[i+1])
		}
	}
	s[0] = upChar(s[0])
	return strings.Replace(string(s), "_", "", -1)
}

func upChar(c rune) rune {
	if c >= 'a' && c <= 'z' {
		return c + ('A' - 'a')
	}
	return c
}
