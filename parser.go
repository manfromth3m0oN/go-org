package main

import (
	"log"
	"strings"
)

type section struct {
	name string
	contents []string
}

func parseTodo(todofile []byte) {
	filestring := string(todoFile)
	lines := strings.Split(filestring, "\n")
	log.Println(lines)
}
