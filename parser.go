package main

import (
	"log"
	"regexp"
	"strings"
)

type section struct {
	name string
	todo []todo
}

type todo struct {
	level    int
	contents string
	parent   *todo
	children []*todo
	path     string
}

var sections map[string]section
var sectionRegex string = "==\\w+=="

func parseTodo(todofile []byte) {
	filestring := string(todoFile)
	lines := strings.Split(filestring, "\n")
	for i, line := range lines {
		var sectionName string
		matched, err := regexp.MatchString(sectionRegex, line)
		checkErr(err)
		if matched == true {
			sectionName = strings.Trim(line, "==")
			log.Printf("Section %v found", sectionName)
			sections[sectionName] = section{name: sectionName}
		} else if line == "" {
			log.Println("Moving to the next section")
		} else {
			if line[0] == 32 { // 32 is an asterisk
				log.Printf("Todo line found at %v", i)
				//sections[sectionName].todo := append(sections[sectionName].todo, todoStruct)
			} else if line[0] == 9 { // 9 is a tab
				log.Printf("At least one tab found on line %v", i)
			} else if line[0] == 62 { // 62 is a right angle bracket
				log.Printf("Notefile line found at %v", i)
			} else {
				log.Printf("Unknown char: %v, line %v", line[0], i)
			}
		}
	}
	log.Println(sections)
}
