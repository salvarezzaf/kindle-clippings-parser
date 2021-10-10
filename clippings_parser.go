package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)


var clippingsFilePath string 

type Clipping struct {
	title string
	author string
	clippingType string
	pageNumber int
	clippingDate string
	content string
}

func New(filePath string) Clipping {
	clippingsFilePath = filePath
	return Clipping{}
}

func (c Clipping) Parse() {
  	content := readClippingsFile()
    transformStringToClipping(content)  
}

func readClippingsFile() string {
	if clippingsFileExists(){
		content, err := ioutil.ReadFile(clippingsFilePath)
		if(err != nil) {
			log.Fatalf("unable to read clippings file %v", err)
		}

		return string(content)
	}
	return ""
}


func clippingsFileExists() bool {
	info, err := os.Stat(clippingsFilePath)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func transformStringToClipping(clippingsAsString string){

	clippingSections := strings.Split(clippingsAsString,"==========")

	for _, section := range clippingSections {
		
		fmt.Println(section)
	}
}