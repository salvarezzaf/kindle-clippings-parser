package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)


type Parser struct {
	clippingsFilePath string
}

type Clipping struct {
	title        string
	author       string
	clippingType string
	pageNumber   int
	clippingDate string
	content      string
}

func New(filePath string) Parser {
	return Parser{clippingsFilePath: filePath}
}

func (p Parser) Parse() []Clipping{
	fileContent := readClippingsFile(p)
	transformStringToClipping(fileContent)

	test := make([]Clipping,0)

	return test
	
}

func readClippingsFile(parser Parser) map[string][]string {

	clippingsSeparatorRegEx := regexp.MustCompile("====*")

	if clippingsFileExists(parser.clippingsFilePath) {

		clippingFile, openError := os.Open(parser.clippingsFilePath)
		isError(openError)
		defer clippingFile.Close()

		contents := make(map[string][]string, 0)
		clipping := make([]string, 0)

		scanner := bufio.NewScanner(clippingFile)
		scannedSectionCounter := 0
		for scanner.Scan() {

			currentLine := scanner.Text()
			if clippingsSeparatorRegEx.MatchString(currentLine) {
				contents["section"+ strconv.Itoa(scannedSectionCounter)]= clipping
				clipping = clipping[:0]
				scannedSectionCounter++
			} else {
				if len(currentLine) > 0 {
					clipping = append(clipping, currentLine)
				}
			}

		}
		return contents
	}

	return nil
}

func isError(anError error) {
	if anError != nil {
		log.Fatalf("unable to read clippings file %v", anError)
	}
}

func clippingsFileExists(clippingsFilePath string) bool {
	info, err := os.Stat(clippingsFilePath)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func transformStringToClipping(clippingSections map[string][]string) []Clipping {

    clippings := make([]Clipping,0)

	for _, section := range clippingSections {

		clippings= append(clippings,extractClippingMetaData(section))

	}
	return clippings	
}

func extractClippingMetaData(clippingSection []string) Clipping {
    
	titleAuthor := strings.FieldsFunc(clippingSection[0], Split)
    pageDetails:= strings.FieldsFunc(clippingSection[1], Split)
	typeAndPage := strings.TrimSpace(pageDetails[0])
	isHighlight :=  regexp.MustCompile(`Your\s*(.*?)\s*on`)
    //matches:= isHighlight.FindAllStringSubmatch(typeAndPage,-1)
    
     

	return Clipping{
		title: titleAuthor[0],
		author: titleAuthor[1],
	

	}
}

func Split(r rune) bool {
	return r == '(' || r == ')' || r == '|' || r == '-'
}
