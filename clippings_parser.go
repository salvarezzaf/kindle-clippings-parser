package main

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const(
	titleAuthorRegexPattern = `(.*?)\s*\((.*?)\)`
	clippingLocRegexPattern = `\s*\\|\s*location\s*(\d+-\d+|\d+)` 
	clippingDateTimeRegexPattern = `\s*\\|\s*Added\s*on\s*(.*?),\s*(\d+)\s*(.*?)\s*(\d+)\s*(\d+:\d+:\d+)`
	clippingTypeAndPageRegexPattern = `Your\s*(.*?)\s*on\s*page\s*([0-9]+)`
)

type Parser struct {
	clippingsFilePath string
}

type Clipping struct {
	title        string
	author       string
	clippingType string
	pageNumber   string
	loc 		 string	
	clippingDate string
	content      string
}

func New(filePath string) Parser {
	return Parser{clippingsFilePath: filePath}
}

func (p Parser) Parse() ([]Clipping, error) {
	fileContent, err := readClippingsFile(p)

	if err != nil {
		return nil, err
	}

	return transformStringToClipping(fileContent), nil

}

func readClippingsFile(parser Parser) (map[string][]string, error) {

	clippingsSeparatorRegEx := regexp.MustCompile("====*")

	if fileExists(parser.clippingsFilePath) {

		clippingFile, openError := os.Open(parser.clippingsFilePath)

		if openError != nil {
			return nil, errors.New("unable to read clippings file")
		}

		defer clippingFile.Close()

		contents := make(map[string][]string, 0)
		clipping := make([]string, 0)

		scanner := bufio.NewScanner(clippingFile)
		scannedSectionCounter := 0
		for scanner.Scan() {

			currentLine := scanner.Text()
			if clippingsSeparatorRegEx.MatchString(currentLine) {
				contents["section"+strconv.Itoa(scannedSectionCounter)] = append([]string(nil), clipping...)
				clipping = clipping[:0]
				scannedSectionCounter++
			} else {
				if len(currentLine) > 0 {
					clipping = append(clipping, currentLine)
				}
			}

		}
		return contents, nil
	}

	return nil, errors.New("Clipping file not found in provided file path")
}

func transformStringToClipping(clippingSections map[string][]string) []Clipping {

	clippings := make([]Clipping, 0)

	for _, section := range clippingSections {

		clippings = append(clippings, extractClippingMetaData(section))

	}
	return clippings
}

func extractClippingMetaData(clippingSection []string) Clipping {

	titleAuthorMatch := findPatternMatchesInString(titleAuthorRegexPattern,clippingSection[0])
	clippingTypeAndpage := findPatternMatchesInString(clippingTypeAndPageRegexPattern, clippingSection[1])
	clippingLoc := findPatternMatchesInString(clippingLocRegexPattern,clippingSection[1])
	clippingDateTime := findPatternMatchesInString(clippingDateTimeRegexPattern,clippingSection[1])
   
	return Clipping{
		title:        titleAuthorMatch[0],
		author:       titleAuthorMatch[1],
		clippingType: clippingTypeAndpage[0],
		pageNumber:   clippingTypeAndpage[1],
		loc: clippingLoc[0],
		clippingDate: strings.Join(clippingDateTime," "),
		content:      clippingSection[2],
	}
}
