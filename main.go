package main

import (
	"io/ioutil"
	"log"
	"os"
)

var noteDir string = "/home/m0on/go/src/github.com/manfromth3m0on/go-org/testNotes/"
var files []os.FileInfo
var todoFile []byte

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getFiles(dir string) []os.FileInfo {
	dirFiles, err := ioutil.ReadDir(dir)
	checkErr(err)
	return dirFiles
}

func getTodoFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	checkErr(err)
	return file
}

func main() {
	log.Println("Welcome to go-org")
	log.Println("-----------------")
	files = getFiles(noteDir)
	log.Println("Tracked Files: ")
	for _, file := range files {
		if file.Name() == "TODO.txt" {
			log.Println("Found TODO.txt")
			todoFile = getTodoFile(noteDir + file.Name())
		}
		log.Println(file.Name())
	}
	parseTodo(todoFile)
}
