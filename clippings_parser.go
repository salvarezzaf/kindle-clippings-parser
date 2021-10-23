package main

import (
	"bufio"
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

func (p Parser) Parse() []Clipping {
	fileContent := readClippingsFile(p)
	return transformStringToClipping(fileContent)

}

func readClippingsFile(parser Parser) map[string][]string {

	clippingsSeparatorRegEx := regexp.MustCompile("====*")

	if fileExists(parser.clippingsFilePath) {

		clippingFile, openError := os.Open(parser.clippingsFilePath)
		isError(openError,"unable to read clippings file") 

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
		return contents
	}

	return nil
}

func transformStringToClipping(clippingSections map[string][]string) []Clipping {

	clippings := make([]Clipping, 0)

	for _, section := range clippingSections {

		clippings = append(clippings, extractClippingMetaData(section))

	}
	return clippings
}

func extractClippingMetaData(clippingSection []string) Clipping {

	titleAuthorDetails := strings.FieldsFunc(clippingSection[0], split)
	typePageAndDateDetails := strings.FieldsFunc(clippingSection[1], split)
	typeAndPageMatches := matchByRegex(`Your\s*(.*?)\s*on\s*page\s*([0-9]+)`, typePageAndDateDetails[0])
	pageNumberToInt, _ := strconv.Atoi(typeAndPageMatches[1])
	clippingDateTime := strings.TrimSpace(typePageAndDateDetails[len(typePageAndDateDetails)-1])

	return Clipping{
		title:        titleAuthorDetails[0],
		author:       titleAuthorDetails[1],
		clippingType: typeAndPageMatches[0],
		pageNumber:   pageNumberToInt,
		clippingDate: clippingDateTime,
		content: clippingSection[2],
	}
}