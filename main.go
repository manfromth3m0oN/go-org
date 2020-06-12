package main

import (
	"io/ioutil"
	"log"
	"os"
)

var noteDir string = "/home/m0on/go/src/github.com/manfromth3m0on/go-org/testNotes"
var files []os.FileInfo

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

func main() {
	log.Println("Welcome to go-org")
	log.Println("-----------------")
	files = getFiles(noteDir)
	log.Println("Tracked Files: ")
	for _, file := range files {
		log.Println(file.Name())
	}
}
