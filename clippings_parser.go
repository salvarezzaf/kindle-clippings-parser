package main

import (
	"bufio"
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

func (p Parser) Parse() []Clipping {
	fileContent := readClippingsFile(p)
	return transformStringToClipping(fileContent)

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

	clippings := make([]Clipping, 0)

	for _, section := range clippingSections {

		clippings = append(clippings, extractClippingMetaData(section))

	}
	return clippings
}

func extractClippingMetaData(clippingSection []string) Clipping {

	titleAuthorDetails := strings.FieldsFunc(clippingSection[0], Split)
	typePageAndDateDetails := strings.FieldsFunc(clippingSection[1], Split)
	typeAndPageMatches := getRegexMatches(`Your\s*(.*?)\s*on\s*page\s*([0-9]+)`, typePageAndDateDetails[0])
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

func getRegexMatches(regexString string, stringToSearch string) []string {
	isHighlight := regexp.MustCompile(regexString)
	matches := isHighlight.FindAllStringSubmatch(stringToSearch, -1)
	matchGroups := make([]string, 0)

	for _, match := range matches {
		for i := 1; i < len(match); i++ {
			if !contains(matchGroups, match[i]) {
				matchGroups = append(matchGroups,strings.TrimSpace(match[i]))
			}
		}
	}
	return matchGroups
}

func Split(r rune) bool {
	return r == '(' || r == ')' || r == '|' || r == '-' || r == ','
}

func contains(aSlice []string, elementToSearch string) bool {
	for _, element := range aSlice {
		if elementToSearch == element {
			return true
		}
	}
	return false
}
