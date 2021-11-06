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

type Parser struct {
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

func New(filePath string) Parser {
	return Parser{clippingsFilePath: filePath}
}
// Parse reads clipping file and returns []Clipping which contain metadata  
// extracted from each clipping section in the file.
func (p Parser) Parse() (map[string][]Clipping, error) {
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
		title:        strings.TrimFunc(titleAuthorMatch[0],IsUnicodeSpecial),
		author:       titleAuthorMatch[1],
		clippingType: clippingTypePageOrLoc[0],
		pageOrLoc:    clippingTypePageOrLoc[3],
		loc:          clippingLocOnly[0],
		clippingDate: strings.Join(clippingDateTime, " "),
		content:      strings.TrimFunc(clippingSection[2],IsUnicodeSpecial),
	}
}
