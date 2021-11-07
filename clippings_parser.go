package main

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	titleAuthorRegexPattern           = `(.*?)\s*\((.*?)\)`
	clippingLocRegexPattern           = `\s*\\|\s*location\s*(\d+-\d+|\d+)`
	clippingDateTimeRegexPattern      = `\s*\\|\s*Added\s*on\s*(.*?),\s*(\d+)\s*(.*?)\s*(\d+)\s*(\d+:\d+:\d+)`
	clippingTypePageOrLocRegexPattern = `Your\s*(.*?)\s*(on|at)\s*(page|location)\s*(\d+-\d+|\d+)`
	
)

var clippingsSeparatorRegEx = regexp.MustCompile("====*")

type ClippingsParser struct {
	clippingsFilePath string
}

type Clipping struct {
	title        string
	author       string 
	clippingType string
	pageOrLoc    string
	loc          string
	clippingDate string
	content      string
}

func NewClippingsParser(filePath string) ClippingsParser {
	return ClippingsParser{clippingsFilePath: filePath}
}
// Parse reads clipping file and returns []Clipping which contain metadata  
// extracted from each clipping section in the file.
func (p ClippingsParser) Parse() (map[string][]Clipping, error) {
	fileContent, err := readClippingsFile(p)

	if err != nil {
		return nil, err
	}

	return transformStringToClipping(fileContent), nil

}

func readClippingsFile(parser ClippingsParser) (map[string][]string, error) {

	if fileExists(parser.clippingsFilePath) {

		clippingFile, openError := os.Open(parser.clippingsFilePath)

		if openError != nil {
			return nil, errors.New("unable to read clippings file")
		}

		defer clippingFile.Close()

		scanner := bufio.NewScanner(clippingFile)
      		
		return readClippingsLineByLine(scanner),nil		
	}

	return nil, errors.New("Clipping file not found in provided file path")
}

func readClippingsLineByLine(scanner *bufio.Scanner) map[string][]string {
	
	contents := make(map[string][]string, 0)
	clipping := make([]string, 0)
    	
	scannedSectionCounter := 0

	for scanner.Scan() {

		currentLine := strings.TrimFunc(scanner.Text(),IsUnicodeSpecial)
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
	return contents
}

func transformStringToClipping(clippingSections map[string][]string) map[string][]Clipping {

	clippings := make(map[string][]Clipping, 0)

	for _, section := range clippingSections {
		currentClipping := extractClippingMetaData(section)

		if clippingsForTitle,found := clippings[currentClipping.title]; found {
            clippingsForTitle = append(clippingsForTitle,currentClipping)
			clippings[currentClipping.title] = clippingsForTitle
		} else {
			clippingForTitle := make([]Clipping,0)
			clippingForTitle = append(clippingForTitle, currentClipping)
			clippings[currentClipping.title] = clippingForTitle
		}		
	}
	return clippings
}

func extractClippingMetaData(clippingSection []string) Clipping {

	titleAuthorMatch := findPatternMatchesInString(titleAuthorRegexPattern, clippingSection[0])
	clippingTypePageOrLoc := findPatternMatchesInString(clippingTypePageOrLocRegexPattern, clippingSection[1])
	clippingLocOnly := findPatternMatchesInString(clippingLocRegexPattern,clippingSection[1])
	clippingDateTime := findPatternMatchesInString(clippingDateTimeRegexPattern, clippingSection[1])

	return Clipping{
		title:        titleAuthorMatch[0],
		author:       titleAuthorMatch[1],
		clippingType: clippingTypePageOrLoc[0],
		pageOrLoc:    clippingTypePageOrLoc[3],
		loc:          clippingLocOnly[0],
		clippingDate: strings.Join(clippingDateTime, " "),
		content:      clippingSection[2],
	}
}
